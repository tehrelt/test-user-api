package app

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/wire"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/tehrelt/test-users-api/internal/config"
	"github.com/tehrelt/test-users-api/internal/service/userservice"
	"github.com/tehrelt/test-users-api/internal/storage/pg/userstorage"
	"github.com/tehrelt/test-users-api/internal/transport/http"
)

func New(ctx context.Context) (*App, func(), error) {
	panic(wire.Build(
		newApp,

		http.New,

		userservice.New,
		wire.Bind(new(userservice.UserSaver), new(*userstorage.UserStorage)),
		wire.Bind(new(userservice.UserProvider), new(*userstorage.UserStorage)),

		userstorage.New,

		_pg,
		config.New,
	))
}

func _pg(ctx context.Context, cfg *config.Config) (*pgxpool.Pool, func(), error) {
	host := cfg.Pg.Host
	port := cfg.Pg.Port
	user := cfg.Pg.Username
	pass := cfg.Pg.Password
	name := cfg.Pg.Name

	cs := fmt.Sprintf(`postgres://%s:%s@%s:%d/%s?sslmode=disable`, user, pass, host, port, name)

	pool, err := pgxpool.Connect(ctx, cs)
	if err != nil {
		return nil, func() {}, err
	}

	slog.Debug("connecting to database", slog.String("conn", cs))
	t := time.Now()
	if err := pool.Ping(ctx); err != nil {
		slog.Error("failed to connect to database", slog.String("err", err.Error()), slog.String("conn", cs))
		return nil, func() { pool.Close() }, err
	}
	slog.Info("connected to database", slog.String("ping", fmt.Sprintf("%2.fs", time.Since(t).Seconds())))

	return pool, func() { pool.Close() }, nil
}
