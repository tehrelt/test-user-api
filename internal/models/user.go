package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Email     string
	CreatedAt time.Time
	UpdatedAt *time.Time
}
