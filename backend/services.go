package main

import (
	"backend/auth"
	"backend/base"
	"backend/branch"
	"backend/permission"
	"backend/role"
	"backend/user"

	"github.com/samber/do"
)

func RegisterServices(i *do.Injector) {
	do.Provide(i, base.NewPasswordService)
	do.Provide(i, base.NewErrorService)
	do.Provide(i, base.NewJwtService)
	do.Provide(i, role.NewService)
	do.Provide(i, permission.NewService)
	do.Provide(i, user.NewService)
	do.Provide(i, auth.NewService)
	do.Provide(i, branch.NewService)
}
