package usecase

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
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
		return problem.BadRequest(problem.InvalidParam{
			Name:    "method",
			Message: err.Error(),
		})
	}

	return nil
}

func (v NewVerification) ToVerification(db *bun.DB) *model.Verification {
	return &model.Verification{
		DB: db,

		ID:     uuid.New(),
		Method: v.Method,
		Status: model.VerificationStatusInProgress,
	}
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
	if err := verification.Create(ctx); err != nil {
		return nil, fmt.Errorf("create verification: %w", err)
	}

	return verification, nil
}
