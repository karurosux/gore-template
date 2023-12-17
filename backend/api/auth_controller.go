package api

import (
	authdto "app/api/dto/auth_dto"
	constants "app/contants"
	"app/service"
	"app/utils"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/samber/do"
)

type AuthController struct {
	authService  *service.AuthService
	errorService *service.ErrorService
}

func NewAuthController(i *do.Injector) (*AuthController, error) {
	return &AuthController{
		authService:  do.MustInvoke[*service.AuthService](i),
		errorService: do.MustInvoke[*service.ErrorService](i),
	}, nil
}

func (ac *AuthController) RegisterRoutes(e *echo.Echo) {
	e.POST(constants.AuthEndpoint+"/login", ac.login).Name = "login"
	e.POST(constants.AuthEndpoint+"/logout", ac.logout).Name = "logout"
}

func (ac *AuthController) login(c echo.Context) error {
	log.Println("Performing login.")

	creds := new(authdto.LoginBodyDTO)

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

func (ac *AuthController) logout(c echo.Context) error {
	utils.ClearSessionToken(c)
	c.Logger().Info("logout")
	return nil
}
