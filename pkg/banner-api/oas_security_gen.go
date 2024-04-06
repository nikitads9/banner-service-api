// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/ogenerrors"
)

// SecurityHandler is handler for security parameters.
type SecurityHandler interface {
	// HandleAdminToken handles AdminToken security.
	// Токен админа с префиксом `AdminToken`, пр. "AdminToken MYTOKEN".
	HandleAdminToken(ctx context.Context, operationName string, t AdminToken) (context.Context, error)
	// HandleUserToken handles UserToken security.
	// Токен пользователя с префиксом `UserToken`, пр. "UserToken MYTOKEN".
	HandleUserToken(ctx context.Context, operationName string, t UserToken) (context.Context, error)
}

func findAuthorization(h http.Header, prefix string) (string, bool) {
	v, ok := h["Authorization"]
	if !ok {
		return "", false
	}
	for _, vv := range v {
		scheme, value, ok := strings.Cut(vv, " ")
		if !ok || !strings.EqualFold(scheme, prefix) {
			continue
		}
		return value, true
	}
	return "", false
}

func (s *Server) securityAdminToken(ctx context.Context, operationName string, req *http.Request) (context.Context, bool, error) {
	var t AdminToken
	const parameterName = "Token"
	value := req.Header.Get(parameterName)
	if value == "" {
		return ctx, false, nil
	}
	t.APIKey = value
	rctx, err := s.sec.HandleAdminToken(ctx, operationName, t)
	if errors.Is(err, ogenerrors.ErrSkipServerSecurity) {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}
	return rctx, true, err
}
func (s *Server) securityUserToken(ctx context.Context, operationName string, req *http.Request) (context.Context, bool, error) {
	var t UserToken
	const parameterName = "Token"
	value := req.Header.Get(parameterName)
	if value == "" {
		return ctx, false, nil
	}
	t.APIKey = value
	rctx, err := s.sec.HandleUserToken(ctx, operationName, t)
	if errors.Is(err, ogenerrors.ErrSkipServerSecurity) {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}
	return rctx, true, err
}

// SecuritySource is provider of security values (tokens, passwords, etc.).
type SecuritySource interface {
	// AdminToken provides AdminToken security value.
	// Токен админа с префиксом `AdminToken`, пр. "AdminToken MYTOKEN".
	AdminToken(ctx context.Context, operationName string) (AdminToken, error)
	// UserToken provides UserToken security value.
	// Токен пользователя с префиксом `UserToken`, пр. "UserToken MYTOKEN".
	UserToken(ctx context.Context, operationName string) (UserToken, error)
}

func (s *Client) securityAdminToken(ctx context.Context, operationName string, req *http.Request) error {
	t, err := s.sec.AdminToken(ctx, operationName)
	if err != nil {
		return errors.Wrap(err, "security source \"AdminToken\"")
	}
	req.Header.Set("Token", t.APIKey)
	return nil
}
func (s *Client) securityUserToken(ctx context.Context, operationName string, req *http.Request) error {
	t, err := s.sec.UserToken(ctx, operationName)
	if err != nil {
		return errors.Wrap(err, "security source \"UserToken\"")
	}
	req.Header.Set("Token", t.APIKey)
	return nil
}
