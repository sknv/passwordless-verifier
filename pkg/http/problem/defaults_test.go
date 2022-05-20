package problem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBadRequest(t *testing.T) {
	type args struct {
		params []InvalidParam
	}

	invalidArgs := []InvalidParam{
		{},
	}

	tests := []struct {
		name string
		args args
		want assert.ValueAssertionFunc
	}{
		{
			name: "when params are provided it sets data field",
			args: args{
				params: invalidArgs,
			},
			want: func(t assert.TestingT, actual any, msgAndArgs ...any) bool {
				expected := BadRequest()
				expected.Data = invalidParams{
					InvalidParams: invalidArgs,
				}
				return assert.Equal(t, expected, actual, msgAndArgs...)
			},
		},
		{
			name: "when params are not provided it skips data field",
			want: func(t assert.TestingT, actual any, msgAndArgs ...any) bool {
				expected := BadRequest()
				return assert.Equal(t, expected, actual, msgAndArgs...)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.want(t, BadRequest(tt.args.params...), "BadRequest(%v)", tt.args.params)
		})
	}
}
