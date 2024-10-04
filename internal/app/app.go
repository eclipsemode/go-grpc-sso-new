package app

import (
	"github.com/eclipsemode/go-grpc-sso-new/internal/app/grpc"
	"github.com/eclipsemode/go-grpc-sso-new/internal/config"
	"github.com/eclipsemode/go-grpc-sso-new/internal/services"
	"github.com/eclipsemode/go-grpc-sso-new/internal/storage/mongo"
	"go.uber.org/zap"
)

type App struct {
	GRPCSrv *grpc.App
}

func New(log *zap.SugaredLogger, cfg *config.Config) *App {
	mongoConfig := &config.MongoConfig{
		Uri:    cfg.Storage.Mongo.Uri,
		DbName: cfg.Storage.Mongo.DbName,
	}

	storage, err := mongo.NewStorage(mongoConfig, log)
	if err != nil {
		log.DPanic("new storage problem shit")
		panic(err)
	}

	authSvc := services.NewAuthService(log, cfg.TokenTTL, storage)

	grpcApp := grpc.New(log, cfg.GRPC.Port, authSvc)

	return &App{
		GRPCSrv: grpcApp,
	}
}
