package usecase

import (
	"testing"

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
			wantErr: false,
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
