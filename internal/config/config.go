package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	Mongo struct {
		URI      string `yaml:"uri" env:"MONGO_URI" env-default:"mongodb://localhost:27017"`
		Username string `yaml:"username" env:"MONGO_USERNAME" env-default:"mongo-admin"`
		Password string `yaml:"password" env:"MONGO_PASSWORD"`
		Database string `yaml:"database" env:"MONGO_DATABASE" env-default:"items"`
	} `yaml:"mongo"`
	GRPC struct {
		Port int `yaml:"port" env:"GRPC_PORT" env-default:"50051"`
	} `yaml:"grpc_server"`
	Logs struct {
		Level string `yaml:"level" env:"LOGS_LEVEL" env-default:"INFO"`
		Path  string `yaml:"path" env:"LOGS_PATH" env-default:"logs.txt"`
	} `yaml:"logs"`
}

var Conf Config

func Init() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Couldn't open current working directory!", err)
	}
	err = cleanenv.ReadConfig(filepath.Join(wd, "config/config.yaml"), &Conf)
	if err != nil {
		log.Fatal("Error on reading config!", err)
	}
}
