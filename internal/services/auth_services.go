package services

import (
	"context"
	"log"
	"resume-analyzer/internal/errors/apperrors"
	"resume-analyzer/internal/models"
	"resume-analyzer/internal/repository"
	"resume-analyzer/internal/utils"
)

type AuthService struct {
	usersRepo repository.UsersRepository
}

func NewAuthService(usersRepo repository.UsersRepository) *AuthService {
	return &AuthService{usersRepo: usersRepo}
}

func (s *AuthService) CreateUser(ctx context.Context, payload *models.CreateUserRequest) error {
	// check if user exists
	_, err := s.usersRepo.GetUserByEmail(ctx, payload.Email)
	if err == nil {
		return apperrors.NewResourceAlreadyExistsError("user")
	}
	// hash password
	hash, err := utils.HashPassword(payload.PasswordHash)
	if err != nil {
		log.Println("Hashing Error: ", err)
		return apperrors.NewInternalServerError()
	}

	payload.PasswordHash = hash
	// create user
	createUserErr := s.usersRepo.CreateUser(ctx, payload)
	if createUserErr != nil {
		return apperrors.NewInternalServerError()
	}

	return nil
}
