package main

import (
	"context"
	"flag"
	stdlog "log"
	"time"

	"github.com/uptrace/bun"

	"github.com/sknv/passwordless-verifier/internal/consumer/telegram"
	"github.com/sknv/passwordless-verifier/internal/gateway/openapi"
	"github.com/sknv/passwordless-verifier/internal/store"
	"github.com/sknv/passwordless-verifier/internal/usecase"
	"github.com/sknv/passwordless-verifier/pkg/application"
	"github.com/sknv/passwordless-verifier/pkg/http/server"
	"github.com/sknv/passwordless-verifier/pkg/log"
	"github.com/sknv/passwordless-verifier/pkg/os"
	"github.com/sknv/passwordless-verifier/pkg/postgres"
	"github.com/sknv/passwordless-verifier/pkg/tracing"
)

const _applicationStopTimeout = time.Second * 30

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

	if err = makeTracing(app, cfg); err != nil {
		logger.WithError(err).Fatal("register tracing")
	}

	db, err := makeDB(app, cfg)
	if err != nil {
		logger.WithError(err).Fatal("register postgres")
	}

	// Register a server
	svc := makeUsecase(cfg, db)
	makeHTTPServer(app, cfg, svc)

	// Register a telegram bot
	if err = makeTelegramBot(app, cfg, svc); err != nil {
		logger.WithError(err).Fatal("register telegram bot")
	}

	// Run the app
	if err = app.Run(); err != nil {
		logger.WithError(err).Fatal("start application")
	}

	<-os.NotifyAboutExit() // wait for the program exit

	// Close the app applying deferred closers
	if err = stopApp(app, _applicationStopTimeout); err != nil {
		logger.WithError(err).Error("stop application")
	}
}

func makeLogger(app *application.Application, config *Config) {
	app.RegisterLogger(log.Config{
		Level:     config.LogConfig.Level,
		Formatter: log.Formatter(config.LogConfig.Formatter),
	})
}

func makeTracing(app *application.Application, config *Config) error {
	return app.RegisterTracing(tracing.Config{
		Host:        config.Jaeger.Host,
		Port:        config.Jaeger.Port,
		ServiceName: config.App.Name,
		Ratio:       config.Jaeger.Ratio,
	})
}

func makeDB(app *application.Application, config *Config) (*bun.DB, error) {
	return app.RegisterPostgres(app.Context(), postgres.Config{
		URL:             config.Postgres.URL,
		MaxOpenConn:     config.Postgres.MaxOpenConn,
		MaxConnLifetime: config.Postgres.MaxConnLifetime.Duration(),
	})
}

func makeUsecase(config *Config, db *bun.DB) *usecase.Usecase {
	return &usecase.Usecase{
		Config: usecase.Config{Deeplink: config.Telegram.Deeplink},
		Store:  &store.DB{DB: db},
	}
}

func makeHTTPServer(app *application.Application, config *Config, usecase *usecase.Usecase) {
	// Create an HTTP server
	e := app.RegisterHTTPServer(application.HTTPServerConfig{
		Address: config.HTTP.Address,
		Config: server.Config{
			Metric: server.MetricConfig{Namespace: config.App.Name},
		},
	})

	srv := &openapi.Server{
		Usecase: usecase,
	}
	srv.Route(e)
}

func makeTelegramBot(app *application.Application, config *Config, usecase *usecase.Usecase) error {
	bot, err := telegram.NewBot(telegram.BotConfig{
		APIToken:          config.Telegram.APIToken,
		PollingTimeout:    config.Telegram.PollingTimeout.Duration(),
		MaxUpdatesAllowed: config.Telegram.MaxUpdatesAllowed,
		CallbackURL:       config.Telegram.CallbackURL,
		Debug:             config.Telegram.Debug,
	}, usecase)
	if err != nil {
		return err
	}

	app.RegisterConsumer(bot)
	return nil
}

func stopApp(app *application.Application, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(app.Context(), timeout)
	defer cancel()

	return app.Stop(ctx)
}
