package model

import (
	"context"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Verifications struct {
	DB *bun.DB
}

func (v *Verifications) FindByID(ctx context.Context, id uuid.UUID) (*Verification, error) {
	verification := &Verification{}
	err := v.DB.NewSelect().
		Model(verification).
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		return nil, ConvertDBError(err)
	}

	verification.DB = v.DB
	return verification, nil
}
