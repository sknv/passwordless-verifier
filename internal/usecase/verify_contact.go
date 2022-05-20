package usecase

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"

	"github.com/sknv/passwordless-verifier/internal/model"
	"github.com/sknv/passwordless-verifier/pkg/log"
)

type VerifyContactParams struct {
	ChatID      int64
	ContactID   int64
	PhoneNumber string
}

func (p VerifyContactParams) Validate() error {
	if p.ChatID != p.ContactID {
		return ErrWrongContact
	}

	return nil
}

func (u *Usecase) VerifyContact(ctx context.Context, params *VerifyContactParams) (*model.Verification, error) {
	ctx, span := otel.Tracer("").Start(ctx, "usecase.VerifyContact")
	defer span.End()

	log.Extract(ctx).
		WithField(fieldParams, params).
		Info("usecase.VerifyContact")

	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("validate params: %w", err)
	}

	verification, err := u.Store.FindLatestVerificationByChatID(ctx, params.ChatID)
	if err != nil {
		return nil, fmt.Errorf("find latest verification by chat id: %w", ConvertStoreError(err))
	}

	verification.LogIn(params.PhoneNumber)
	if err = u.Store.UpdateVerificationAndCreateSession(ctx, verification); err != nil {
		return nil, fmt.Errorf("update verification and create session: %w", ConvertStoreError(err))
	}

	return verification, nil
}
