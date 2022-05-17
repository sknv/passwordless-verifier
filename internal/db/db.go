package db

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
)

type DB struct {
	DB *bun.DB
}

func (d *DB) Create(ctx context.Context, model any) (sql.Result, error) {
	return d.DB.NewInsert().
		Model(model).
		Exec(ctx)
}

func (d *DB) Find(ctx context.Context, dest any, where string, args ...any) error {
	return d.DB.NewSelect().
		Model(dest).
		Where(where, args...).
		Scan(ctx)
}

func (d *DB) Update(ctx context.Context, model any, columns ...string) (sql.Result, error) {
	return d.DB.NewUpdate().
		Model(model).
		Column(columns...).
		WherePK().
		Exec(ctx)
}
