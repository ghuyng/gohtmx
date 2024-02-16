package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/ghuyng/gohtmx/internal/config"
	"github.com/ghuyng/gohtmx/internal/web/app"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// Load config from file and environment
	cfg, err := config.Load("env/web.env")
	if err != nil {
		panic(err)
	}

	// Setup logger
	loggerCfg := zap.NewProductionConfig()
	loggerCfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger := zap.Must(loggerCfg.Build()).Sugar()
	defer logger.Sync()

	// Create new app
	app := app.NewApp(cfg, logger)

	errc := make(chan error, 1)
	go func() {
		errc <- app.Start()
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	select {
	case err := <-errc:
		logger.Errorw("failed to serve", zap.Error(err))
	case sig := <-sigs:
		logger.Infow("terminating", zap.String("signal", sig.String()))
	}

	if err := app.Shutdown(); err != nil {
		logger.Fatal(err.Error())
	}
	logger.Info("app shutdown successfully")
}
