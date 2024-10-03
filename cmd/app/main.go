package main

import (
	"github.com/eclipsemode/go-grpc-sso-new/internal/app"
	"github.com/eclipsemode/go-grpc-sso-new/internal/config"
	"github.com/eclipsemode/go-grpc-sso-new/internal/lib/logger"
	"github.com/gorilla/mux"
)

func main() {
	cfg := config.MustLoad(config.EnvLocal)

	log := logger.New()

	r := mux.NewRouter()

	_ = r

	application := app.New(log, cfg)

	_ = application

	log.Info("App initialized successfully")
}
