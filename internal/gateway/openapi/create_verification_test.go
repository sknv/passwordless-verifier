package openapi

import (
	"bytes"
	"context"
	"encoding/json"
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

func TestServer_CreateVerification(t *testing.T) {
	type fields struct {
		usecase Usecase
	}
	type args struct {
		req string
	}

	verificationUUID := uuid.New()

	tests := []struct {
		name          string
		prepareFields func() *fields
		args          args
		want          *openapi.Verification
	}{
		{
			name: "when usecase returns a verification it renders a verification response",
			prepareFields: func() *fields {
				in := &usecase.NewVerification{
					Method: model.VerificationMethodTelegram,
				}
				out := &model.Verification{
					ID:     verificationUUID,
					Method: model.VerificationMethodTelegram,
					Status: model.VerificationStatusInProgress,
				}

				return &fields{
					usecase: &UsecaseMock{
						CreateVerificationFunc: func(_ context.Context, newVerification *usecase.NewVerification) (*model.Verification, error) {
							assert.Equalf(t, in, newVerification, "usecase.CreateVerification(%v)", newVerification)
							return out, nil
						},
					},
				}
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

			// Prepare fields
			fields := tt.prepareFields()

			// Construct test object
			srv := &Server{
				Usecase: fields.usecase,
			}
			s := openapi.ServerInterfaceWrapper{Handler: srv}

			err := s.CreateVerification(c)
			assert.NoErrorf(t, err, "CreateVerification(%v)", tt.args.req)

			resp := &openapi.Verification{}
			err = json.Unmarshal(rec.Body.Bytes(), resp)
			assert.NoError(t, err)
			assert.Equalf(t, tt.want, resp, "CreateVerification(%v)", tt.args.req)
		})
	}
}
