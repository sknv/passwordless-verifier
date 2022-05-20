package store

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"

	"github.com/sknv/passwordless-verifier/internal/model"
)

func (d *DB) findLatestSessionByVerificationID(ctx context.Context, verificationID uuid.UUID) (*model.Session, error) {
	session := &model.Session{}
	err := d.DB.NewSelect().
		Model(session).
		Where("verification_id = ?", verificationID).
		Order("created_at DESC").
		Limit(1).
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) { // session may not exist
			return nil, nil
		}
		return nil, err
	}

	return session, nil
}

func createSession(ctx context.Context, db bun.IDB, session *model.Session) error {
	session.CreatedAt = time.Now()

	_, err := db.NewInsert().
		Model(session).
		Exec(ctx)
	return err
}
