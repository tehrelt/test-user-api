package userstorage

import (
	"context"
	"log/slog"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/tehrelt/test-users-api/internal/common"
	"github.com/tehrelt/test-users-api/internal/models"
	"github.com/tehrelt/test-users-api/internal/storage"
)

func (us *UserStorage) Find(ctx context.Context, id uuid.UUID) (*models.User, error) {

	fn := "userstorage.Find"

	log, ok := common.ExtractLogger(ctx)
	if ok {
		log = slog.Default()
	}

	log = log.With(slog.String("fn", fn))
	log.Debug("finding user", slog.String("id", id.String()))

	query, args, err := squirrel.Select("id", "last_name", "first_name", "email", "created_at", "updated_at").
		From("users").
		Where(squirrel.Eq{"id": id.String()}).
		ToSql()

	if err != nil {
		log.Error("failed to build query", slog.Any("err", err))
		return nil, err
	}

	user := new(userEntry)
	if err := us.pool.QueryRow(ctx, query, args...).
		Scan(&user.id, &user.lastName, &user.firstName, &user.email, &user.createdAt, &user.updatedAt); err != nil {
		if err == pgx.ErrNoRows {
			log.Debug("user not found", slog.String("id", id.String()))
			return nil, storage.ErrUserNotFound
		}

		log.Error("failed to find user", slog.String("id", id.String()), slog.Any("err", err))

		return nil, err

	}

	u, err := user.toModel()
	if err != nil {
		log.Error("failed to convert user entry to model", slog.String("id", id.String()), slog.Any("err", err))
		return nil, err
	}

	return u, nil
}
