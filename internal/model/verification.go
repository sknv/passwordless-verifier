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
	DB            *bun.DB `bun:"-"`

	ID        uuid.UUID          `bun:"id,nullzero"`
	Method    VerificationMethod `bun:"method"`
	Status    VerificationStatus `bun:"status"`
	Deeplink  string             `bun:"deeplink"`
	CreatedAt time.Time          `bun:"created_at,nullzero"`
	UpdatedAt time.Time          `bun:"updated_at,nullzero"`
}

func (v *Verification) Create(ctx context.Context) error {
	_, err := v.DB.NewInsert().
		Model(v).
		Exec(ctx)
	return err
}
