package middleware

import (
	"app/entities"

	"github.com/labstack/echo/v4"
)

func SuperAdminOnly() func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := GetUserFromContext(c)
			passed := user.Role.RoleType == entities.SuperAdmin

			if !passed {
				c.Error(echo.NewHTTPError(echo.ErrUnauthorized.Code, "Unauthorized"))
				return nil
			}

			return next(c)
		}
	}
}
