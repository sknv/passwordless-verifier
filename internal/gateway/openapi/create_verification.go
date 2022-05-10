package openapi

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"

	"github.com/sknv/passwordless-verifier/api/openapi"
	"github.com/sknv/passwordless-verifier/internal/gateway/openapi/view"
	"github.com/sknv/passwordless-verifier/pkg/log"
)

func (s *Server) CreateVerification(c echo.Context) error {
	ctx, span := otel.Tracer("").Start(c.Request().Context(), "openapi.CreateVerification")
	defer span.End()

	logger := log.Extract(ctx)

	req := &openapi.NewVerification{}
	if err := c.Bind(req); err != nil {
		logger.WithError(err).Error(msgDecodeJSON)
		return BindRequestError(err)
	}

	log.AddFields(ctx, logrus.Fields{fieldRequest: req})

	return c.JSON(http.StatusOK, view.ToVerification())
}
