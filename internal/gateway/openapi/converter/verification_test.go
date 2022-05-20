package converter

import (
	"testing"
	"time"

	openapiTypes "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/sknv/passwordless-verifier/api/openapi"
	"github.com/sknv/passwordless-verifier/internal/model"
	"github.com/sknv/passwordless-verifier/internal/usecase"
)

func TestFromNewVerification(t *testing.T) {
	type args struct {
		newVerification *openapi.NewVerification
	}

	tests := []struct {
		name string
		args args
		want *usecase.NewVerification
	}{
		{
			name: "it converts an http request to a usecase request",
			args: args{
				newVerification: &openapi.NewVerification{
					Method: openapi.VerificationMethodTelegram,
				},
			},
			want: &usecase.NewVerification{
				Method: model.VerificationMethodTelegram,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := FromNewVerification(tt.args.newVerification)
			assert.Equalf(t, tt.want, got, "FromNewVerification(%v)", tt.args.newVerification)
		})
	}
}

func TestToVerification(t *testing.T) {
	type args struct {
		verification *model.Verification
	}

	verificationID, now, phone := uuid.New(), time.Now(), "+79001002030"

	tests := []struct {
		name string
		args args
		want *openapi.Verification
	}{
		{
			name: "it converts a model to an http response",
			args: args{
				verification: &model.Verification{
					ID:        verificationID,
					Method:    model.VerificationMethodTelegram,
					Deeplink:  "https://t.me/example_bot?start=123",
					Status:    model.VerificationStatusInProgress,
					CreatedAt: now,
					Session: &model.Session{
						PhoneNumber: phone,
						CreatedAt:   now,
					},
				},
			},
			want: &openapi.Verification{
				Id:        openapiTypes.UUID(verificationID.String()),
				Method:    openapi.VerificationMethodTelegram,
				Deeplink:  "https://t.me/example_bot?start=123",
				Status:    openapi.VerificationStatusInProgress,
				CreatedAt: now,
				Session: &openapi.Session{
					PhoneNumber: phone,
					CreatedAt:   now,
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equalf(t, tt.want, ToVerification(tt.args.verification), "ToVerification(%v)", tt.args.verification)
		})
	}
}
