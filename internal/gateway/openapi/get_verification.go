package openapi

import (
	"net/http"

	openapiTypes "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/labstack/echo/v4"

	"github.com/sknv/passwordless-verifier/internal/gateway/openapi/converter"
	"github.com/sknv/passwordless-verifier/internal/usecase"
)

// GetVerification handler
// GET /verifications/{id}
func (s *Server) GetVerification(c echo.Context, id openapiTypes.UUID) error {
	ctx := c.Request().Context()

	verification, err := s.Usecase.GetVerification(ctx, &usecase.GetVerificationParams{ID: string(id)})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, converter.ToVerification(verification))
}
