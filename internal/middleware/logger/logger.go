package logger

import (
	"log/slog"
	"net/http"
	"net/url"
)

type Middleware func(next http.Handler) http.Handler

func Wrap(h http.Handler, middlewares ...Middleware) http.Handler {
	switch len(middlewares) {
	case 0:
		return h
	case 1:
		return middlewares[0](h)
	default:
		for i := len(middlewares) - 1; i >= 0; i-- {
			h = middlewares[i](h)
		}
		return h
	}
}

func LogRequests(find RouteFinder, logger *slog.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//var opName, opID string

			/* 			if route, ok := find(r.Method, r.URL); ok {
				opName = route.Name()
				opID = route.OperationID()
			} */

			log := logger.With(
				slog.String("http_method", r.Method),
				slog.String("path", r.URL.Path),
				slog.String("remote_addr", r.RemoteAddr),
				slog.String("user_agent", r.UserAgent()),
			)

			log.Info("request acquired")

			next.ServeHTTP(w, r)
		})
	}
}

// Server is a generic ogen server type.
type Server[R Route] interface {
	FindPath(method string, u *url.URL) (r R, _ bool)
}

// Route is a generic ogen route type.
type Route interface {
	Name() string
	OperationID() string
	PathPattern() string
}

// RouteFinder finds Route by given URL.
type RouteFinder func(method string, u *url.URL) (Route, bool)

// MakeRouteFinder creates RouteFinder from given server.
func MakeRouteFinder[R Route, S Server[R]](server S) RouteFinder {
	return func(method string, u *url.URL) (Route, bool) {
		return server.FindPath(method, u)
	}
}
