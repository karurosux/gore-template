package middleware

import (
	userdto "app/service/dto/user_dto"

	"github.com/labstack/echo/v4"
)

type AppContext struct {
	echo.Context
	User userdto.UserWithRoleAndPermissions
}

func AppendAppContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &AppContext{
			Context: c,
			User:    userdto.UserWithRoleAndPermissions{},
		}
		return next(cc)
	}
}

func GetUserFromContext(c echo.Context) userdto.UserWithRoleAndPermissions {
	cc := c.(*AppContext)
	return cc.User
}
