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

	App struct {
		Name    string `env:"APP_NAME"`
		Version string `env:"APP_VERSION"`
	}

	Http struct {
		Host string `env:"HTTP_HOST"`
		Port string `env:"HTTP_PORT"`
	}

	Pg struct {
		Host     string `env:"PG_HOST"`
		Port     int    `env:"PG_PORT"`
		Username string `env:"PG_USER"`
		Password string `env:"PG_PASS"`
		Name     string `env:"PG_NAME"`
	}
}

func New() *Config {
	config := new(Config)

	if err := cleanenv.ReadEnv(config); err != nil {
		slog.Error("error when reading env", slog.String("err", err.Error()))
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
