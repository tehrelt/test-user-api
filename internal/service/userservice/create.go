package userservice

import (
	"context"
	"log/slog"

	"github.com/tehrelt/test-users-api/internal/common"
	"github.com/tehrelt/test-users-api/internal/models"
	"github.com/tehrelt/test-users-api/internal/service"
	"github.com/tehrelt/test-users-api/internal/storage"
)

func (service *UserService) Create(ctx context.Context, in *service.CreateUserDto) (*models.User, error) {

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
	user, err := service.saver.Create(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return user, nil
}
