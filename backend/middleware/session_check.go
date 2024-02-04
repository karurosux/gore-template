package middleware

import (
	"app/service"
	"app/utils"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/samber/do"
	"github.com/samber/lo"
)

var PUBLIC_ROUTES []string = []string{
	"/api/v1/auth/login",
}

type SessionCheck struct {
	userServ     *service.UserService
	errorService *service.ErrorService
	jwtService   *service.JwtService
}

func NewSessionCheck(i *do.Injector) (*SessionCheck, error) {
	return &SessionCheck{
		userServ:     do.MustInvoke[*service.UserService](i),
		errorService: do.MustInvoke[*service.ErrorService](i),
		jwtService:   do.MustInvoke[*service.JwtService](i),
	}, nil
}

func (sc *SessionCheck) SessionCheckMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var url string = c.Request().URL.String()

		if sc.isPublicRoute(url) {
			// No need to pass for the other validations.
			return next(c)
		}

		sessionToken := utils.GetSessionToken(c)
		if lo.IsEmpty(sessionToken) {
			c.Error(sc.errorService.UnauthorizedError(nil))
			return nil
		}

		_, claims, err := sc.jwtService.ValidateToken(sessionToken)
		if err != nil {
			c.Error(sc.errorService.UnauthorizedError(err))
			return nil
		}

		foundUser, err := sc.userServ.FindByEmailWithRoleAndPermissions(claims.Email)
		if err != nil {
			c.Error(sc.errorService.UnauthorizedError(err))
			return nil
		}

		cc := c.(*AppContext)
		cc.User = foundUser

		return next(cc)
	}
}

func (sc *SessionCheck) isPublicRoute(url string) bool {
	return !strings.Contains(url, "/api") || lo.Some(PUBLIC_ROUTES, []string{url})
}
