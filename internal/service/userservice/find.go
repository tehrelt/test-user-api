package userservice

import (
	"context"
	"errors"
	"log/slog"

	"github.com/google/uuid"
	"github.com/tehrelt/test-users-api/internal/common"
	"github.com/tehrelt/test-users-api/internal/models"
	"github.com/tehrelt/test-users-api/internal/service"
	"github.com/tehrelt/test-users-api/internal/storage"
)

func (s *UserService) Find(ctx context.Context, id uuid.UUID) (*models.User, error) {
	fn := "userservice.Find"
	l, ok := common.ExtractLogger(ctx)
	if !ok {
		l = slog.Default()
	}
	l = l.With(slog.String("fn", fn))

	l.Info("finding user", slog.String("id", id.String()))
	user, err := s.provider.Find(ctx, id)
	if err != nil {
		if errors.Is(err, storage.ErrUserNotFound) {
			l.Info("user not found", slog.String("id", id.String()))
			return nil, service.ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}
