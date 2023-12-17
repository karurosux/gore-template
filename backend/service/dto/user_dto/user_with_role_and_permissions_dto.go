package userdto

import (
	"app/entities"
	roledto "app/service/dto/role_dto"

	"github.com/samber/lo"
)

type UserWithRoleAndPermissions struct {
	UserDTO `tstype:",extends,required"`
	Role    roledto.RoleWithPermissionsDTO `json:"role" tstype:"RoleWithPermissionsDTO"`
}

func ToUserWithRoleAndPermissions(u entities.User, r roledto.RoleWithPermissionsDTO) UserWithRoleAndPermissions {
	return UserWithRoleAndPermissions{
		UserDTO: ToUserDto(u),
		Role:    r,
	}
}

func ToUsersWithRoleAndPermissions(u []entities.User) []UserWithRoleAndPermissions {
	return lo.Map(u, func(u entities.User, idx int) UserWithRoleAndPermissions {
		return UserWithRoleAndPermissions{
			UserDTO: ToUserDto(u),
			Role:    roledto.ToRoleWithPermissionsDTO(u.Role),
		}
	})
}
