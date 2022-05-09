package application

import (
	"context"

	"github.com/sknv/passwordless-verifier/pkg/log"
)

// Application is a core object.
type Application struct {
	ctx context.Context
}

func NewApplication(ctx context.Context) *Application {
	return &Application{
		ctx: ctx,
	}
}

// Context returns the application context.
func (a *Application) Context() context.Context {
	return a.ctx
}

// Run the application.
func (a *Application) Run(ctx context.Context) error {
	logger := log.Extract(ctx)
	logger.Info("starting application...")

	logger.Info("application started")
	return nil
}

// Stop the application.
func (a *Application) Stop(ctx context.Context) error {
	logger := log.Extract(ctx)
	logger.Info("stopping application...")

	logger.Info("application stopped")
	return nil
}
