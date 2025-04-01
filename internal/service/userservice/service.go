package userservice

import (
	"context"

	"github.com/google/uuid"
	"github.com/tehrelt/test-users-api/internal/models"
	"github.com/tehrelt/test-users-api/internal/storage"
)

type UserSaver interface {
	Create(ctx context.Context, in *storage.CreateUserDto) (*models.User, error)
	Update(ctx context.Context, in *storage.UpdateUserDto) (*models.User, error)
}

type UserProvider interface {
	Find(ctx context.Context, id uuid.UUID) (*models.User, error)
}

type UserService struct {
	saver    UserSaver
	provider UserProvider
}

func New(saver UserSaver, provider UserProvider) *UserService {
	return &UserService{
		saver:    saver,
		provider: provider,
	}
}
