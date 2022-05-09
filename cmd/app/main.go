package main

import (
	"context"
	"flag"
	stdlog "log"
	"time"

	"github.com/sknv/passwordless-verifier/pkg/application"
	"github.com/sknv/passwordless-verifier/pkg/log"
	"github.com/sknv/passwordless-verifier/pkg/os"
)

const (
	applicationStartTimeout = time.Second * 30
	applicationStopTimeout  = time.Second * 30
)

func main() {
	configPath := ConfigFilePathFlag()
	flag.Parse()

	cfg, err := ParseConfig(*configPath)
	if err != nil {
		stdlog.Fatalf("parse config: %s", err)
	}

	// Create an application core
	app := application.NewApplication(context.Background())

	// Prepare dependencies
	makeLogger(app, cfg)
	logger := log.Extract(app.Context())

	// Run the app
	if err = runApp(app, applicationStartTimeout); err != nil {
		logger.WithError(err).Fatal("start application")
	}

	<-os.NotifyAboutExit() // wait for the program exit

	// Close the app applying deferred closers
	if err = stopApp(app, applicationStopTimeout); err != nil {
		logger.WithError(err).Error("stop application")
	}
}

func makeLogger(app *application.Application, config *Config) {
	app.RegisterLogger(log.Config{Level: config.LogConfig.Level})
}

func runApp(app *application.Application, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(app.Context(), timeout)
	defer cancel()

	return app.Run(ctx)
}

func stopApp(app *application.Application, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(app.Context(), timeout)
	defer cancel()

	return app.Stop(ctx)
}
