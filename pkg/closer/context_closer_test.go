package closer

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCloseWithContext(t *testing.T) {
	doneCtx, cancel := context.WithTimeout(context.Background(), -1)
	defer cancel()

	tests := []struct {
		name    string
		ctx     context.Context
		closer  PlainCloser
		wantErr bool
	}{
		{
			name:    "it ignores empty closer",
			ctx:     context.Background(),
			closer:  nil,
			wantErr: false,
		},
		{
			name:    "it closes successfully",
			ctx:     context.Background(),
			closer:  func() error { return nil },
			wantErr: false,
		},
		{
			name:    "when closer returns an error it returns an error",
			ctx:     context.Background(),
			closer:  func() error { return errors.New("any") },
			wantErr: true,
		},
		{
			name:    "for a context timeout it returns an error",
			ctx:     doneCtx,
			closer:  func() error { return nil },
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := CloseWithContext(tt.ctx, tt.closer)
			assert.Equalf(t, tt.wantErr, err != nil, "errors do not match, err=%s", err)
		})
	}
}
