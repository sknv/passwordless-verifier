package application

import (
	"context"
	"fmt"

	"github.com/sknv/passwordless-verifier/pkg/log"
)

type Worker interface {
	Run(ctx context.Context)
	Close(ctx context.Context) error
}

func (a *Application) RegisterWorker(worker Worker) {
	a.worker = worker
}

func (a *Application) runWorker(ctx context.Context) {
	if a.worker == nil {
		return // no worker registered
	}

	logger := log.Extract(ctx)
	logger.Info("starting worker...")
	defer logger.Info("worker started")

	go a.worker.Run(ctx) // start in its own goroutine

	// Remember to stop the worker
	a.closers.Add(func(closeCtx context.Context) error {
		logger.Info("stopping worker...")
		defer logger.Info("worker stopped")

		if err := a.worker.Close(closeCtx); err != nil {
			return fmt.Errorf("stop worker: %w", err)
		}

		return nil
	})
}
