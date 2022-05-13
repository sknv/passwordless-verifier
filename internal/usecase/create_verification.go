package usecase

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"

	"github.com/sknv/passwordless-verifier/internal/model"
	"github.com/sknv/passwordless-verifier/pkg/log"
)

type NewVerification struct {
	Method model.VerificationMethod
}

func (v NewVerification) Validate() error {
	return v.Method.Validate()
}

func (u *Usecase) CreateVerification(ctx context.Context, newVerification *NewVerification) (*model.Verification, error) {
	ctx, span := otel.Tracer("").Start(ctx, "usecase.CreateVerification")
	defer span.End()

	log.AddFields(ctx, logrus.Fields{fieldParams: newVerification})

	if err := newVerification.Validate(); err != nil {
		return nil, fmt.Errorf("validate params: %w", err)
	}

	verification := &model.Verification{
		ID:     uuid.New(),
		Method: newVerification.Method,
		Status: model.VerificationStatusInProgress,
	}

	return verification, nil
}
