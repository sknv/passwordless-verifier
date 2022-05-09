package metric

import (
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var NewPrometheus = prometheus.NewPrometheus

// WithPrometheus enables prometheus metrics for an echo server.
func WithPrometheus(e *echo.Echo, subsystem string, skipper middleware.Skipper, metrics ...[]*prometheus.Metric) {
	prom := NewPrometheus(subsystem, skipper, metrics...)
	prom.Use(e)
}
