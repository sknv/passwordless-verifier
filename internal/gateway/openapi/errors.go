package openapi

import (
	"github.com/labstack/echo/v4"

	"github.com/sknv/passwordless-verifier/pkg/http/problem"
)

func BindRequestError(err error) error {
	switch typedErr := err.(type) { //nolint:errorlint // no wrapping errors on transport layer allowed by convention
	case *echo.HTTPError:
		prb := problem.BadRequest()
		prb.Detail, _ = typedErr.Message.(string)

		return prb
	default:
		return err
	}
}
