package main

import (
	"backend/auth"
	"backend/base/model"
	"backend/branch"
	"backend/permission"
	"backend/role"
	"backend/user"

	"github.com/labstack/echo/v4"
	"github.com/samber/do"
)

func RegisterControllers(i *do.Injector) {
	do.Provide(i, user.NewController)
	do.Provide(i, role.NewController)
	do.Provide(i, auth.NewController)
	do.Provide(i, permission.NewController)
	do.Provide(i, branch.NewController)
}

func RegisterRoutes(e *echo.Echo, i *do.Injector) {
	controllers := &[]model.Controller{
		do.MustInvoke[*user.Controller](i),
		do.MustInvoke[*role.Controller](i),
		do.MustInvoke[*auth.Controller](i),
		do.MustInvoke[*permission.Controller](i),
		do.MustInvoke[*branch.Controller](i),
	}

	for i := 0; i < len(*controllers); i++ {
		(*controllers)[i].RegisterRoutes(e)
	}
}
