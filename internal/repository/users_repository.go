package repository

import (
	"context"
	"resume-analyzer/internal/models"
)

type UsersRepository interface {
	CreateUser(ctx context.Context, payload *models.CreateUserRequest) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}