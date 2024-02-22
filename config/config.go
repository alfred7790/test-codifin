package config

import (
	"fmt"
	"github.com/jinzhu/configor"
	"log"
	"os"
)

type Config struct {
	Host      Host
	DB        DB
	DebugMode bool `env:"DEBUG_MODE" default:"true"`
}

type Host struct {
	Environment string
	Port        string `env:"HOST_PORT" default:"8080"`
}

type DB struct {
	Host     string `env:"DB_HOST" default:"db_products"`
	Port     string `env:"DB_PORT"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Name     string `env:"DB_NAME"`
	Retries  int    `env:"DB_RETRIES" default:"3"`
}

var config Config

func init() {
	environment := os.Getenv("APP_ENV")
	if environment == "" {
		environment = "develop"
	}

	fileName := fmt.Sprintf("config/yaml/config_%s.yaml", environment)
	err := configor.Load(&config, fileName)
	if err != nil {
		log.Fatalf("file (%s) can't be charged, reason: %s", fileName, err.Error())
	}
}

func GetConfig() *Config {
	return &config
}
