package application

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sknv/passwordless-verifier/pkg/closer"
	"github.com/sknv/passwordless-verifier/pkg/tracing"
)

func TestApplication_RegisterTracing(t *testing.T) {
	type args struct {
		config tracing.Config
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "when tracing initialized successfully it returns no error",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			a := &Application{
				ctx:     context.Background(),
				closers: &closer.Closers{},
			}
			err := a.RegisterTracing(tt.args.config)
			assert.Equalf(t, tt.wantErr, err != nil, "RegisterTracing(%v)", tt.args.config)
		})
	}
}
