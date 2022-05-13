package model

import (
	"database/sql"
	"errors"

	"github.com/sknv/passwordless-verifier/pkg/http/problem"
)

func ConvertDBError(err error) error {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		businessProblem := problem.Business("not-found", "object not found")
		businessProblem.Err = err

		return businessProblem
	default:
		return err
	}
}
