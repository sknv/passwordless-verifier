package problem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem_Error(t *testing.T) {
	type fields struct {
		status int
		title  string
		detail string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "it returns an error with expected format",
			fields: fields{
				status: 400,
				title:  "any-title",
				detail: "any-detail",
			},
			want: "status = 400, title = any-title, detail = any-detail",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			p := &Problem{
				Title:  tt.fields.title,
				Status: tt.fields.status,
				Detail: tt.fields.detail,
			}
			assert.Equalf(t, tt.want, p.Error(), "Error()")
		})
	}
}
