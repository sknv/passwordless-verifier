package openapi

import (
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

//go:generate moq -out mocks_test.go -fmt goimports . Usecase

func TestServer_Route(t *testing.T) {
	type args struct {
		e *echo.Echo
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "it registers some handlers",
			args: args{
				e: echo.New(),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := &Server{}
			s.Route(tt.args.e)
			assert.Positivef(t, len(tt.args.e.Routes()), "Route(e)")
		})
	}
}
