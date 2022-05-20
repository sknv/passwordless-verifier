package usecase

import (
	"context"

	"github.com/google/uuid"

	"github.com/sknv/passwordless-verifier/internal/model"
)

const fieldParams = "params"

type Config struct {
	Deeplink string
}

type Store interface {
	FindVerificationByID(ctx context.Context, id uuid.UUID) (*model.Verification, error)
	FindVerificationByIDWithSession(ctx context.Context, id uuid.UUID) (*model.Verification, error)
	FindLatestVerificationByChatID(ctx context.Context, chatID int64) (*model.Verification, error)
	CreateVerification(ctx context.Context, verification *model.Verification) error
	UpdateVerification(ctx context.Context, verification *model.Verification) error
	UpdateVerificationAndCreateSession(ctx context.Context, verification *model.Verification) error
}

type Usecase struct {
	Config Config
	Store  Store
}
