package model

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type VerificationMethod string

const (
	VerificationMethodTelegram VerificationMethod = "telegram"
)

func (m VerificationMethod) Validate() error {
	if m == VerificationMethodTelegram {
		return nil
	}

	return fmt.Errorf("unknown verification method: %s", m)
}

type VerificationStatus string

const (
	VerificationStatusInProgress VerificationStatus = "in_progress"
	VerificationStatusCompleted  VerificationStatus = "completed"
)

type Verification struct {
	bun.BaseModel `bun:"table:verifications"`
	DB            DB `bun:"-"`

	ID        uuid.UUID          `bun:"id,nullzero"`
	Method    VerificationMethod `bun:"method"`
	Status    VerificationStatus `bun:"status"`
	Deeplink  string             `bun:"deeplink"`
	CreatedAt time.Time          `bun:"created_at,nullzero"`
	UpdatedAt time.Time          `bun:"updated_at,nullzero"`
}

func NewVerification(db DB) *Verification {
	return &Verification{
		DB: db,

		ID: uuid.New(),
	}
}

func (v *Verification) Create(ctx context.Context, deeplinkFormat string) error {
	v.Deeplink = v.formatDeeplink(deeplinkFormat)
	v.Status = VerificationStatusInProgress

	_, err := v.DB.Create(ctx, v)
	return err
}

func (v *Verification) formatDeeplink(format string) string {
	return fmt.Sprintf(format, v.ID)
}
