package log

import (
	"io"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestBuild(t *testing.T) {
	type args struct {
		level   logrus.Level
		options []Option
	}
	type want struct {
		level logrus.Level
		out   io.Writer
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "it builds global logger and apply options successfully",
			args: args{
				level: logrus.DebugLevel,
				options: []Option{
					func(l *logrus.Logger) {
						l.SetOutput(os.Stderr)
					},
				},
			},
			want: want{
				level: logrus.DebugLevel,
				out:   os.Stderr,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			Build(tt.args.level, tt.args.options...)
			assert.Equalf(t, tt.want, want{
				level: L().Level,
				out:   L().Out,
			}, "Build(%v, %v)", tt.args.level, tt.args.options)
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
