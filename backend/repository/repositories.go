package repository

import "github.com/samber/do"

func RegisterRepositories(i *do.Injector) {
	do.Provide(i, NewPermissionsRepository)
	do.Provide(i, NewUserRepository)
	do.Provide(i, NewBranchRepository)
	do.Provide(i, NewRoleRepository)
	do.Provide(i, NewCustomerRepository)
}
