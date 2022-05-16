package usecase

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/sknv/passwordless-verifier/internal/model"
)

func TestGetVerificationParams_TypedID(t *testing.T) {
	type fields struct {
		ID string
	}

	id := uuid.New()

	tests := []struct {
		name   string
		fields fields
		want   uuid.UUID
	}{
		{
			name: "when valid id is provided it returns a filled uuid",
			fields: fields{
				ID: id.String(),
			},
			want: id,
		},
		{
			name: "when invalid id is provided it returns an empty uuid",
			fields: fields{
				ID: "invalid-uuid",
			},
			want: uuid.UUID{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			p := GetVerificationParams{
				ID: tt.fields.ID,
			}
			assert.Equalf(t, tt.want, p.TypedID(), "TypedID()")
		})
	}
}

func TestGetVerificationParams_Validate(t *testing.T) {
	type fields struct {
		id string
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "when invalid method is provided it returns an error",
			wantErr: true,
		},
		{
			name: "when valid method is provided it returns no error",
			fields: fields{
				id: uuid.New().String(),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			p := GetVerificationParams{
				ID: tt.fields.id,
			}
			err := p.Validate()
			assert.Equalf(t, tt.wantErr, err != nil, "Validate()")
		})
	}
}

func TestUsecase_GetVerification(t *testing.T) {
	type fields struct {
		db model.DB
	}
	type args struct {
		params *GetVerificationParams
	}

	id := uuid.New()

	tests := []struct {
		name          string
		prepareFields func() *fields
		args          args
		want          func(*fields) assert.ValueAssertionFunc
		wantErr       bool
	}{
		{
			name:          "when args are not valid it returns an error",
			prepareFields: func() *fields { return &fields{} },
			args: args{
				params: &GetVerificationParams{},
			},
			wantErr: true,
		},
		{
			name: "when args are valid it finds and returns a verification",
			prepareFields: func() *fields {
				return &fields{
					db: &DBMock{
						FindFunc: func(_ context.Context, dest any, _ string, _ ...any) error {
							verification, _ := dest.(*model.Verification)
							verification.ID = id

							return nil
						},
					},
				}
			},
			args: args{
				params: &GetVerificationParams{
					ID: id.String(),
				},
			},
			want: func(f *fields) assert.ValueAssertionFunc {
				return func(t assert.TestingT, actual any, msgAndArgs ...any) bool {
					want := &model.Verification{
						DB: f.db,

						ID: id,
					}

					return assert.Equal(t, want, actual, msgAndArgs)
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

			u := &Usecase{
				DB: fields.db,
			}
			got, err := u.GetVerification(context.Background(), tt.args.params)
			assert.Equalf(t, tt.wantErr, err != nil, "GetVerification(ctx, %v)", tt.args.params)
			if tt.want != nil {
				tt.want(fields)(t, got, "GetVerification(ctx, %v)", tt.args.params)
			}
		})
	}
}
