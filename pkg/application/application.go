package application

import (
	"context"

	"github.com/sknv/passwordless-verifier/pkg/closer"
	"github.com/sknv/passwordless-verifier/pkg/log"
)

// Application is a core object.
type Application struct {
	Closers *closer.Closers

	ctx        context.Context
	httpServer *preparedHTTPServer
}

func NewApplication(ctx context.Context) *Application {
	return &Application{
		Closers: &closer.Closers{},

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

	a.runHTTPServer(ctx, a.httpServer)

	logger.Info("application started")
	return nil
}

// Stop the application.
func (a *Application) Stop(ctx context.Context) error {
	logger := log.Extract(ctx)
	logger.Info("stopping application...")

	if err := a.Closers.Close(ctx); err != nil {
		return err
	}

	logger.Info("application stopped")
	return nil
}
