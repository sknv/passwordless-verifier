package tracing

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"

	"github.com/sknv/passwordless-verifier/pkg/log"
	"github.com/sknv/passwordless-verifier/pkg/tracing"
)

var Middleware = otelecho.Middleware

const _fieldTraceID = "trace_id"

type Middlewarer interface {
	Use(middleware ...echo.MiddlewareFunc)
}

// WithTraceID middleware injects a trace id into the context of each request.
func WithTraceID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		traceID, ok := tracing.GetTraceID(ctx)
		if ok {
			log.AddFields(ctx, logrus.Fields{_fieldTraceID: traceID})
		}

		return next(c)
	}
}

// WithTracer middleware enables tracing and injects a trace id into the context of each request.
func WithTracer(m Middlewarer, service string, options ...otelecho.Option) {
	m.Use(Middleware(service, options...), WithTraceID)
}
