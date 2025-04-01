package storage

import "github.com/google/uuid"

type CreateUserDto struct {
	FirstName string
	LastName  string
	Email     string
}

type UpdateUserDto struct {
	Id        uuid.UUID
	FirstName *string
	LastName  *string
	Email     *string
}
