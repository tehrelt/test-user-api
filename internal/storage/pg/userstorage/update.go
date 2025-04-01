package userstorage

import (
	"context"
	"log/slog"

	"github.com/Masterminds/squirrel"
	"github.com/tehrelt/test-users-api/internal/common"
	"github.com/tehrelt/test-users-api/internal/models"
	"github.com/tehrelt/test-users-api/internal/storage"
	"github.com/tehrelt/test-users-api/internal/storage/pg"
)

func (us *UserStorage) Update(ctx context.Context, in *storage.UpdateUserDto) (result *models.User, err error) {

	fn := "userstorage.Update"

	log, ok := common.ExtractLogger(ctx)
	if !ok {
		log = slog.Default()
	}
	log = log.With(slog.String("fn", fn))

	log.Debug("updating user", slog.Any("in", in))
	builder := squirrel.Update(pg.USERS_TABLE).Where(squirrel.Eq{"id": in.Id.String()}).PlaceholderFormat(squirrel.Dollar)

	tx, err := us.pool.Begin(ctx)
	if err != nil {
		log.Error("failed to begin transaction", slog.Any("err", err))
		return nil, err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
			return
		}

		_ = tx.Commit(ctx)
	}()

	if in.FirstName != nil {
		builder = builder.Set("first_name", *in.FirstName)
	}

	if in.LastName != nil {
		builder = builder.Set("last_name", *in.LastName)
	}

	if in.Email != nil {
		builder = builder.Set("email", *in.Email)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		log.Error("failed to build query", slog.Any("err", err))
		return nil, err
	}

	u := new(userEntry)
	if err := tx.
		QueryRow(ctx, query, args...).
		Scan(&u.id, &u.firstName, &u.lastName, &u.email, &u.createdAt, &u.updatedAt); err != nil {

		log.Error("failed to update user", slog.Any("err", err))
		return nil, err
	}

	result, err = u.toModel()
	if err != nil {
		log.Error("failed to convert user entry to model", slog.Any("err", err))
		return nil, err
	}

	return result, nil
}
