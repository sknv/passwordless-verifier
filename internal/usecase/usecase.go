package usecase

import (
	"github.com/sknv/passwordless-verifier/internal/model"
)

const fieldParams = "params"

type Config struct {
	DeeplinkFormat string
}

type Usecase struct {
	Config Config
	DB     model.DB
}
