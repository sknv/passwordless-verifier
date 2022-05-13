package openapi

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/sknv/passwordless-verifier/api/openapi"
	"github.com/sknv/passwordless-verifier/internal/gateway/openapi/converter"
)

func (s *Server) CreateVerification(c echo.Context) error {
	ctx := c.Request().Context()

	req := &openapi.NewVerification{}
	if err := c.Bind(req); err != nil {
		return err
	}

	verification, err := s.Usecase.CreateVerification(ctx, converter.FromNewVerification(req))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, converter.ToVerification(verification))
}
