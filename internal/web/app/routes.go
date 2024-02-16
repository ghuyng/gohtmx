package app

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func (a *App) setupRoutes() error {
	a.router.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:      true,
		LogStatus:   true,
		LogMethod:   true,
		LogRemoteIP: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			a.logger.Infow("request",
				zap.String("uri", v.URI),
				zap.Int("status", v.Status),
				zap.String("method", v.Method),
				zap.String("remote_ip", v.RemoteIP),
			)
			return nil
		},
	}))
	a.router.Use(middleware.Recover())
	a.router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: a.cfg.Server.AllowOrigins,
		AllowMethods: a.cfg.Server.AllowMethods,
	}))

	a.router.Static("/assets", "internal/web/view/assets")
	a.router.GET("/users", a.userHandler.GetUsers)
	return nil
}
