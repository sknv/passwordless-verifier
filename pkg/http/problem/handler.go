package problem

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/sknv/passwordless-verifier/pkg/log"
)

// HTTPErrorHandler renders echo JSON error response.
func HTTPErrorHandler(respErr error, c echo.Context) {
	// Log the error first
	log.Extract(c.Request().Context()).
		WithError(respErr).
		Error("http error")

	// Committed response means it was already written to the output
	if c.Response().Committed {
		return
	}

	var (
		problem Problem
		httpErr *echo.HTTPError
	)
	switch {
	case errors.As(respErr, &problem):
		// do nothing
	case errors.As(respErr, &httpErr):
		// Not found error represents a request to an unknown resource
		if httpErr.Code == http.StatusNotFound || httpErr.Code == http.StatusMethodNotAllowed {
			problem = NotFound()
			break
		}

		// Transform to a common problem
		problem = New(httpErr.Code, http.StatusText(httpErr.Code))
		if strMsg, ok := httpErr.Message.(string); ok && strMsg != problem.Title { // do not repeat the same info
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
