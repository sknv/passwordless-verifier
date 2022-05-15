package application

import (
	"context"

	"golang.org/x/sync/errgroup"

	"github.com/sknv/passwordless-verifier/pkg/closer"
	"github.com/sknv/passwordless-verifier/pkg/log"
)

// Application is a core object.
type Application struct {
	ctx        context.Context
	closers    *closer.Closers
	httpServer *preparedHTTPServer
	worker     Worker
}

func NewApplication(ctx context.Context) *Application {
	return &Application{
		ctx:     ctx,
		closers: &closer.Closers{},
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

	if err := runParallel(
		ctx,
		a.runHTTPServer,
		a.runWorker,
	); err != nil {
		return err
	}

	logger.Info("application started")
	return nil
}

// Stop the application.
func (a *Application) Stop(ctx context.Context) error {
	logger := log.Extract(ctx)
	logger.Info("stopping application...")

	if err := a.closers.Close(ctx); err != nil {
		return err
	}

	logger.Info("application stopped")
	return nil
}

func runParallel(ctx context.Context, fns ...func(context.Context)) error {
	group, groupCtx := errgroup.WithContext(ctx)
	for _, fn := range fns {
		fn := fn // remember to copy
		group.Go(func() error {
			fn(groupCtx)
			return nil
		})
	}

	return group.Wait()
}
