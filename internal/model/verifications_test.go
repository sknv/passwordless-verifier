package model

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestVerifications_FindByID(t *testing.T) {
	type fields struct {
		db DB
	}
	type args struct {
		id uuid.UUID
	}

	id := uuid.New()

	tests := []struct {
		name          string
		prepareFields func() *fields
		args          args
		want          func(*fields) *Verification
		wantErr       bool
	}{
		{
			name: "when db returns an error it returns an error",
			prepareFields: func() *fields {
				return &fields{
					db: &DBMock{
						FindFunc: func(context.Context, any, string, ...any) error { return errors.New("any-error") },
					},
				}
			},
			args: args{
				id: id,
			},
			wantErr: true,
		},
		{
			name: "when db call is ok it returns a result with db field has been set",
			prepareFields: func() *fields {
				return &fields{
					db: &DBMock{
						FindFunc: func(_ context.Context, dest any, _ string, _ ...any) error {
							verification, _ := dest.(*Verification)
							verification.ID = id

							return nil
						},
					},
				}
			},
			args: args{
				id: id,
			},
			want: func(f *fields) *Verification {
				return &Verification{
					DB: f.db,

					ID: id,
				}
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Prepare fields
			fields := tt.prepareFields()

			v := &Verifications{
				DB: fields.db,
			}
			got, err := v.FindByID(context.Background(), tt.args.id)
			assert.Equalf(t, tt.wantErr, err != nil, "FindByID(ctx, %v)", tt.args.id)
			if tt.want != nil {
				assert.Equalf(t, tt.want(fields), got, "FindByID(ctx, %v)", tt.args.id)
			}
		})
	}
}
