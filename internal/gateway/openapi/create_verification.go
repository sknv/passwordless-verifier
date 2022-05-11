package openapi

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"github.com/sknv/passwordless-verifier/api/openapi"
	"github.com/sknv/passwordless-verifier/internal/gateway/openapi/view"
	"github.com/sknv/passwordless-verifier/pkg/log"
)

func (s *Server) CreateVerification(c echo.Context) error {
	ctx := c.Request().Context()

	req := &openapi.NewVerification{}
	if err := c.Bind(req); err != nil {
		return err
	}

	log.AddFields(ctx, logrus.Fields{fieldRequest: req})

	return c.JSON(http.StatusOK, view.ToVerification())
}
