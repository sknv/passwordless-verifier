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
		id string
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
				id: id.String(),
			},
			want: id,
		},
		{
			name: "when invalid id is provided it returns an empty uuid",
			fields: fields{
				id: "invalid-uuid",
			},
			want: uuid.UUID{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			p := GetVerificationParams{
				ID: tt.fields.id,
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
			name:    "when invalid id is provided it returns an error",
			wantErr: true,
		},
		{
			name: "when valid id is provided it returns no error",
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
		store Store
	}
	type args struct {
		params *GetVerificationParams
	}

	verificationID := uuid.New()

	tests := []struct {
		name          string
		prepareFields func() *fields
		args          args
		want          *model.Verification
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
					store: &StoreMock{
						FindVerificationByIDWithSessionFunc: func(_ context.Context, id uuid.UUID) (*model.Verification, error) {
							return &model.Verification{ID: id}, nil
						},
					},
				}
			},
			args: args{
				params: &GetVerificationParams{
					ID: verificationID.String(),
				},
			},
			want: &model.Verification{
				ID: verificationID,
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
				Store: fields.store,
			}
			got, err := u.GetVerification(context.Background(), tt.args.params)
			assert.Equalf(t, tt.wantErr, err != nil, "GetVerification(ctx, %v)", tt.args.params)
			assert.Equalf(t, tt.want, got, "GetVerification(ctx, %v)", tt.args.params)
		})
	}
}
