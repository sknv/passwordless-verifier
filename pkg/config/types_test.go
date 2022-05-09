package config

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDuration_UnmarshalText(t *testing.T) {
	type args struct {
		text string
	}

	tests := []struct {
		name    string
		args    args
		want    time.Duration
		wantErr bool
	}{
		{
			name: "when a valid duration string is provided it parses the string successfully",
			args: args{
				text: "1m",
			},
			want: time.Minute,
		},
		{
			name: "when a invalid duration string is provided it returns an error",
			args: args{
				text: "invalid-duration",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var dur Duration
			err := dur.UnmarshalText([]byte(tt.args.text))
			assert.Equalf(t, tt.wantErr, err != nil, "UnmarshalText(%v)", tt.args.text)
			assert.Equalf(t, tt.want, dur.Duration(), "UnmarshalText(%v)", tt.args.text)
		})
	}
}
