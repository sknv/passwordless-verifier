package application

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/sknv/passwordless-verifier/pkg/http/server"
	"github.com/sknv/passwordless-verifier/pkg/log"
)

type HTTPServerConfig struct {
	server.Config

	Address string
}

func (a *Application) RegisterHTTPServer(config HTTPServerConfig, opts ...server.Option) *echo.Echo {
	srv := server.New(config.Config, opts...)
	a.httpServer = &preparedHTTPServer{
		address: config.Address,
		server:  srv,
	}

	return srv
}

type preparedHTTPServer struct {
	address string
	server  *echo.Echo
}

func (a *Application) runHTTPServer(ctx context.Context) {
	if a.httpServer == nil {
		return // no HTTP server registered
	}

	logger := log.Extract(ctx).WithField("address", a.httpServer.address)
	logger.Info("starting http server...")
	defer logger.Info("http server started")

	go func() {
		//nolint:errorlint // expect exactly the specified error
		if err := a.httpServer.server.Start(a.httpServer.address); err != nil && err != http.ErrServerClosed {
			logger.WithError(err).Fatal("start http server")
		}
	}()

	// Remember to stop the server
	a.closers.Add(func(closeCtx context.Context) error {
		logger.Info("stopping http server...")
		defer logger.Info("http server stopped")

		if err := a.httpServer.server.Shutdown(closeCtx); err != nil {
			return fmt.Errorf("shutdown http server: %w", err)
		}

		return nil
	})
}
