package usecase

import (
	"context"
	"fmt"

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

func (v NewVerification) ToVerification(db model.DB) *model.Verification {
	verification := model.NewVerification(db)
	verification.Method = v.Method

	return verification
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

	verification := newVerification.ToVerification(u.DB)
	if err := verification.Create(ctx, u.Config.DeeplinkFormat); err != nil {
		return nil, fmt.Errorf("create verification: %w", err)
	}

	return verification, nil
}
