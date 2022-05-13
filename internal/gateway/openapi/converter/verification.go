package converter

import (
	openapiTypes "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/google/uuid"

	"github.com/sknv/passwordless-verifier/api/openapi"
	"github.com/sknv/passwordless-verifier/internal/model"
	"github.com/sknv/passwordless-verifier/internal/usecase"
)

func ToVerification(verification *model.Verification) *openapi.Verification {
	return &openapi.Verification{
		Id: openapiTypes.UUID(uuid.NewString()),
	}
}

func FromNewVerification(newVerification *openapi.NewVerification) *usecase.NewVerification {
	return &usecase.NewVerification{
		Method: model.VerificationMethod(newVerification.Method),
	}
}
