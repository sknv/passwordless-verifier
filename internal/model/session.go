package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Session struct {
	bun.BaseModel `bun:"table:sessions"`

	ID             uuid.UUID `bun:"id,pk,nullzero"`
	VerificationID uuid.UUID `bun:"verification_id"`
	PhoneNumber    string    `bun:"phone_number"`
	CreatedAt      time.Time `bun:"created_at,nullzero"`

	// Relations
	Verification *Verification `bun:"rel:belongs-to"`
}
