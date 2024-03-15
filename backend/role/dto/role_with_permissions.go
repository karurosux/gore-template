package dto

import (
	permissionsdto "backend/permission/dto"
	"backend/role/entity"
)

type RoleWithPermissionsDTO struct {
	RoleDTO     `tstype:",extends,required"`
	Permissions []permissionsdto.PermissionsDTO `json:"permissions" tstype:"PermissionsDTO[]"`
}

func ToRoleWithPermissionsDTO(r entity.Role) RoleWithPermissionsDTO {
	rdto := ToRoleDTO(r)
	return RoleWithPermissionsDTO{
		RoleDTO:     rdto,
		Permissions: []permissionsdto.PermissionsDTO{},
	}
}
