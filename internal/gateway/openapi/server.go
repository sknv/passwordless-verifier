package openapi

import (
	"github.com/labstack/echo/v4"

	"github.com/sknv/passwordless-verifier/api/openapi"
)

const _apiPrefix = "/api"

type Server struct{}

func (s *Server) Route(e *echo.Echo) {
	api := e.Group(_apiPrefix)
	openapi.RegisterHandlers(api, s)
}
