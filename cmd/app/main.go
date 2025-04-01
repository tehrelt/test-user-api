package main

import (
	"context"
	"flag"
	"log"
	"log/slog"

	"github.com/joho/godotenv"
	"github.com/tehrelt/test-users-api/internal/app"
)

var (
	local = flag.Bool("local", false, "run in local mode")
)

func main() {
	flag.Parse()

	if local != nil && *local {
		if err := godotenv.Load(); err != nil {
			log.Fatal("failed to load .env file")
		}
	}

	ctx := context.Background()

	instance, cleanup, err := app.New(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()

	if err := instance.Run(ctx); err != nil {
		slog.Error("failed to run app")
	}
}
