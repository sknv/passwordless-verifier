package usecase

import (
	"github.com/uptrace/bun"
)

const fieldParams = "params"

type Usecase struct {
	DB *bun.DB
}
