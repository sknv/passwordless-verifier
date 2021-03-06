package problem

import (
	"net/http"
)

const (
	_badRequestType          = "bad-request"
	_unauthorizedType        = "unauthorized"
	_notFoundType            = "not-found"
	_internalServerErrorType = "internal-server-error"
)

type invalidParams struct {
	InvalidParams []InvalidParam `json:"invalidParams,omitempty"`
}

func BadRequest(params ...InvalidParam) *Problem {
	problem := New(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	problem.Type = _badRequestType
	if len(params) > 0 {
		problem.Data = invalidParams{
			InvalidParams: params,
		}
	}

	return problem
}

func Unauthorized() *Problem {
	problem := New(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
	problem.Type = _unauthorizedType

	return problem
}

func NotFound() *Problem {
	problem := New(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	problem.Type = _notFoundType

	return problem
}

func Business(problemType, title string) *Problem {
	problem := New(http.StatusUnprocessableEntity, title)
	problem.Type = problemType

	return problem
}

func InternalServerError() *Problem {
	problem := New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	problem.Type = _internalServerErrorType

	return problem
}
