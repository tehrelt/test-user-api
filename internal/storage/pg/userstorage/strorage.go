package userstorage

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/tehrelt/test-users-api/internal/models"
)

type UserStorage struct {
	pool *pgxpool.Pool
}

type userEntry struct {
	id        string
	firstName string
	lastName  string
	email     string
	createdAt time.Time
	updatedAt *time.Time
}

func (u *userEntry) toModel() (*models.User, error) {
	id, err := uuid.Parse(u.id)
	if err != nil {
		return nil, err
	}
	return &models.User{
		Id:        id,
		FirstName: u.firstName,
		LastName:  u.lastName,
		Email:     u.email,
		CreatedAt: u.createdAt,
		UpdatedAt: u.updatedAt,
	}, nil
}

func NewUserStorage(pool *pgxpool.Pool) *UserStorage {
	return &UserStorage{
		pool: pool,
	}
}
