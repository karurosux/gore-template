package base

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/samber/do"
)

type ErrorService struct{}

func NewErrorService(i *do.Injector) (*ErrorService, error) {
	return &ErrorService{}, nil
}

func (es *ErrorService) InternalServerError(err error) *echo.HTTPError {
	return echo.NewHTTPError(echo.ErrInternalServerError.Code, err.Error())
}

func (es *ErrorService) NotFoundError(err error) *echo.HTTPError {
	return echo.NewHTTPError(echo.ErrNotFound.Code, err.Error())
}

func (es *ErrorService) BadRequestError(err error) *echo.HTTPError {
	return echo.NewHTTPError(echo.ErrBadRequest.Code, err.Error())
}

func (es *ErrorService) UnauthorizedError(err error) *echo.HTTPError {
	if err != nil {
		log.Println(err)
	}
	return echo.NewHTTPError(echo.ErrUnauthorized.Code, "Unauthorized")
}
