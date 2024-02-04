package middleware

import (
	"database/sql/driver"
	"app/entities"
	permissionsdto "app/service/dto/permissions_dto"

	"github.com/labstack/echo/v4"
)

type PermissionLevel string

func (ct *PermissionLevel) Scan(value interface{}) error {
	*ct = PermissionLevel(value.([]byte))
	return nil
}

func (ct PermissionLevel) Value() (driver.Value, error) {
	return string(ct), nil
}

const (
	WRITE PermissionLevel = "WRITE"
	READ  PermissionLevel = "READ"
)

type AllowedPermissions struct {
	Cat entities.PermissionCategoryVal
	Val PermissionLevel
}

func PermissionCheck(allowed []AllowedPermissions) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := GetUserFromContext(c)
			passed := false
			permissionsMap := map[string]permissionsdto.PermissionsDTO{}

			for i := 0; i < len(user.Role.Permissions); i++ {
				permission := (user.Role.Permissions)[i]
				permissionsMap[string(permission.Category)] = permission
			}

			for i := 0; i < len(allowed); i++ {
				curr := allowed[i]
				userPermission := permissionsMap[string(curr.Cat)]

				switch {
				case userPermission.Category == curr.Cat && userPermission.Read && curr.Val == READ:
					passed = true
				case userPermission.Category == curr.Cat && userPermission.Write && curr.Val == WRITE:
					passed = true
				}
				if passed {
					break
				}
			}

			if !passed {
				c.Error(echo.NewHTTPError(echo.ErrUnauthorized.Code, "Unauthorized"))
				return nil
			}

			return next(c)
		}
	}
}
