package model

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestVerificationMethod_Validate(t *testing.T) {
	tests := []struct {
		name    string
		m       VerificationMethod
		wantErr bool
	}{
		{
			name: "when allowed method is provided it returns no error",
			m:    VerificationMethodTelegram,
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

func TestFormatDeeplink(t *testing.T) {
	type args struct {
		deeplink       string
		verificationID uuid.UUID
	}

	deeplink, verificationID := "https://t.me/example_bot", uuid.New()

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "it returns an expected format",
			args: args{
				deeplink:       deeplink,
				verificationID: verificationID,
			},
			want: fmt.Sprintf("%s?start=%s", deeplink, verificationID),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := FormatDeeplink(tt.args.deeplink, tt.args.verificationID)
			assert.Equalf(t, tt.want, got, "FormatDeeplink(%v, %v)", tt.args.deeplink, tt.args.verificationID)
		})
	}
}

func TestVerification_LogIn(t *testing.T) {
	type fields struct {
		id uuid.UUID
	}
	type args struct {
		phoneNumber string
	}

	verificationID, phone := uuid.New(), "+79001002030"

	tests := []struct {
		name   string
		fields fields
		args   args
		want   assert.ValueAssertionFunc
	}{
		{
			name: "it sets all the fields successfully",
			fields: fields{
				id: verificationID,
			},
			args: args{
				phoneNumber: phone,
			},
			want: func(t assert.TestingT, actual interface{}, msgAndArgs ...interface{}) bool {
				want := &Verification{
					ID:     verificationID,
					Status: VerificationStatusCompleted,
					Session: &Session{
						VerificationID: verificationID,
						PhoneNumber:    phone,
					},
				}

				got, _ := actual.(*Verification)
				got.Session.ID = uuid.UUID{} // ignore fields when compare

				return assert.Equal(t, want, got, msgAndArgs...)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			v := &Verification{
				ID: tt.fields.id,
			}
			v.LogIn(tt.args.phoneNumber)
			tt.want(t, v, "LogIn(%v)", tt.args.phoneNumber)
		})
	}
}
