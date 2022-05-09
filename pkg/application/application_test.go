package application

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplication_Context(t *testing.T) {
	type fields struct {
		ctx context.Context
	}

	tests := []struct {
		name   string
		fields fields
		want   context.Context
	}{
		{
			name: "it returns the provided ctx field",
			fields: fields{
				ctx: context.Background(),
			},
			want: context.Background(),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			a := &Application{
				ctx: tt.fields.ctx,
			}
			assert.Equalf(t, tt.want, a.Context(), "Context()")
		})
	}
}

func TestNewApplication(t *testing.T) {
	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name string
		args args
		want *Application
	}{
		{
			name: "it returns the expected application",
			args: args{
				ctx: context.Background(),
			},
			want: &Application{
				ctx: context.Background(),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equalf(t, tt.want, NewApplication(tt.args.ctx), "NewApplication(%v)", tt.args.ctx)
		})
	}
}
