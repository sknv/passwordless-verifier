package usecase

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bun"

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

func TestNewVerification_ToVerification(t *testing.T) {
	type fields struct {
		method model.VerificationMethod
	}
	type args struct {
		db *bun.DB
	}

	db := &bun.DB{}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   assert.ValueAssertionFunc
	}{
		{
			name: "it constructs a new verification",
			fields: fields{
				method: model.VerificationMethodTelegram,
			},
			args: args{
				db: db,
			},
			want: func(t assert.TestingT, actual any, msgAndArgs ...any) bool {
				want := &model.Verification{
					DB: db,

					Method: model.VerificationMethodTelegram,
					Status: model.VerificationStatusInProgress,
				}

				got, _ := actual.(*model.Verification)
				got.ID = uuid.UUID{} // ignore id field when compare

				return assert.Equal(t, want, got, msgAndArgs)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			v := NewVerification{
				Method: tt.fields.method,
			}
			tt.want(t, v.ToVerification(tt.args.db), "ToVerification(%v)", tt.args.db)
		})
	}
}
