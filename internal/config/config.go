package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	// Основной конфиг
	Env string `yaml:"env" env-default:"local"` //Также есть опцияenv-required,
	// при ее использовании нужно будет явно указать окуржение
	ConfigDatabase `yaml:"db"`
	HTTPServer     `yaml:"http-server"` //Встраивание структуры:
	// все публичные поля и методы встроенной структуры доступны здесь как собственные
}
type ConfigDatabase struct {
	Host     string `yaml:"host" env-default:"localhost"`
	Port     string `yaml:"port" env-default:"5432"`
	User     string `yaml:"user" env-default:"postgres"`
	Password string `yaml:"password" env-default:"postgres"`
	Database string `yaml:"database" env-default:"postgres"`
}

type HTTPServer struct {
	Address     string        `yaml:"port" default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" default:"5s"`
	IdleTimeout time.Duration `yaml:"idleTimeout" default:"60s"`
}

func MustLoad() Config {
	//configPath := os.Getenv("CONFIG_PATH")

	configPath := "config/config.yaml"
	if configPath == "" {

		log.Fatal("CONFIG_PATH environment variable not set")
	}
	//check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file does not exist: %s", configPath)
	}

	var config Config
	err := cleanenv.ReadConfig(configPath, &config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}
