package auth

import (
	//"github.com/nikitads9/banner-service-api/internal/app/api"
	"context"
	"errors"
	"log/slog"
	"strings"

	"github.com/nikitads9/banner-service-api/internal/app/service/jwt"
	"github.com/nikitads9/banner-service-api/internal/logger/sl"
	desc "github.com/nikitads9/banner-service-api/pkg/banner-api"
)

type scope string

const (
	keyScope scope = ""
)

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

// HandleAdminToken ...
func (s Security) HandleAdminToken(ctx context.Context, _ string, t desc.AdminToken) (context.Context, error) {
	const op = "middleware.auth.handleAdminToken"

	log := s.logger.With(
		slog.String("op", op),
	)

	if t.APIKey == "" || !strings.HasPrefix(t.APIKey, "AdminToken ") {
		log.Error("missing token ", sl.Err(errMissingToken))
		return ctx, errMissingToken
	}

	token := strings.TrimPrefix(t.APIKey, "AdminToken ")
	scope, err := s.jwtService.VerifyToken(ctx, token)
	if err != nil {
		log.Error("issue verifying jwt token", sl.Err(err))
		return ctx, err
	}

	ctx = withScope(ctx, scope)

	return ctx, nil
}

// HandleUserToken ...
func (s Security) HandleUserToken(ctx context.Context, _ string, t desc.UserToken) (context.Context, error) {
	const op = "middleware.auth.handleUserToken"

	log := s.logger.With(
		slog.String("op", op),
	)

	if t.APIKey == "" || !strings.HasPrefix(t.APIKey, "UserToken ") {
		log.Error("missing token ", sl.Err(errMissingToken))
		return ctx, errMissingToken
	}

	token := strings.TrimPrefix(t.APIKey, "UserToken ")
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
	if scope, ok := ctx.Value(keyScope).(string); ok {
		return scope
	}

	return ""
}

func withScope(ctx context.Context, scope string) context.Context {
	return context.WithValue(ctx, keyScope, scope)
}
