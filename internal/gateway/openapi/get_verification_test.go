package openapi

import (
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

func TestServer_GetVerification(t *testing.T) {
	type fields struct {
		usecase Usecase
	}
	type args struct {
		id openapiTypes.UUID
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
				in := &usecase.GetVerificationParams{
					ID: verificationUUID.String(),
				}
				out := &model.Verification{
					ID:     verificationUUID,
					Method: model.VerificationMethodTelegram,
					Status: model.VerificationStatusInProgress,
				}

				return &fields{
					usecase: &UsecaseMock{
						GetVerificationFunc: func(_ context.Context, params *usecase.GetVerificationParams) (*model.Verification, error) {
							assert.Equalf(t, in, params, "usecase.GetVerification(%v)", params)
							return out, nil
						},
					},
				}
			},
			args: args{
				id: openapiTypes.UUID(verificationUUID.String()),
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
			req := httptest.NewRequest(http.MethodPost, "/", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			c := echo.New().NewContext(req, rec)
			c.SetPath("/api/verifications/:id")
			c.SetParamNames("id")
			c.SetParamValues(string(tt.args.id))

			// Prepare fields
			fields := tt.prepareFields()

			// Construct test object
			srv := &Server{
				Usecase: fields.usecase,
			}
			s := openapi.ServerInterfaceWrapper{Handler: srv}

			err := s.GetVerification(c)
			assert.NoErrorf(t, err, "GetVerification(%v)", tt.args.id)

			resp := &openapi.Verification{}
			err = json.Unmarshal(rec.Body.Bytes(), resp)
			assert.NoError(t, err)
			assert.Equalf(t, tt.want, resp, "GetVerification(%v)", tt.args.id)
		})
	}
}
