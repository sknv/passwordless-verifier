package health

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const _healthCheckPath = "/health"

type Router interface {
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// Handler handles an echo server health check.
func Handler(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

// WithHealthCheck enables health checks for an echo server.
func WithHealthCheck(r Router) {
	r.GET(_healthCheckPath, Handler)
}
