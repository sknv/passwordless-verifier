package tracing

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTraceID(t *testing.T) {
	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name      string
		args      args
		wantValid bool
	}{
		{
			name: "when an empty context is provided it returns an invalid flaf",
			args: args{
				ctx: context.Background(),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			_, valid := GetTraceID(tt.args.ctx)
			assert.Equalf(t, tt.wantValid, valid, "GetTraceID(%v)", tt.args.ctx)
		})
	}
}
