package postgres

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"resume-analyzer/internal/config"
	"resume-analyzer/internal/errors/apperrors"
	"resume-analyzer/internal/models"
)

type UserRepository struct {
	db *config.DB
}

func NewUserRepository(db *config.DB) *UserRepository {
	return &UserRepository{db: db}
}

const (
	CREATEUSER = `
	INSERT INTO users (username, email, password_hash)
	VALUES ($1, $2, $3)
	`
	GETUSERBYEMAIL= `
	SELECT * FROM users 
	WHERE email = $1
	`
)

func (r *UserRepository) CreateUser(ctx context.Context, payload *models.CreateUserRequest) error {
	_, err := r.db.ExecContext(ctx, CREATEUSER, payload.Username, payload.Email, payload.PasswordHash)
	if err != nil {
		log.Println("Error creating user in db: ",  err)
		return apperrors.NewInternalServerError()
	}
	return nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User 

	err := r.db.GetContext(ctx, GETUSERBYEMAIL, email)
	switch {
	case err == nil:
		return &user, nil
	case errors.Is(err, sql.ErrNoRows):
		return nil, apperrors.NewNotFoundError("user")
	default:
		log.Println("Db error: ", err)
		return nil, apperrors.NewInternalServerError()
	}
}

