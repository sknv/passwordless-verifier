package openapi

import (
	"github.com/labstack/echo/v4"
)

const _apiPrefix = "/api"

type Server struct{}

func (s *Server) Route(e *echo.Echo) {
	api := e.Group(_apiPrefix)
	api.POST("/verifications", s.CreateVerification)
}
