package usecase

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sknv/passwordless-verifier/pkg/http/problem"
)

func TestConvertStoreError(t *testing.T) {
	type args struct {
		err error
	}

	anyErr := errors.New("any-error")

	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "when error is sql.ErrNoRows it returns a problem.Business",
			args: args{
				err: sql.ErrNoRows,
			},
			wantErr: func(t assert.TestingT, err error, msgAndArgs ...any) bool {
				var prb *problem.Problem
				return errors.As(err, &prb) &&
					assert.Equal(t, http.StatusUnprocessableEntity, prb.Status, msgAndArgs...)
			},
		},
		{
			name: "for a random error it proxies the one",
			args: args{
				err: anyErr,
			},
			wantErr: func(t assert.TestingT, err error, msgAndArgs ...any) bool {
				return errors.Is(err, anyErr)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.wantErr(t, ConvertStoreError(tt.args.err), fmt.Sprintf("ConvertStoreError(%v)", tt.args.err))
		})
	}
}
