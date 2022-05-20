package converter

import (
	"github.com/sknv/passwordless-verifier/api/openapi"
	"github.com/sknv/passwordless-verifier/internal/model"
)

func ToSession(session *model.Session) *openapi.Session {
	if session == nil {
		return nil
	}

	return &openapi.Session{
		PhoneNumber: session.PhoneNumber,
		CreatedAt:   session.CreatedAt,
	}
}
