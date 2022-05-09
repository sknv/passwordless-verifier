package problem

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/sknv/passwordless-verifier/pkg/log"
)

// HTTPErrorHandler renders echo JSON error response.
func HTTPErrorHandler(respErr error, c echo.Context) {
	// Committed response means it was already written to the output
	if c.Response().Committed {
		return
	}

	var problem *Problem
	switch typedErr := respErr.(type) { //nolint:errorlint // no wrapping errors on transport layer allowed by convention
	case *Problem:
		problem = typedErr
	case *echo.HTTPError:
		// Not found error represents a request to an unknown resource
		if typedErr.Code == http.StatusNotFound || typedErr.Code == http.StatusMethodNotAllowed {
			problem = NotFound()
			break
		}

		// Log the error because it was raised from the framework itself
		log.Extract(c.Request().Context()).
			WithError(respErr).
			Warn("echo http error")

		// Transform to a common problem
		problem = New(typedErr.Code, http.StatusText(typedErr.Code))
		if strMsg, ok := typedErr.Message.(string); ok && strMsg != problem.Title { // do not repeat the same info
			problem.Detail = strMsg
		}
	default:
		problem = InternalServerError()
	}

	// HEAD request expects a response without a body
	if c.Request().Method == http.MethodHead {
		if err := c.NoContent(problem.Status); err != nil {
			log.Extract(c.Request().Context()).
				WithError(err).
				Warn("send no content")
		}
		return
	}

	if err := c.JSON(problem.Status, problem); err != nil {
		log.Extract(c.Request().Context()).
			WithError(err).
			Warn("render json")
	}
}

// WithHTTPErrorHandler sets custom HTTP error handler for the provided echo instance.
func WithHTTPErrorHandler(e *echo.Echo) {
	e.HTTPErrorHandler = HTTPErrorHandler
}
