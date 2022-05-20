package converter

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sknv/passwordless-verifier/api/openapi"
	"github.com/sknv/passwordless-verifier/internal/model"
)

func TestToSession(t *testing.T) {
	type args struct {
		session *model.Session
	}

	now, phone := time.Now(), "+79001002030"

	tests := []struct {
		name string
		args args
		want *openapi.Session
	}{
		{
			name: "when nil session provided it returns nil",
		},
		{
			name: "when valid session provided it converts a model to an http response",
			args: args{
				session: &model.Session{
					PhoneNumber: phone,
					CreatedAt:   now,
				},
			},
			want: &openapi.Session{
				PhoneNumber: phone,
				CreatedAt:   now,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equalf(t, tt.want, ToSession(tt.args.session), "ToSession(%v)", tt.args.session)
		})
	}
}
