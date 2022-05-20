package store

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"

	"github.com/sknv/passwordless-verifier/internal/model"
)

func (d *DB) FindVerificationByID(ctx context.Context, id uuid.UUID) (*model.Verification, error) {
	verification := &model.Verification{}
	err := d.DB.NewSelect().
		Model(verification).
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	return verification, nil
}

func (d *DB) FindVerificationByIDWithSession(ctx context.Context, id uuid.UUID) (*model.Verification, error) {
	verification, err := d.FindVerificationByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("select verification: %w", err)
	}

	// Load a session
	verification.Session, err = d.findLatestSessionByVerificationID(ctx, verification.ID)
	if err != nil {
		return nil, fmt.Errorf("select session: %w", err)
	}

	return verification, nil
}

func (d *DB) FindLatestVerificationByChatID(ctx context.Context, chatID int64) (*model.Verification, error) {
	verification := &model.Verification{}
	err := d.DB.NewSelect().
		Model(verification).
		Where("chat_id = ?", chatID).
		Order("created_at DESC").
		Limit(1).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	return verification, nil
}

func (d *DB) CreateVerification(ctx context.Context, verification *model.Verification) error {
	verification.CreatedAt = time.Now()

	_, err := d.DB.NewInsert().
		Model(verification).
		Exec(ctx)
	return err
}

func (d *DB) UpdateVerification(ctx context.Context, verification *model.Verification) error {
	return updateVerification(ctx, d.DB, verification)
}

func (d *DB) UpdateVerificationAndCreateSession(ctx context.Context, verification *model.Verification) error {
	// Apply changes in transaction
	return d.DB.RunInTx(ctx, nil, func(txCtx context.Context, tx bun.Tx) error {
		if err := updateVerification(txCtx, tx, verification); err != nil {
			return fmt.Errorf("update verification: %w", err)
		}

		if err := createSession(txCtx, tx, verification.Session); err != nil {
			return fmt.Errorf("create session: %w", err)
		}

		return nil
	})
}

func updateVerification(ctx context.Context, db bun.IDB, verification *model.Verification) error {
	verification.UpdatedAt = time.Now()

	_, err := db.NewUpdate().
		Model(verification).
		Column("status", "chat_id", "updated_at").
		WherePK().
		Exec(ctx)
	return err
}
