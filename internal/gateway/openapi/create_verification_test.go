package openapi

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	openapiTypes "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/sknv/passwordless-verifier/api/openapi"
	"github.com/sknv/passwordless-verifier/internal/model"
	"github.com/sknv/passwordless-verifier/internal/usecase"
)

//go:generate moq -out mocks_test.go . Usecase

func TestServer_CreateVerification(t *testing.T) {
	type (
		createVerificationContract struct {
			in    *usecase.NewVerification
			out   *model.Verification
			err   error
			times int
		}

		usecaseContract struct {
			createVerification createVerificationContract
		}
	)
	type fields struct {
		usecase usecaseContract
	}
	type args struct {
		req string
	}

	verificationUUID := uuid.New()

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *openapi.Verification
		wantErr bool
	}{
		{
			name: "when usecase returns an error it renders an error response",
			fields: fields{
				usecase: usecaseContract{
					createVerification: createVerificationContract{
						in: &usecase.NewVerification{
							Method: model.VerificationMethodTelegram,
						},
						err:   errors.New("any-error"),
						times: 1,
					},
				},
			},
			args: args{
				req: `{"method": "telegram"}`,
			},
			wantErr: true,
		},
		{
			name: "when usecase returns an verification it renders a verification response",
			fields: fields{
				usecase: usecaseContract{
					createVerification: createVerificationContract{
						in: &usecase.NewVerification{
							Method: model.VerificationMethodTelegram,
						},
						out: &model.Verification{
							ID:     verificationUUID,
							Method: model.VerificationMethodTelegram,
							Status: model.VerificationStatusInProgress,
						},
						times: 1,
					},
				},
			},
			args: args{
				req: `{"method": "telegram"}`,
			},
			want: &openapi.Verification{
				Id:     openapiTypes.UUID(verificationUUID.String()),
				Method: openapi.VerificationMethodTelegram,
				Status: openapi.VerificationStatusInProgress,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Prepare request
			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(tt.args.req)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			c := echo.New().NewContext(req, rec)
			c.SetPath("/api/verifications")

			// Prepare deps
			usecaseMock := &UsecaseMock{
				CreateVerificationFunc: func(ctx context.Context, newVerification *usecase.NewVerification) (*model.Verification, error) {
					assert.Equalf(t, tt.fields.usecase.createVerification.in, newVerification,
						"usecase.CreateVerification(%v)", newVerification)
					return tt.fields.usecase.createVerification.out, tt.fields.usecase.createVerification.err
				},
			}

			// Construct test object
			srv := &Server{
				Usecase: usecaseMock,
			}
			s := openapi.ServerInterfaceWrapper{Handler: srv}

			err := s.CreateVerification(c)
			assert.Equalf(t, tt.fields.usecase.createVerification.times, len(usecaseMock.CreateVerificationCalls()),
				"usecase.CreateVerificationCalls")
			assert.Equalf(t, tt.wantErr, err != nil, "CreateVerification(%v)", tt.args.req)

			if tt.want == nil {
				return
			}

			resp := &openapi.Verification{}
			err = json.Unmarshal(rec.Body.Bytes(), resp)
			assert.NoError(t, err)
			assert.Equalf(t, tt.want, resp, "CreateVerification(%v)", tt.args.req)
		})
	}
}
