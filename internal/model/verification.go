package model

import (
	"context"
	"database/sql"
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

	ID        uuid.UUID          `bun:"id,pk,nullzero"`
	Method    VerificationMethod `bun:"method"`
	Status    VerificationStatus `bun:"status"`
	Deeplink  string             `bun:"deeplink"`
	ChatID    sql.NullInt64      `bun:"chat_id"`
	CreatedAt time.Time          `bun:"created_at,nullzero"`
	UpdatedAt time.Time          `bun:"updated_at,nullzero"`
}

func NewVerification(db DB) *Verification {
	return &Verification{
		DB: db,

		ID: uuid.New(),
	}
}

func (v *Verification) SetChatID(chatID int64) {
	v.ChatID.Int64, v.ChatID.Valid = chatID, true
}

func (v *Verification) Create(ctx context.Context, deeplink string) error {
	v.Deeplink = v.formatDeeplink(deeplink)
	v.Status = VerificationStatusInProgress

	_, err := v.DB.Create(ctx, v)
	return err
}

func (v *Verification) Update(ctx context.Context) error {
	v.UpdatedAt = time.Now()

	_, err := v.DB.Update(ctx, v, "status", "chat_id", "updated_at")
	return err
}

func (v *Verification) formatDeeplink(format string) string {
	return fmt.Sprintf("%s?start=%s", format, v.ID) // formatted deeplink should base on verification method
}
