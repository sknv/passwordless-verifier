package usecase

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel"

	"github.com/sknv/passwordless-verifier/pkg/http/problem"
	"github.com/sknv/passwordless-verifier/pkg/log"
)

type SetVerificationChatParams struct {
	ID     string
	ChatID int64
}

func (p SetVerificationChatParams) Validate() error {
	if _, err := uuid.Parse(p.ID); err != nil {
		badRequest := problem.BadRequest(problem.InvalidParam{
			Name:    "id",
			Message: err.Error(),
		})
		badRequest.Err = err
		return badRequest
	}

	return nil
}

func (p SetVerificationChatParams) TypedID() uuid.UUID {
	id, _ := uuid.Parse(p.ID)
	return id
}

func (u *Usecase) SetVerificationChat(ctx context.Context, params *SetVerificationChatParams) error {
	ctx, span := otel.Tracer("").Start(ctx, "usecase.SetVerificationChat")
	defer span.End()

	log.Extract(ctx).
		WithField(fieldParams, params).
		Info("usecase.SetVerificationChat")

	if err := params.Validate(); err != nil {
		return fmt.Errorf("validate params: %w", err)
	}

	verification, err := u.Store.FindVerificationByID(ctx, params.TypedID())
	if err != nil {
		return fmt.Errorf("find verification by id: %w", ConvertStoreError(err))
	}

	verification.ChatID = params.ChatID
	if err = u.Store.UpdateVerification(ctx, verification); err != nil {
		return fmt.Errorf("update verification: %w", ConvertStoreError(err))
	}

	return nil
}
