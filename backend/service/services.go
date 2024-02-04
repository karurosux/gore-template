package service

import (
	"github.com/samber/do"
)

func RegisterServices(i *do.Injector) {
	do.Provide(i, NewPermissionsService)
	do.Provide(i, NewUserService)
	do.Provide(i, NewRoleService)
	do.Provide(i, NewErrorService)
	do.Provide(i, NewPasswordService)
	do.Provide(i, NewAuthService)
	do.Provide(i, NewJwtService)
	do.Provide(i, NewBranchService)
	do.Provide(i, NewCustomerService)
}
