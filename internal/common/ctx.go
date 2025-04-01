package common

import (
	"context"
	"log/slog"
)

type ctxKey string

var logKey ctxKey = "logger"

func InjectLogger(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, logKey, logger)
}

func ExtractLogger(ctx context.Context) (*slog.Logger, bool) {
	log, ok := ctx.Value(logKey).(*slog.Logger)
	return log, ok
}
