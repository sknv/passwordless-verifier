package usecase

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel"

	"github.com/sknv/passwordless-verifier/internal/model"
	"github.com/sknv/passwordless-verifier/pkg/http/problem"
	"github.com/sknv/passwordless-verifier/pkg/log"
)

type NewVerification struct {
	Method model.VerificationMethod
}

func (v NewVerification) Validate() error {
	if err := v.Method.Validate(); err != nil {
		badRequest := problem.BadRequest(problem.InvalidParam{
			Name:    "method",
			Message: err.Error(),
		})
		badRequest.Err = err
		return badRequest
	}

	return nil
}

func (u *Usecase) CreateVerification(
	ctx context.Context, newVerification *NewVerification,
) (*model.Verification, error) {
	ctx, span := otel.Tracer("").Start(ctx, "usecase.CreateVerification")
	defer span.End()

	log.Extract(ctx).
		WithField(fieldParams, newVerification).
		Info("usecase.CreateVerification")

	if err := newVerification.Validate(); err != nil {
		return nil, fmt.Errorf("validate params: %w", err)
	}

	verificationID := uuid.New()
	verification := &model.Verification{
		ID:       verificationID,
		Method:   newVerification.Method,
		Deeplink: model.FormatDeeplink(u.Config.Deeplink, verificationID),
		Status:   model.VerificationStatusInProgress,
	}
	if err := u.Store.CreateVerification(ctx, verification); err != nil {
		return nil, fmt.Errorf("create verification: %w", ConvertStoreError(err))
	}

	return verification, nil
}
