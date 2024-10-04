package main

import (
	"github.com/eclipsemode/go-grpc-sso-new/internal/app"
	"github.com/eclipsemode/go-grpc-sso-new/internal/config"
	"github.com/eclipsemode/go-grpc-sso-new/internal/lib/logger"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad(config.EnvLocal)

	log := logger.New()

	r := mux.NewRouter()

	_ = r

	application := app.New(log, cfg)

	go application.GRPCSrv.MustRun()

	err := gracefulShutdown(log, application)
	if err != nil {
		log.Fatal("failed to gracefully shutdown the server", zap.Error(err))
	}
}

func gracefulShutdown(log *zap.SugaredLogger, app *app.App) error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	sign := <-c

	err := app.GRPCSrv.Stop()
	if err != nil {
		log.Errorw("graceful shutdown", "error", err)
		return err
	}

	log.Log(zap.InfoLevel, zap.String("signal", sign.String()))

	return nil
}
