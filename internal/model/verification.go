package model

import (
	"fmt"

	"github.com/google/uuid"
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
	ID     uuid.UUID
	Method VerificationMethod
	Status VerificationStatus
}
