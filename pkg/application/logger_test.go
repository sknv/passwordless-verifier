package application

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sknv/passwordless-verifier/pkg/log"
)

func TestApplication_RegisterLogger(t *testing.T) {
	type args struct {
		config log.Config
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "when an empty config is provided it registers logger successfully",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			a := &Application{
				ctx: context.Background(),
			}
			a.RegisterLogger(tt.args.config)
			assert.Equalf(t, log.L(), log.Extract(a.Context()), "RegisterLogger(%v)", tt.args.config)
		})
	}
}
