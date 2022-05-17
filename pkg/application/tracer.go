package application

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"

	"github.com/sknv/passwordless-verifier/pkg/log"
	"github.com/sknv/passwordless-verifier/pkg/tracing"
)

func (a *Application) RegisterTracing(config tracing.Config) error {
	if err := tracing.Init(config); err != nil {
		return err
	}

	// Remember to flush tracer, if any
	a.closers.Add(func(closeCtx context.Context) error {
		logger := log.Extract(closeCtx)
		logger.Info("flushing tracer...")
		defer logger.Info("tracer flushed")

		provider, _ := otel.GetTracerProvider().(*tracesdk.TracerProvider)
		if err := provider.Shutdown(closeCtx); err != nil {
			return fmt.Errorf("shutdown tracer: %w", err)
		}

		return nil
	})

	return nil
}
