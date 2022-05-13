package problem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem_Error(t *testing.T) {
	type fields struct {
		problemType string
		status      int
		title       string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "it returns an error with expected format",
			fields: fields{
				problemType: "bad-request",
				status:      400,
				title:       "any-title",
			},
			want: "type = bad-request, status = 400, title = any-title",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			p := &Problem{
				Type:   tt.fields.problemType,
				Title:  tt.fields.title,
				Status: tt.fields.status,
			}
			assert.Equalf(t, tt.want, p.Error(), "Error()")
		})
	}
}
