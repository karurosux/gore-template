package auth

import (
	"errors"
	"backend/base"
	"backend/user"

	"github.com/samber/do"
)

type Service struct {
	service         *user.Service
	jwtService      *base.JwtService
	passwordService *base.PasswordService
}

func NewService(i *do.Injector) (*Service, error) {
	return &Service{
		service:         do.MustInvoke[*user.Service](i),
		jwtService:      do.MustInvoke[*base.JwtService](i),
		passwordService: do.MustInvoke[*base.PasswordService](i),
	}, nil
}

func (as *Service) Login(email string, password string) (string, error) {
	foundUser, err := as.service.FindByEmailWithPassword(email)
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
