package model

import (
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

func FormatDeeplink(deeplink string, verificationID uuid.UUID) string {
	return fmt.Sprintf("%s?start=%s", deeplink, verificationID)
}

type Verification struct {
	bun.BaseModel `bun:"table:verifications"`

	ID        uuid.UUID          `bun:"id,pk,nullzero"`
	Method    VerificationMethod `bun:"method"`
	Status    VerificationStatus `bun:"status"`
	Deeplink  string             `bun:"deeplink"`
	ChatID    int64              `bun:"chat_id,nullzero"`
	CreatedAt time.Time          `bun:"created_at,nullzero"`
	UpdatedAt time.Time          `bun:"updated_at,nullzero"`

	// Relations
	Session *Session `bun:"rel:has-one"`
}

func (v *Verification) LogIn(phoneNumber string) {
	v.Status = VerificationStatusCompleted
	v.Session = &Session{
		ID:             uuid.New(),
		VerificationID: v.ID,
		PhoneNumber:    phoneNumber,
	}
}
