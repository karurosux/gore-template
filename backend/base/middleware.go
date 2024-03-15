package base

import (
	"backend/user/dto"

	"github.com/labstack/echo/v4"
)

type BackendContext struct {
	echo.Context
	User dto.UserWithRoleAndPermissions
}

func AppendBackendContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &BackendContext{
			Context: c,
			User:    dto.UserWithRoleAndPermissions{},
		}
		return next(cc)
	}
}

func GetUserFromContext(c echo.Context) dto.UserWithRoleAndPermissions {
	cc := c.(*BackendContext)
	return cc.User
}
