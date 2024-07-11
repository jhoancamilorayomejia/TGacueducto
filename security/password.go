package security

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// ErrInvalidPassword error when invalid password
var ErrInvalidPassword = errors.New("invalid password")

// EncryptPassword return encrypted password
func EncryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// VerifyPassword verify that the password is correct
func VerifyPassword(hashPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		return ErrInvalidPassword
	}

	return nil
}
