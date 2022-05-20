package usecase

import (
	"database/sql"
	"errors"

	"github.com/sknv/passwordless-verifier/pkg/http/problem"
)

// ErrWrongContact wrong contact provided
var ErrWrongContact = errors.New("wrong contact")

func ConvertStoreError(err error) error {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		businessProblem := problem.Business("not-found", "object not found")
		businessProblem.Err = err

		return businessProblem
	default:
		return err
	}
}
