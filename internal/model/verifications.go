package model

import (
	"context"

	"github.com/google/uuid"
)

type Verifications struct {
	DB DB
}

func (v *Verifications) FindByID(ctx context.Context, id uuid.UUID) (*Verification, error) {
	verification := &Verification{}
	if err := v.DB.Find(ctx, verification, "id = ?", id); err != nil {
		return nil, ConvertDBError(err)
	}

	verification.DB = v.DB
	return verification, nil
}
