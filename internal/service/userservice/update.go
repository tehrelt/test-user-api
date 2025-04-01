package userservice

import (
	"context"
	"errors"
	"log/slog"

	"github.com/tehrelt/test-users-api/internal/common"
	"github.com/tehrelt/test-users-api/internal/models"
	"github.com/tehrelt/test-users-api/internal/service"
	"github.com/tehrelt/test-users-api/internal/storage"
)

func (s *UserService) Update(ctx context.Context, in *service.UpdateUserDto) (*models.User, error) {

	fn := "userservice.Update"
	l, ok := common.ExtractLogger(ctx)
	if !ok {
		l = slog.Default()
	}
	l = l.With(slog.String("fn", fn))

	l.Info("updating user", slog.Any("in", in))

	user, err := s.saver.Update(ctx, &storage.UpdateUserDto{
		Id:        in.Id,
		FirstName: in.FirstName,
		LastName:  in.LastName,
	})
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			l.Info("user not found", slog.Any("err", err))
			return nil, service.ErrUserNotFound
		}

		l.Error("error updating user", slog.Any("err", err))
		return nil, err
	}

	return user, nil
}
