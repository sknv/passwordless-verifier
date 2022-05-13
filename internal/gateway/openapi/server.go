package openapi

import (
	"context"

	"github.com/labstack/echo/v4"

	"github.com/sknv/passwordless-verifier/api/openapi"
	"github.com/sknv/passwordless-verifier/internal/model"
	"github.com/sknv/passwordless-verifier/internal/usecase"
)

const _apiPrefix = "/api"

type Usecase interface {
	CreateVerification(ctx context.Context, newVerification *usecase.NewVerification) (*model.Verification, error)
	GetVerification(ctx context.Context, params *usecase.GetVerificationParams) (*model.Verification, error)
}

type Server struct {
	Usecase Usecase
}

func (s *Server) Route(e *echo.Echo) {
	api := e.Group(_apiPrefix)
	openapi.RegisterHandlers(api, s)
}
