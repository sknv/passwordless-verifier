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

type GetVerificationParams struct {
	ID string
}

func (p GetVerificationParams) Validate() error {
	if _, err := uuid.Parse(p.ID); err != nil {
		return problem.BadRequest(problem.InvalidParam{
			Name:    "id",
			Message: err.Error(),
		})
	}

	return nil
}

func (p GetVerificationParams) TypedID() uuid.UUID {
	id, _ := uuid.Parse(p.ID)
	return id
}

func (u *Usecase) GetVerification(ctx context.Context, params *GetVerificationParams) (*model.Verification, error) {
	ctx, span := otel.Tracer("").Start(ctx, "usecase.GetVerification")
	defer span.End()

	log.Extract(ctx).
		WithField(fieldParams, params).
		Info("usecase.GetVerification")

	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("validate params: %w", err)
	}

	verifications := &model.Verifications{DB: u.DB}
	verification, err := verifications.FindByID(ctx, params.TypedID())
	if err != nil {
		return nil, fmt.Errorf("find verification by id: %w", err)
	}

	return verification, nil
}
