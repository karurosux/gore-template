package service

import (
	"github.com/samber/do"
	"golang.org/x/crypto/bcrypt"
)

type PasswordService struct{}

func NewPasswordService(i *do.Injector) (*PasswordService, error) {
	return &PasswordService{}, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (ps *PasswordService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (ps *PasswordService) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
