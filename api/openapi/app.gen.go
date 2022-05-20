// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.10.1 DO NOT EDIT.
package openapi

import (
	"fmt"
	"net/http"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/labstack/echo/v4"
)

// Defines values for VerificationMethod.
const (
	VerificationMethodTelegram VerificationMethod = "telegram"
)

// Defines values for VerificationStatus.
const (
	VerificationStatusCompleted VerificationStatus = "completed"

	VerificationStatusInProgress VerificationStatus = "in_progress"
)

// NewVerification defines model for NewVerification.
type NewVerification struct {
	Method VerificationMethod `json:"method"`
}

// Problem defines model for Problem.
type Problem struct {
	Data   *map[string]interface{} `json:"data,omitempty"`
	Detail *string                 `json:"detail,omitempty"`
	Status int                     `json:"status"`
	Title  string                  `json:"title"`
	Type   string                  `json:"type"`
}

// Session defines model for Session.
type Session struct {
	CreatedAt   time.Time `json:"createdAt"`
	PhoneNumber string    `json:"phoneNumber"`
}

// Verification defines model for Verification.
type Verification struct {
	CreatedAt time.Time          `json:"createdAt"`
	Deeplink  string             `json:"deeplink"`
	Id        openapi_types.UUID `json:"id"`
	Method    VerificationMethod `json:"method"`
	Session   *Session           `json:"session,omitempty"`
	Status    VerificationStatus `json:"status"`
}

// VerificationMethod defines model for VerificationMethod.
type VerificationMethod string

// VerificationStatus defines model for VerificationStatus.
type VerificationStatus string

// CreateVerificationJSONBody defines parameters for CreateVerification.
type CreateVerificationJSONBody NewVerification

// CreateVerificationJSONRequestBody defines body for CreateVerification for application/json ContentType.
type CreateVerificationJSONRequestBody CreateVerificationJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /verifications)
	CreateVerification(ctx echo.Context) error

	// (GET /verifications/{id})
	GetVerification(ctx echo.Context, id openapi_types.UUID) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// CreateVerification converts echo context to params.
func (w *ServerInterfaceWrapper) CreateVerification(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreateVerification(ctx)
	return err
}

// GetVerification converts echo context to params.
func (w *ServerInterfaceWrapper) GetVerification(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id openapi_types.UUID

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetVerification(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/verifications", wrapper.CreateVerification)
	router.GET(baseURL+"/verifications/:id", wrapper.GetVerification)

}
