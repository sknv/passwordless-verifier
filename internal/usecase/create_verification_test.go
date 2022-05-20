package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/sknv/passwordless-verifier/internal/model"
)

func TestNewVerification_Validate(t *testing.T) {
	type fields struct {
		method model.VerificationMethod
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
				method: model.VerificationMethodTelegram,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			v := NewVerification{Method: tt.fields.method}
			err := v.Validate()
			assert.Equalf(t, tt.wantErr, err != nil, "Validate()")
		})
	}
}

func TestUsecase_CreateVerification(t *testing.T) {
	type fields struct {
		store Store
	}
	type args struct {
		newVerification *NewVerification
	}

	tests := []struct {
		name          string
		prepareFields func() *fields
		args          args
		want          assert.ValueAssertionFunc
		wantErr       bool
	}{
		{
			name:          "when args are not valid it returns an error",
			prepareFields: func() *fields { return &fields{} },
			args: args{
				newVerification: &NewVerification{},
			},
			wantErr: true,
		},
		{
			name: "when args are valid it creates and returns a verification",
			prepareFields: func() *fields {
				return &fields{
					store: &StoreMock{
						CreateVerificationFunc: func(context.Context, *model.Verification) error { return nil },
					},
				}
			},
			args: args{
				newVerification: &NewVerification{
					Method: model.VerificationMethodTelegram,
				},
			},
			want: func(t assert.TestingT, actual any, msgAndArgs ...any) bool {
				want := &model.Verification{
					Method: model.VerificationMethodTelegram,
					Status: model.VerificationStatusInProgress,
				}

				got, _ := actual.(*model.Verification)
				got.ID, got.Deeplink, got.CreatedAt = uuid.UUID{}, "", time.Time{} // ignore fields when compare

				return assert.Equal(t, want, got, msgAndArgs...)
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
			got, err := u.CreateVerification(context.Background(), tt.args.newVerification)
			assert.Equalf(t, tt.wantErr, err != nil, "CreateVerification(ctx, %v)", tt.args.newVerification)
			if tt.want != nil {
				tt.want(t, got, "CreateVerification(ctx, %v)", tt.args.newVerification)
			}
		})
	}
}
