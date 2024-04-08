package auth

import (
	//"github.com/nikitads9/banner-service-api/internal/app/api"
	"context"
	"errors"
	"log/slog"
	"strings"

	"github.com/nikitads9/banner-service-api/internal/app/service/jwt"
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

type Security struct {
	logger     *slog.Logger
	jwtService jwt.Service
}

func NewSecurity(logger *slog.Logger, jwtService jwt.Service) Security {
	return Security{
		logger:     logger,
		jwtService: jwtService,
	}
}

func (s Security) HandleAdminToken(ctx context.Context, operationName string, t desc.AdminToken) (context.Context, error) {
	const op = "middleware.auth.handleAdminToken"

	log := s.logger.With(
		slog.String("op", op),
	)

	if t.APIKey == "" || !strings.HasPrefix(t.APIKey, "AdminToken ") {
		log.Error("missing token ", errMissingToken)
		//api.WriteWithError(w, http.StatusUnauthorized, errMissingToken.Error())
		return ctx, errMissingToken
	}

	token := strings.TrimPrefix(t.APIKey, "AdminToken ")
	scope, err := s.jwtService.VerifyToken(ctx, token)
	if err != nil {
		log.Error("issue verifying jwt token", err)
		//api.WriteWithError(w, http.StatusUnauthorized, errInvalidToken.Error())
		return ctx, err
	}

	ctx = withScope(ctx, scope)

	return ctx, nil
}

func (s Security) HandleUserToken(ctx context.Context, operationName string, t desc.UserToken) (context.Context, error) {
	const op = "middleware.auth.handleUserToken"

	log := s.logger.With(
		slog.String("op", op),
	)

	if t.APIKey == "" || !strings.HasPrefix(t.APIKey, "UserToken ") {
		log.Error("missing token ", errMissingToken)
		return ctx, errMissingToken
	}

	token := strings.TrimPrefix(t.APIKey, "UserToken ")
	scope, err := s.jwtService.VerifyToken(ctx, token)
	if err != nil {
		log.Error("issue verifying jwt token", err)
		return ctx, err
	}

	//TODO: return 403 here if scope != user
	ctx = withScope(ctx, scope)
	return ctx, nil
}

func ScopeFromContext(ctx context.Context) string {
	if scope, ok := ctx.Value(keyScope).(string); ok {
		return scope
	}

	return ""
}

func withScope(ctx context.Context, scope string) context.Context {
	return context.WithValue(ctx, keyScope, scope)
}
