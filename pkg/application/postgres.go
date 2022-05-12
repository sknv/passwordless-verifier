package application

import (
	"context"
	"fmt"

	"github.com/uptrace/bun"

	"github.com/sknv/passwordless-verifier/pkg/closer"
	"github.com/sknv/passwordless-verifier/pkg/log"
	"github.com/sknv/passwordless-verifier/pkg/postgres"
)

func (a *Application) RegisterPostgres(
	ctx context.Context, config postgres.Config, options ...postgres.Option,
) (*bun.DB, error) {
	logger := log.Extract(ctx)
	logger.Info("opening postgres connection...")

	pg := postgres.Connect(config, options...)

	logger.Info("postgres connection opened")

	// Ping the db
	logger.Info("checking postgres...")

	if err := pg.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping postgres: %w", err)
	}

	logger.Info("postgres checked")

	// Remember to close the db connection
	a.Closers.Add(func(context.Context) error {
		logger.Info("closing postgres connection...")
		defer logger.Info("postgres connection closed")

		if err := closer.CloseWithContext(ctx, pg.Close); err != nil {
			return fmt.Errorf("close postrgres: %w", err)
		}

		return nil
	})

	return pg, nil
}
