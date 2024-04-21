package auth

import (
	"context"
	"errors"
	"log/slog"
	"strings"

	"github.com/nikitads9/banner-service-api/internal/app/service/jwt"
	"github.com/nikitads9/banner-service-api/internal/logger/sl"
	desc "github.com/nikitads9/banner-service-api/pkg/banner-api"
)

type KeyScope struct{}

var (
	errMissingToken = errors.New("missing bearer token")
	//errInvalidToken = errors.New("token is invalid")
)

// Security ...
type Security struct {
	logger     *slog.Logger
	jwtService jwt.Service
}

// NewSecurity ...
func NewSecurity(logger *slog.Logger, jwtService jwt.Service) Security {
	return Security{
		logger:     logger,
		jwtService: jwtService,
	}
}

// HandleBearer ...
func (s Security) HandleBearer(ctx context.Context, _ string, t desc.Bearer) (context.Context, error) {
	const op = "middleware.auth.handleToken"

	log := s.logger.With(
		slog.String("op", op),
	)

	if t.APIKey == "" || !strings.HasPrefix(t.APIKey, "Bearer ") {
		log.Error("missing token ", sl.Err(errMissingToken))
		return ctx, errMissingToken
	}

	token := strings.TrimPrefix(t.APIKey, "Bearer ")
	scope, err := s.jwtService.VerifyToken(ctx, token)
	if err != nil {
		log.Error("issue verifying jwt token", sl.Err(err))
		return ctx, err
	}

	ctx = withScope(ctx, scope)

	return ctx, nil
}

// ScopeFromContext ...
func ScopeFromContext(ctx context.Context) string {
	if scope, ok := ctx.Value(KeyScope{}).(string); ok {
		return scope
	}

	return ""
}

func withScope(ctx context.Context, scope string) context.Context {
	return context.WithValue(ctx, KeyScope{}, scope)
}
