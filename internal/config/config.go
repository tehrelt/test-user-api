package config

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type EnvType string

const (
	EnvProd  EnvType = "prod"
	EnvDev   EnvType = "dev"
	EnvLocal EnvType = "local"
)

type Config struct {
	Env EnvType `env:"ENV"`

	Http struct {
		Host string `env:"HTTP_HOST"`
		Port string `env:"HTTP_PORT"`
	}

	Pg struct {
		Host     string `env:"PG_HOST"`
		Port     string `env:"PG_PORT"`
		Username string `env:"PG_USER"`
		Password string `env:"PG_PASS"`
	}
}

func New() *Config {
	config := new(Config)

	if err := cleanenv.ReadEnv(config); err != nil {
		slog.Error("error when reading env", slog.String("err", err))
		header := fmt.Sprintf("%s - %s", os.Getenv("APP_NAME"), os.Getenv("APP_VERSION"))

		usage := cleanenv.FUsage(os.Stdout, config, &header)
		usage()

		os.Exit(-1)
	}

	setupLogger(config)

	slog.Debug("read config", slog.Any("cfg", config))

	return config
}

func setupLogger(cfg *Config) {
	var log *slog.Logger

	switch cfg.Env {
	case EnvProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	case EnvDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	default:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	}

	slog.SetDefault(log)
}
