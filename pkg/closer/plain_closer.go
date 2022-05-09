package closer

import (
	"context"
	"fmt"
)

// PlainCloser describes a closer without using context.
type PlainCloser func() error

// CloseContext tries to apply the plain closer gracefully using the provided context.
func CloseContext(ctx context.Context, closer PlainCloser) error {
	if closer == nil {
		return nil
	}

	var err error
	closed := make(chan struct{})
	go func() {
		err = closer()
		close(closed)
	}()

	select {
	case <-closed:
		return err
	case <-ctx.Done():
		return fmt.Errorf("context done: %w", ctx.Err())
	}
}
