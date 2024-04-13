package banner

import (
	"context"
	"crypto/tls"
	"net/http"

	"github.com/nikitads9/banner-service-api/internal/app/api"
	"github.com/nikitads9/banner-service-api/internal/logger/sl"
	"github.com/nikitads9/banner-service-api/internal/middleware/auth"
	"github.com/nikitads9/banner-service-api/internal/middleware/logger"
	desc "github.com/nikitads9/banner-service-api/pkg/banner-api"

	"errors"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	configType string
	pathConfig string

	bannerImpl      *api.Implementation
	serviceProvider *serviceProvider
	server          *http.Server
}

// NewApp ...
func NewApp(ctx context.Context, configType string, pathConfig string) (*App, error) {
	a := &App{
		configType: configType,
		pathConfig: pathConfig,
	}
	err := a.initDeps(ctx)

	return a, err
}

func (a *App) initDeps(ctx context.Context) error {

	inits := []func(context.Context) error{
		a.initServiceProvider,
		a.initServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider(a.configType, a.pathConfig)

	return nil
}

// Run ...
func (a *App) Run() error {
	defer func() {
		a.serviceProvider.db.Close() //nolint:errcheck
	}()

	err := a.startServer()
	if err != nil {
		a.serviceProvider.GetLogger().Error("failed to start server: %s", err)
		return err
	}

	return nil
}

func (a *App) initServer(ctx context.Context) error {
	tracer := a.serviceProvider.GetTracer(ctx)
	bannerService := a.serviceProvider.GetBannerService(ctx)
	a.bannerImpl = api.NewImplementation(bannerService, tracer)
	srv, err := desc.NewServer(a.bannerImpl, auth.NewSecurity(a.serviceProvider.GetLogger(), a.serviceProvider.GetJWTService(ctx)))
	if err != nil {
		return err
	}

	address := a.serviceProvider.GetConfig().GetAddress(a.serviceProvider.GetConfig().Server.Host, a.serviceProvider.GetConfig().Server.Port)

	routeFinder := logger.MakeRouteFinder(srv)
	a.server = &http.Server{
		Addr: address,
		Handler: logger.Wrap(srv,
			logger.LogRequests(routeFinder, a.serviceProvider.GetLogger()),
		),
		ReadTimeout:  a.serviceProvider.GetConfig().GetServerConfig().Timeout,
		WriteTimeout: a.serviceProvider.GetConfig().GetServerConfig().Timeout,
		IdleTimeout:  a.serviceProvider.GetConfig().GetServerConfig().IdleTimeout,
		TLSConfig: &tls.Config{
			MinVersion:               tls.VersionTLS13,
			PreferServerCipherSuites: true,
		},
	}

	return nil
}

func (a *App) startServer() error {
	if a.server == nil {
		a.serviceProvider.GetLogger().Error("server was not initialized")
		return errors.New("server was not initialized")
	}
	a.serviceProvider.GetLogger().Info("starting server", slog.String("address", a.serviceProvider.GetConfig().GetServerConfig().Port))

	done := make(chan os.Signal, 1)
	errChan := make(chan error)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		defer close(errChan)
		if err := a.server.ListenAndServe(); err != nil {
			a.serviceProvider.GetLogger().Error("", sl.Err(err))
			errChan <- err
		}
	}()

	a.serviceProvider.GetLogger().Info("server started")

	select {
	case err := <-errChan:
		return err
	case <-done:
		a.serviceProvider.GetLogger().Info("stopping server")
		// TODO: move timeout to config
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := a.server.Shutdown(ctx); err != nil {
			a.serviceProvider.GetLogger().Error("failed to stop server", sl.Err(err))
			return err
		}

		a.serviceProvider.GetLogger().Info("server stopped")
	}

	return nil
}
