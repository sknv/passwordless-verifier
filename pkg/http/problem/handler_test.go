package problem

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHTTPErrorHandler(t *testing.T) {
	type (
		requestContract struct {
			method string
		}

		responseContract struct {
			err         error
			isCommitted bool
		}

		want struct {
			code    int
			hasBody bool
		}
	)
	type args struct {
		req  requestContract
		resp responseContract
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "when a response is already committed it renders no body",
			args: args{
				req: requestContract{method: http.MethodGet},
				resp: responseContract{
					err:         BadRequest(),
					isCommitted: true,
				},
			},
			want: want{code: http.StatusBadRequest},
		},
		{
			name: "for a HEAD request it renders no body",
			args: args{
				req: requestContract{method: http.MethodHead},
				resp: responseContract{
					err: Business("problem-type", "problem-title"),
				},
			},
			want: want{code: http.StatusUnprocessableEntity},
		},
		{
			name: "when a problem is provided it renders the one",
			args: args{
				req: requestContract{method: http.MethodGet},
				resp: responseContract{
					err: Unauthorized(),
				},
			},
			want: want{
				code:    http.StatusUnauthorized,
				hasBody: true,
			},
		},
		{
			name: "when an echo.HTTPError is provided it maps one to a problem",
			args: args{
				req: requestContract{method: http.MethodGet},
				resp: responseContract{
					err: echo.NewHTTPError(http.StatusUnprocessableEntity).SetInternal(errors.New("any-error")),
				},
			},
			want: want{
				code:    http.StatusUnprocessableEntity,
				hasBody: true,
			},
		},
		{
			name: "when a method is not allowed it maps one to a NotFound problem",
			args: args{
				req: requestContract{method: http.MethodGet},
				resp: responseContract{
					err: echo.NewHTTPError(http.StatusMethodNotAllowed),
				},
			},
			want: want{
				code:    http.StatusNotFound,
				hasBody: true,
			},
		},
		{
			name: "when a random error is provided it renders a default problem",
			args: args{
				req: requestContract{method: http.MethodGet},
				resp: responseContract{
					err: errors.New("any-error"),
				},
			},
			want: want{
				code:    http.StatusInternalServerError,
				hasBody: true,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			req := httptest.NewRequest(tt.args.req.method, "/", nil)
			resp := httptest.NewRecorder()

			e := echo.New()
			c := e.NewContext(req, resp)
			if tt.args.resp.isCommitted {
				_ = c.NoContent(http.StatusBadRequest)
			}

			HTTPErrorHandler(tt.args.resp.err, c)
			assert.Equalf(t, tt.want.code, resp.Code, "HTTPErrorHandler(%v, %v)", tt.args.resp.err, c)
			assert.Equalf(t, tt.want.hasBody, resp.Body.Len() > 0, "HTTPErrorHandler(%v, %v)", tt.args.resp.err, c)
		})
	}
}
