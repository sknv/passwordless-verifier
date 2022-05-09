package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"

	"github.com/sknv/passwordless-verifier/pkg/log"
)

// DefaultLoggerConfig is the default Logger middleware config.
var DefaultLoggerConfig = LoggerConfig{
	Skipper: middleware.DefaultSkipper,
}

// LoggerConfig defines the config for Logger middleware.
type LoggerConfig struct {
	// Skipper defines a function to skip middleware.
	Skipper middleware.Skipper
}

// Logger returns a middleware that logs HTTP requests.
func Logger() echo.MiddlewareFunc {
	return LoggerWithConfig(DefaultLoggerConfig)
}

// LoggerWithConfig returns a Logger middleware with config.
func LoggerWithConfig(config LoggerConfig) echo.MiddlewareFunc {
	// Defaults
	if config.Skipper == nil {
		config.Skipper = DefaultLoggerConfig.Skipper
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			req := c.Request()
			res := c.Response()
			start := time.Now()

			var err error
			if err = next(c); err != nil {
				c.Error(err)
			}
			latency := time.Since(start)

			log.Extract(req.Context()).WithFields(logrus.Fields{
				"op":            "http",
				"id":            res.Header().Get(echo.HeaderXRequestID),
				"remote_ip":     c.RealIP(),
				"host":          req.Host,
				"method":        req.Method,
				"uri":           req.RequestURI,
				"user_agent":    req.UserAgent(),
				"status":        res.Status,
				"latency":       latency,
				"latency_human": latency.String(),
				"bytes_in":      req.Header.Get(echo.HeaderContentLength),
				"bytes_out":     res.Size,
			}).Info()

			return err
		}
	}
}
