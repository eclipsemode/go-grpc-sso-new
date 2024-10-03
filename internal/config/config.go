package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"time"
)

type EnvType string

const (
	EnvLocal EnvType = "local"
	EnvProd  EnvType = "prod"
)

type Config struct {
	Env      string        `yaml:"env" env-default:"local"`
	TokenTTL time.Duration `yaml:"token_ttl" env-default:"1h"`
	Storage  StorageConfig `yaml:"storage"`
	GRPC     GRPCConfig    `yaml:"grpc"`
}

type StorageConfig struct {
	Mongo MongoConfig `yaml:"mongo"`
}

type MongoConfig struct {
	Uri    string `yaml:"uri"`
	DbName string `yaml:"db_name"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func MustLoad(envType EnvType) *Config {
	configPath := fmt.Sprintf("./config/%s.yaml", envType)

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatal(err)
	}

	return &cfg
}
