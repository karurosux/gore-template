package api

import (
	"app/model"

	"github.com/labstack/echo/v4"
	"github.com/samber/do"
)

func RegisterControllers(i *do.Injector) {
	do.Provide(i, NewUserController)
	do.Provide(i, NewBranchController)
	do.Provide(i, NewRoleController)
	do.Provide(i, NewAuthController)
	do.Provide(i, NewPermissionsController)
	do.Provide(i, NewCustomerController)
}

func RegisterRoutes(e *echo.Echo, i *do.Injector) {
	controllers := &[]model.Controller{
		do.MustInvoke[*UserController](i),
		do.MustInvoke[*BranchController](i),
		do.MustInvoke[*RoleController](i),
		do.MustInvoke[*AuthController](i),
		do.MustInvoke[*PermissionsController](i),
		do.MustInvoke[*CustomerController](i),
	}

	for i := 0; i < len(*controllers); i++ {
		(*controllers)[i].RegisterRoutes(e)
	}
}
