package converter

import (
	openapiTypes "github.com/deepmap/oapi-codegen/pkg/types"

	"github.com/sknv/passwordless-verifier/api/openapi"
	"github.com/sknv/passwordless-verifier/internal/model"
	"github.com/sknv/passwordless-verifier/internal/usecase"
)

var verificationStatuses = map[model.VerificationStatus]openapi.VerificationStatus{
	model.VerificationStatusInProgress: openapi.VerificationStatusInProgress,
	model.VerificationStatusCompleted:  openapi.VerificationStatusCompleted,
}

func FromNewVerification(newVerification *openapi.NewVerification) *usecase.NewVerification {
	return &usecase.NewVerification{
		Method: model.VerificationMethod(newVerification.Method),
	}
}

func ToVerification(verification *model.Verification) *openapi.Verification {
	return &openapi.Verification{
		Id:     openapiTypes.UUID(verification.ID.String()),
		Method: openapi.VerificationMethod(verification.Method),
		Status: verificationStatuses[verification.Status],
	}
}
