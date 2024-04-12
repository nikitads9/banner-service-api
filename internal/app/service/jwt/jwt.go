package jwt

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"time"

	"github.com/nikitads9/banner-service-api/internal/logger/sl"

	"github.com/golang-jwt/jwt/v5"
)

// Service is an interface that represents all the capabilities for the JWT service.
type Service interface {
	GenerateToken(ctx context.Context, scope string) (string, error)
	VerifyToken(ctx context.Context, token string) (string, error)
}

type service struct {
	jwtSecret  string
	expiration time.Duration
	log        *slog.Logger
}

// New creates a service with a provided JWT secret string and expiration (hourly) number. It implements
// the JWT Service interface.
func NewJWTService(jwtSecret string, expiration time.Duration, log *slog.Logger) Service {
	return &service{jwtSecret, expiration, log}
}

var (
	ErrUnsupportedSign = errors.New("unexpected signing method")
	ErrNoScope         = errors.New("scope not set")
	ErrInvalidToken    = errors.New("invalid token")

	ErrParseScope = errors.New("parsing scope failed")
	ErrParseExp   = errors.New("parsing token expiration failed")
)

// GenerateToken takes a user ID and
func (s *service) GenerateToken(ctx context.Context, scope string) (string, error) {
	const op = "service.jwt.GenerateToken"

	log := s.log.With(
		slog.String("op", op),
	)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"scope": scope,
		"exp":   time.Now().Add(s.expiration).Unix(),
	})

	log.Info("token generated", slog.String("scope:", scope))

	return token.SignedString([]byte(s.jwtSecret))
}

// VerifyToken parses and validates a jwt token. It returns the scope if the token is valid.
func (s *service) VerifyToken(ctx context.Context, tokenString string) (string, error) {
	const op = "service.jwt.VerifyToken"

	log := s.log.With(
		slog.String("op", op),
	)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Error("unexpected signing method: ", token.Header["alg"])
			return nil, ErrUnsupportedSign
		}
		return []byte(s.jwtSecret), nil
	}, jwt.WithJSONNumber())

	if err != nil {
		log.Error("parsing token failed: ", sl.Err(err))
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !token.Valid || !ok {
		log.Error("invalid token", sl.Err(ErrInvalidToken))
		return "", ErrInvalidToken
	}

	scope := claims["scope"]
	userScope, ok := scope.(string)
	if !ok {
		log.Error("issue parsing user scope", sl.Err(ErrParseScope))
		return "", ErrParseScope

	}

	if userScope == "" {
		log.Error("empty user scope", sl.Err(ErrNoScope))
		return "", ErrNoScope
	}

	exp, err := claims["exp"].(json.Number).Int64()
	if err != nil {
		log.Error("issue parsing token expiration", sl.Err(err))
		return "", ErrParseExp

	}

	if exp < time.Now().Unix() {
		log.Error("token expired", sl.Err(jwt.ErrTokenExpired))
		return "", jwt.ErrTokenExpired
	}

	return userScope, nil

}
