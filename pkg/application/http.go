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

func (a *Application) runHTTPServer(ctx context.Context, srv *preparedHTTPServer) {
	if srv == nil {
		return // no HTTP server registered
	}

	logger := log.Extract(ctx).WithField("address", srv.address)
	logger.Info("starting http server...")
	defer logger.Info("http server started")

	go func() {
		//nolint:errorlint // expect exactly the specified error
		if err := srv.server.Start(srv.address); err != nil && err != http.ErrServerClosed {
			logger.WithError(err).Fatal("start http server")
		}
	}()

	// Remember to stop the server
	a.Closers.Add(func(ctx context.Context) error {
		logger.Info("stopping http server...")
		defer logger.Info("http server stopped")

		if err := srv.server.Shutdown(ctx); err != nil {
			return fmt.Errorf("shutdown http server: %w", err)
		}

		return nil
	})
}
