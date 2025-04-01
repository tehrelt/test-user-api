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

func (serv *UserService) Create(ctx context.Context, in *service.CreateUserDto) (*models.User, error) {

	fn := "userservice.Create"
	l, ok := common.ExtractLogger(ctx)
	if !ok {
		l = slog.Default()
	}
	l = l.With(slog.String("fn", fn))

	newUser := &storage.CreateUserDto{
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Email:     in.Email,
	}

	l.Info("creating user")
	user, err := serv.saver.Create(ctx, newUser)
	if err != nil {
		if errors.Is(err, storage.ErrUserAlreadyExists) {
			l.Info("user already exists")
			return nil, service.ErrUserAlreadyExists
		}

		l.Error("failed to create user", slog.String("err", err.Error()))
		return nil, err
	}

	return user, nil
}
