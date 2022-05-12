package postgres

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	type args struct {
		config Config
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "it connects successfully",
			args: args{
				config: Config{
					URL: "postgres://localhost:5432/test?sslmode=disable",
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.NotNilf(t, Connect(tt.args.config), "Connect(%v)", tt.args.config)
		})
	}
}
