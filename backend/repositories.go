package main

import (
	"backend/branch"
	"backend/permission"
	"backend/role"
	"backend/user"

	"github.com/samber/do"
)

func RegisterRepositories(i *do.Injector) {
	do.Provide(i, branch.NewRepository)
	do.Provide(i, role.NewRepository)
	do.Provide(i, user.NewRepository)
	do.Provide(i, permission.NewRepository)
}
