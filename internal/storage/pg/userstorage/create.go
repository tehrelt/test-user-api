package userstorage

import (
	"context"
	"errors"
	"log/slog"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx"
	"github.com/tehrelt/test-users-api/internal/common"
	"github.com/tehrelt/test-users-api/internal/models"
	"github.com/tehrelt/test-users-api/internal/storage"
	"github.com/tehrelt/test-users-api/internal/storage/pg"
)

func (us *UserStorage) Create(ctx context.Context, in *storage.CreateUserDto) (user *models.User, err error) {

	fn := "userstorage.Create"

	log, ok := common.ExtractLogger(ctx)
	if !ok {
		log = slog.Default()
	}
	log = log.With(slog.String("fn", fn))

	log.Debug("creating user", slog.Any("in", in))

	query, args, err := squirrel.Insert(pg.USERS_TABLE).
		Columns("last_name", "first_name", "email").
		Values(in.LastName, in.FirstName, in.Email).
		Suffix("RETURNING id, last_name, first_name, email, created_at, updated_at").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	u := new(userEntry)
	tx, err := us.pool.Begin(ctx)
	if err != nil {
		log.Error("failed to begin transaction", slog.String("err", err.Error()))
		return nil, err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
			return
		}

		_ = tx.Commit(ctx)
	}()

	if err = tx.QueryRow(ctx, query, args...).
		Scan(&u.id, &u.lastName, &u.firstName, &u.email, &u.createdAt, &u.updatedAt); err != nil {
		var pgErr pgx.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				log.Warn("user already exists")
				return nil, storage.ErrUserAlreadyExists
			}
			log.Error("failed to create user on pg side", slog.Any("err", pgErr))
			return nil, err
		}

		log.Error("failed to create user", slog.String("err", err.Error()))
		return nil, err
	}

	user, err = u.toModel()
	if err != nil {
		log.Error("failed to convert user entry to model", slog.String("err", err.Error()))
		return nil, err
	}

	return user, nil
}
