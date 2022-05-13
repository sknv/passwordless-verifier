package application

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sknv/passwordless-verifier/pkg/closer"
)

func TestApplication_Context(t *testing.T) {
	type fields struct {
		ctx context.Context
	}

	ctx := context.WithValue(context.Background(), "key", "val")

	tests := []struct {
		name   string
		fields fields
		want   context.Context
	}{
		{
			name: "it returns the provided ctx field",
			fields: fields{
				ctx: ctx,
			},
			want: ctx,
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
				Closers: &closer.Closers{},

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

func TestApplication_Run(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name: "when no component is registered it runs successfully",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			a := &Application{
				Closers: &closer.Closers{},
				ctx:     context.Background(),
			}
			err := a.Run(context.Background())
			assert.Equalf(t, tt.wantErr, err != nil, "Run(ctx)")
		})
	}
}

func TestApplication_Stop(t *testing.T) {
	type fields struct {
		closers []closer.Closer
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "when a closer returns an error it returns an error",
			fields: fields{
				closers: []closer.Closer{
					func(context.Context) error { return errors.New("any-error") },
				},
			},
			wantErr: true,
		},
		{
			name: "when closers complete successfully it returns no error",
			fields: fields{
				closers: []closer.Closer{
					func(context.Context) error { return nil },
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			closers := &closer.Closers{}
			for _, cls := range tt.fields.closers {
				closers.Add(cls)
			}

			a := &Application{
				Closers: closers,
				ctx:     context.Background(),
			}
			err := a.Stop(context.Background())
			assert.Equalf(t, tt.wantErr, err != nil, "Stop(ctx)")
		})
	}
}
