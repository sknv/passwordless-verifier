package openapi

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/sknv/passwordless-verifier/api/openapi"
)

func TestServer_CreateVerification(t *testing.T) {
	type args struct {
		req *openapi.NewVerification
	}

	tests := []struct {
		name    string
		args    args
		want    *openapi.Verification
		wantErr bool
	}{
		{
			name: "it returns no error",
			args: args{
				req: &openapi.NewVerification{
					Method: openapi.VerificationMethodTelegram,
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			request, err := json.Marshal(tt.args.req)
			assert.NoError(t, err)

			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(request))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			c := echo.New().NewContext(req, rec)
			c.SetPath("/api/verifications")

			srv := &Server{}
			s := openapi.ServerInterfaceWrapper{Handler: srv}

			err = s.CreateVerification(c)
			assert.Equalf(t, tt.wantErr, err != nil, "CreateVerification(%v)", tt.args.req)

			if tt.want == nil {
				return
			}

			resp := &openapi.Verification{}
			err = json.Unmarshal(rec.Body.Bytes(), resp)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, resp)
		})
	}
}
