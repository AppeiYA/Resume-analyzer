package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	cost := bcrypt.DefaultCost

	hash, hashErr := bcrypt.GenerateFromPassword([]byte(password), cost)
	if hashErr != nil {
		return "", hashErr
	}
	return string(hash), nil
}

func CheckPassword(password string, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
