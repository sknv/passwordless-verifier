package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	glog "github.com/labstack/gommon/log"

	"github.com/sknv/passwordless-verifier/pkg/http/health"
	"github.com/sknv/passwordless-verifier/pkg/http/logger"
	"github.com/sknv/passwordless-verifier/pkg/http/metric"
	"github.com/sknv/passwordless-verifier/pkg/http/problem"
	"github.com/sknv/passwordless-verifier/pkg/http/render"
	pkgmware "github.com/sknv/passwordless-verifier/pkg/http/server/middleware"
	"github.com/sknv/passwordless-verifier/pkg/http/tracing"
	"github.com/sknv/passwordless-verifier/pkg/log"
)

// Option configures echo instance.
type Option func(*echo.Echo)

type (
	MetricConfig struct {
		Namespace string
		Skipper   middleware.Skipper
	}

	TracingConfig struct {
		Namespace string
	}

	Config struct {
		Metric  MetricConfig
		Tracing TracingConfig
	}
)

// New creates an echo.Echo instance.
func New(config Config, options ...Option) *echo.Echo {
	e := echo.New()
	e.HideBanner, e.HidePort = true, true
	logger.WithLogger(e)
	render.WithJSONSerializer(e)
	problem.WithHTTPErrorHandler(e)
	health.WithHealthCheck(e)

	// Prepend default middleware, order matters
	e.Use(
		middleware.RequestID(),
		pkgmware.WithContextLogger(log.L()),
		pkgmware.WithLogRequestID,
		pkgmware.Logger(),
	)
	metric.WithPrometheus(e, config.Metric.Namespace, config.Metric.Skipper)
	tracing.WithTracer(e, config.Tracing.Namespace)

	// Apply options
	for _, opt := range options {
		opt(e)
	}

	// Append default middleware
	WithRecover(e)

	return e
}

// WithRecover is a recover middleware.
func WithRecover(e *echo.Echo) {
	recoverConfig := middleware.DefaultRecoverConfig
	recoverConfig.LogLevel = glog.ERROR
	e.Use(middleware.RecoverWithConfig(recoverConfig))
}
