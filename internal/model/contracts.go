package model

import (
	"context"
	"database/sql"
)

type DB interface {
	Create(ctx context.Context, model any) (sql.Result, error)
	Find(ctx context.Context, dest any, where string, args ...any) error
}
