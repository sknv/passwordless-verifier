package closer

import (
	"context"
	"fmt"
)

// PlainCloser describes a closer without using context.
type PlainCloser func() error

// CloseWithContext tries to apply the plain closer gracefully using the provided context.
func CloseWithContext(ctx context.Context, closer PlainCloser) error {
	if closer == nil {
		return nil
	}

	var err error
	done := make(chan struct{})
	go func() {
		err = closer()
		close(done)
	}()

	select {
	case <-done:
		return err
	case <-ctx.Done():
		return fmt.Errorf("context done: %w", ctx.Err())
	}
}
