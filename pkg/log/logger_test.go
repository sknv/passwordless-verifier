package log

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestBuild(t *testing.T) {
	type args struct {
		level   logrus.Level
		options []Option
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "it builds global logger successfully",
			args: args{
				level: logrus.DebugLevel,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			Build(tt.args.level, tt.args.options...)
			assert.Equalf(t, tt.args.level, L().Level, "Build(%v, %v)", tt.args.level, tt.args.options)
		})
	}
}

func TestParseLevel(t *testing.T) {
	type args struct {
		level string
	}

	tests := []struct {
		name string
		args args
		want logrus.Level
	}{
		{
			name: "when known level is provided it parses it successfully",
			args: args{
				level: "debug",
			},
			want: logrus.DebugLevel,
		},
		{
			name: "when unknown level is provided it uses the default one",
			args: args{
				level: "unknown",
			},
			want: DefaultLevel,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equalf(t, tt.want, ParseLevel(tt.args.level), "ParseLevel(%v)", tt.args.level)
		})
	}
}
