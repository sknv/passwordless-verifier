package usecase

import (
	"github.com/sknv/passwordless-verifier/internal/model"
)

const fieldParams = "params"

type Config struct {
	DeeplinkFormat string
}

type DB interface {
	model.DB
}

type Usecase struct {
	Config Config
	DB     DB
}
