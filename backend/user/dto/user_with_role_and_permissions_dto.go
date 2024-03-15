package dto

import (
	roledto "backend/role/dto"
	"backend/user/entity"

	"github.com/samber/lo"
)

type UserWithRoleAndPermissions struct {
	UserDTO `tstype:",extends,required"`
	Role    roledto.RoleWithPermissionsDTO `json:"role" tstype:"RoleWithPermissionsDTO"`
}

func ToUserWithRoleAndPermissions(u entity.User, r roledto.RoleWithPermissionsDTO) UserWithRoleAndPermissions {
	return UserWithRoleAndPermissions{
		UserDTO: ToUserDTO(u),
		Role:    r,
	}
}

func ToUsersWithRoleAndPermissions(u []entity.User) []UserWithRoleAndPermissions {
	return lo.Map(u, func(u entity.User, idx int) UserWithRoleAndPermissions {
		return UserWithRoleAndPermissions{
			UserDTO: ToUserDTO(u),
			Role:    roledto.ToRoleWithPermissionsDTO(u.Role),
		}
	})
}
