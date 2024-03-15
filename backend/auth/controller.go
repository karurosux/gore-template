package auth

import (
	"log"
	"backend/auth/dto"
	"backend/base"
	constants "backend/base/contants"
	"backend/base/utils"

	"github.com/labstack/echo/v4"
	"github.com/samber/do"
)

type Controller struct {
	authService  *Service
	errorService *base.ErrorService
}

func NewController(i *do.Injector) (*Controller, error) {
	return &Controller{
		authService:  do.MustInvoke[*Service](i),
		errorService: do.MustInvoke[*base.ErrorService](i),
	}, nil
}

func (ac *Controller) RegisterRoutes(e *echo.Echo) {
	e.POST(constants.AuthEndpoint+"/login", ac.login).Name = "login"
	e.POST(constants.AuthEndpoint+"/logout", ac.logout).Name = "logout"
}

func (ac *Controller) login(c echo.Context) error {
	log.Println("Performing login.")

	creds := new(dto.LoginBodyDTO)

	if err := c.Bind(creds); err != nil {
		return ac.errorService.InternalServerError(err)
	}

	if err := c.Validate(creds); err != nil {
		return ac.errorService.BadRequestError(err)
	}

	token, err := ac.authService.Login(creds.Email, creds.Password)
	if err != nil {
		return ac.errorService.UnauthorizedError(err)
	}

	utils.SetSessionToken(c, token)

	return c.NoContent(200)
}

func (ac *Controller) logout(c echo.Context) error {
	utils.ClearSessionToken(c)
	c.Logger().Info("logout")
	return nil
}
