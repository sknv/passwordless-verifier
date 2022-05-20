package application

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sknv/passwordless-verifier/pkg/closer"
	"github.com/sknv/passwordless-verifier/pkg/http/server"
)

func TestApplication_RegisterHTTPServer(t *testing.T) {
	type args struct {
		config HTTPServerConfig
		opts   []server.Option
	}

	tests := []struct {
		name        string
		args        args
		wantPresent bool
	}{
		{
			name:        "it registers an http server successfully",
			wantPresent: true,
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
			got := a.RegisterHTTPServer(tt.args.config, tt.args.opts...)
			assert.Equalf(t, tt.wantPresent, got != nil, "RegisterHTTPServer(%v, %v)", tt.args.config, tt.args.opts)
		})
	}
}
