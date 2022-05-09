package log

import (
	"context"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestExtract1(t *testing.T) {
	type args struct {
		logger logrus.FieldLogger
		fields logrus.Fields
	}

	tests := []struct {
		name string
		args args
		want assert.ValueAssertionFunc
	}{
		{
			name: "when there is no logger it returns an empty one",
			want: func(t assert.TestingT, actual any, msgAndArgs ...any) bool {
				return assert.Equal(t, _nullLogger, actual, msgAndArgs...)
			},
		},
		{
			name: "when there are no fields it returns a provided logger",
			args: args{
				logger: L(),
			},
			want: func(t assert.TestingT, actual any, msgAndArgs ...any) bool {
				return assert.Equal(t, L(), actual, msgAndArgs...)
			},
		},
		{
			name: "when a logger with fields is provided it returns the one",
			args: args{
				logger: _nullLogger,
				fields: logrus.Fields{
					"key": "value",
				},
			},
			want: func(t assert.TestingT, actual any, msgAndArgs ...any) bool {
				return assert.NotEqual(t, _nullLogger, actual, msgAndArgs...)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := ToContext(context.Background(), tt.args.logger)
			AddFields(ctx, tt.args.fields)
			resultCtx := Extract(ctx)
			tt.want(t, resultCtx, "Extract(%v, %v)", tt.args.logger, tt.args.fields)
		})
	}
}
