package mongo

import (
	"context"
	"fmt"
	"github.com/eclipsemode/go-grpc-sso-new/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type Storage struct {
	Db *mongo.Database
}

func NewStorage(cfg *config.MongoConfig, zap *zap.SugaredLogger) (*Storage, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().
		ApplyURI(fmt.Sprintf("%s/%s", cfg.Uri, cfg.DbName)).
		SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	zap.Info("Successfully connected to MongoDB")

	return &Storage{
		Db: client.Database(cfg.DbName),
	}, nil
}
