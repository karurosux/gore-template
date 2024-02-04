package service

import (
	"errors"

	"github.com/samber/do"
)

type AuthService struct {
	userServ        *UserService
	jwtService      *JwtService
	passwordService *PasswordService
}

func NewAuthService(i *do.Injector) (*AuthService, error) {
	return &AuthService{
		userServ:        do.MustInvoke[*UserService](i),
		jwtService:      do.MustInvoke[*JwtService](i),
		passwordService: do.MustInvoke[*PasswordService](i),
	}, nil
}

func (as *AuthService) Login(email string, password string) (string, error) {
	foundUser, err := as.userServ.FindByEmailWithPassword(email)
	if err != nil {
		return "", err
	}

	if !as.passwordService.CheckPasswordHash(password, foundUser.Password) {
		return "", errors.New("Unauthorized")
	}

	token, err := as.jwtService.GenerateToken(email)
	if err != nil {
		return "", err
	}

	return token, nil
}
