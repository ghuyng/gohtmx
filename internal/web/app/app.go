package app

import (
	"context"
	"net/http"
	"time"

	"github.com/ghuyng/gohtmx/internal/config"
	"github.com/ghuyng/gohtmx/internal/web/handlers"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type App struct {
	cfg         *config.Config
	logger      *zap.SugaredLogger
	router      *echo.Echo
	server      *http.Server
	userHandler *handlers.UserHandler
}

func NewApp(cfg *config.Config, logger *zap.SugaredLogger) *App {
	router := echo.New()
	app := &App{
		cfg:         cfg,
		logger:      logger,
		router:      router,
		userHandler: handlers.NewUserHandler(logger),
	}

	app.setupRoutes()
	return app
}

func (a *App) Start() error {
	a.server = &http.Server{
		Addr:              a.cfg.Server.Addr,
		Handler:           a.router,
		ReadTimeout:       time.Duration(a.cfg.Server.ReadTimeoutInSeconds) * time.Second,
		ReadHeaderTimeout: time.Duration(a.cfg.Server.ReadHeaderTimeoutInSeconds) * time.Second,
		WriteTimeout:      time.Duration(a.cfg.Server.WriteTimeoutInSeconds) * time.Second,
		IdleTimeout:       time.Duration(a.cfg.Server.IdleTimeoutInSeconds) * time.Second,
	}
	a.logger.Info("starting server on ", a.cfg.Server.Addr)
	return a.server.ListenAndServe()
}

func (a *App) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10) // 10 seconds to shutdown
	defer cancel()

	a.logger.Info("shutting down server")
	return a.server.Shutdown(ctx)
}
