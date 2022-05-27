package converter

import (
	"github.com/sknv/passwordless-verifier/api/openapi"
	"github.com/sknv/passwordless-verifier/internal/model"
	"github.com/sknv/passwordless-verifier/internal/usecase"
)

var verificationStatuses = map[model.VerificationStatus]openapi.VerificationStatus{
	model.VerificationStatusInProgress: openapi.InProgress,
	model.VerificationStatusCompleted:  openapi.Completed,
}

func FromNewVerification(newVerification *openapi.NewVerification) *usecase.NewVerification {
	return &usecase.NewVerification{
		Method: model.VerificationMethod(newVerification.Method),
	}
}

func ToVerification(verification *model.Verification) *openapi.Verification {
	return &openapi.Verification{
		Id:        verification.ID,
		Method:    openapi.VerificationMethod(verification.Method),
		Deeplink:  verification.Deeplink,
		Status:    verificationStatuses[verification.Status],
		CreatedAt: verification.CreatedAt,
		Session:   ToSession(verification.Session),
	}
}
