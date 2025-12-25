package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id uuid.UUID `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Email string `json:"email" db:"email"`
	PasswordHash string `json:"hash" db:"password_hash"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type CreateUserRequest struct {
	Username string `json:"username" db:"username" validate:"required,min=6,max=32"`
	Email string `json:"email" db:"email" validate:"required,email"`
	PasswordHash string `json:"hash" db:"password_hash" validate:"required,min=8"`
}