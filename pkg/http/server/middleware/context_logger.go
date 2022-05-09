package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"github.com/sknv/passwordless-verifier/pkg/log"
)

const _fieldRequestID = "request_id"

// WithContextLogger injects the provided logger into request context.
func WithContextLogger(logger *logrus.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctxLog := log.ToContext(c.Request().Context(), logger)
			c.SetRequest(c.Request().WithContext(ctxLog))

			return next(c)
		}
	}
}

// WithLogRequestID is a middleware that injects a request id into the context of each request.
func WithLogRequestID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		requestID := c.Response().Header().Get(echo.HeaderXRequestID)
		log.AddFields(c.Request().Context(), logrus.Fields{_fieldRequestID: requestID})

		return next(c)
	}
}
