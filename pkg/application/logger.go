package application

import (
	"github.com/sknv/passwordless-verifier/pkg/log"
)

func (a *Application) RegisterLogger(config log.Config) {
	log.Init(config)

	// Update the application context with the logger
	a.ctx = log.ToContext(a.ctx, log.L())
}
