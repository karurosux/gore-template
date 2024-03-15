package middleware

import (
	"backend/base"
	"backend/base/utils"
	"backend/user"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/samber/do"
	"github.com/samber/lo"
)

var PUBLIC_ROUTES []string = []string{
	"/api/v1/auth/login",
}

type SessionCheck struct {
	userServ     *user.Service
	errorService *base.ErrorService
	jwtService   *base.JwtService
}

func NewSessionCheck(i *do.Injector) (*SessionCheck, error) {
	return &SessionCheck{
		userServ:     do.MustInvoke[*user.Service](i),
		errorService: do.MustInvoke[*base.ErrorService](i),
		jwtService:   do.MustInvoke[*base.JwtService](i),
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

		cc := c.(*base.BackendContext)
		cc.User = foundUser

		return next(cc)
	}
}

func (sc *SessionCheck) isPublicRoute(url string) bool {
	return !strings.Contains(url, "/api") || lo.Some(PUBLIC_ROUTES, []string{url})
}
