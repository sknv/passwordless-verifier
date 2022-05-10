package view

import (
	openapiTypes "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/google/uuid"

	"github.com/sknv/passwordless-verifier/api/openapi"
)

func ToVerification() *openapi.Verification {
	return &openapi.Verification{
		Id: openapiTypes.UUID(uuid.NewString()),
	}
}
