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
		level     logrus.Level
		formatter logrus.Formatter
		options   []Option
	}
	type want struct {
		level     logrus.Level
		formatter logrus.Formatter
		out       io.Writer
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "it builds global logger and apply options successfully",
			args: args{
				level:     logrus.DebugLevel,
				formatter: &logrus.JSONFormatter{},
				options: []Option{
					func(l *logrus.Logger) {
						l.SetOutput(os.Stderr)
					},
				},
			},
			want: want{
				level:     logrus.DebugLevel,
				formatter: &logrus.JSONFormatter{},
				out:       os.Stderr,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			Build(tt.args.level, tt.args.formatter, tt.args.options...)
			assert.Equalf(t, tt.want, want{
				level:     L().Level,
				formatter: L().Formatter,
				out:       L().Out,
			}, "Build(%v, %v, %v)", tt.args.level, tt.args.formatter, tt.args.options)
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
			want: logrus.InfoLevel,
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

func TestGetFormatter(t *testing.T) {
	type args struct {
		formatter Formatter
	}

	tests := []struct {
		name string
		args args
		want assert.ValueAssertionFunc
	}{
		{
			name: "when text formatter provided it returns a logrus.TextFormatter",
			args: args{
				formatter: FormatterText,
			},
			want: func(t assert.TestingT, actual any, msgAndArgs ...any) bool {
				_, ok := actual.(*logrus.TextFormatter)
				return assert.True(t, ok, msgAndArgs)
			},
		},
		{
			name: "when any other formatter provided it returns a default formatter",
			args: args{},
			want: func(t assert.TestingT, actual any, msgAndArgs ...any) bool {
				_, ok := actual.(*logrus.JSONFormatter)
				return assert.True(t, ok, msgAndArgs)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.want(t, GetFormatter(tt.args.formatter), "GetFormatter(%v)", tt.args.formatter)
		})
	}
}
