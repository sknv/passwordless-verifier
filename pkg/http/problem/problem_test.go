package problem

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem_Error(t *testing.T) {
	type fields struct {
		err         error
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
			name: "when only required fields are provided it returns the expected message",
			fields: fields{
				status: 400,
				title:  "any-title",
			},
			want: "status = 400, title = any-title",
		},
		{
			name: "when type is also provided it returns the expected message",
			fields: fields{
				problemType: "bad-request",
				status:      400,
				title:       "any-title",
			},
			want: "status = 400, title = any-title, type = bad-request",
		},
		{
			name: "when a wrapped error presents it returns the expected message",
			fields: fields{
				err:         errors.New("any-error"),
				problemType: "bad-request",
				status:      400,
				title:       "any-title",
			},
			want: "status = 400, title = any-title, type = bad-request, err = any-error",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			p := &Problem{
				Err:    tt.fields.err,
				Type:   tt.fields.problemType,
				Title:  tt.fields.title,
				Status: tt.fields.status,
			}
			assert.Equalf(t, tt.want, p.Error(), "Error()")
		})
	}
}
