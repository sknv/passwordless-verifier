package openapi

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"

	"github.com/sknv/passwordless-verifier/api/openapi"
	"github.com/sknv/passwordless-verifier/internal/gateway/openapi/view"
	"github.com/sknv/passwordless-verifier/pkg/http/problem"
	"github.com/sknv/passwordless-verifier/pkg/log"
)

func (s *Server) CreateVerification(c echo.Context) error {
	ctx, span := otel.Tracer("").Start(c.Request().Context(), "openapi.CreateVerification")
	defer span.End()

	logger := log.Extract(ctx)

	req := &openapi.CreateVerificationRequest{}
	if err := c.Bind(req); err != nil {
		logger.WithError(err).Error(msgDecodeJSON)
		return err
	}

	logger = logger.WithField(fieldRequest, req)

	if err := c.JSON(http.StatusOK, view.ToVerification()); err != nil {
		logger.WithError(err).Error(msgRenderJSON)
		return problem.InternalServerError()
	}

	logger.Debug("ok")
	return nil
}
