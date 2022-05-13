package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerificationMethod_Validate(t *testing.T) {
	tests := []struct {
		name    string
		m       VerificationMethod
		wantErr bool
	}{
		{
			name:    "when allowed method is provided it returns no error",
			m:       VerificationMethodTelegram,
			wantErr: false,
		},
		{
			name:    "when unknown method is provided it returns an error",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := tt.m.Validate()
			assert.Equalf(t, tt.wantErr, err != nil, "Validate()")
		})
	}
}
