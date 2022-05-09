package closer

import (
	"context"
	"errors"
	"testing"

	"github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/assert"
)

func TestClosers_Close(t *testing.T) {
	type fields struct {
		closers []Closer
	}

	err1, err2 := errors.New("1"), errors.New("2")

	tests := []struct {
		name    string
		fields  fields
		wantErr error
	}{
		{
			name: "it closes successfully",
			fields: fields{
				closers: []Closer{
					func(context.Context) error {
						return nil
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "it closes with a single error",
			fields: fields{
				closers: []Closer{
					func(context.Context) error {
						return nil
					},
					func(context.Context) error {
						return err1
					},
				},
			},
			wantErr: multierror.Append(err1, nil),
		},
		{
			name: "when multiple errors are raised it returns them in correct order",
			fields: fields{
				closers: []Closer{
					func(context.Context) error {
						return err1
					},
					func(context.Context) error {
						return err2
					},
				},
			},
			wantErr: multierror.Append(err2, err1),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c := &Closers{}
			for _, cls := range tt.fields.closers {
				c.Add(cls)
			}

			assert.Equalf(t, tt.wantErr, c.Close(context.Background()), "Close(%v)", tt.fields.closers)
		})
	}
}
