package postgres

import (
	"context"

	"github.com/uptrace/bun"

	"github.com/sknv/passwordless-verifier/pkg/log"
)

type Logger struct{}

// WithLogger registers a db logger hook.
func WithLogger(db *bun.DB) {
	db.AddQueryHook(Logger{})
}

func (Logger) BeforeQuery(ctx context.Context, _ *bun.QueryEvent) context.Context { return ctx }

func (Logger) AfterQuery(ctx context.Context, event *bun.QueryEvent) {
	log.Extract(ctx).Debug(event)
}
