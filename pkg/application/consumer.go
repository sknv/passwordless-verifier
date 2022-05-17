package application

import (
	"context"
	"fmt"

	"github.com/sknv/passwordless-verifier/pkg/log"
)

type Consumer interface {
	Run(ctx context.Context)
	Close(ctx context.Context) error
}

func (a *Application) RegisterConsumer(worker Consumer) {
	a.consumer = worker
}

func (a *Application) runConsumer(ctx context.Context) {
	if a.consumer == nil {
		return // no consumer registered
	}

	logger := log.Extract(ctx)
	logger.Info("starting consumer...")
	defer logger.Info("consumer started")

	go a.consumer.Run(ctx) // start in its own goroutine

	// Remember to stop the consumer
	a.closers.Add(func(closeCtx context.Context) error {
		logger.Info("stopping consumer...")
		defer logger.Info("consumer stopped")

		if err := a.consumer.Close(closeCtx); err != nil {
			return fmt.Errorf("stop consumer: %w", err)
		}

		return nil
	})
}
