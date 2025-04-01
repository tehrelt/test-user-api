package common

import (
	"context"
	"log/slog"
)

type ctxKey string

var LogKey ctxKey = "logger"

func InjectLogger(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, LogKey, logger)
}

func ExtractLogger(ctx context.Context) (*slog.Logger, bool) {
	log, ok := ctx.Value(LogKey).(*slog.Logger)
	return log, ok
}
