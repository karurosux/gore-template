package roledto

import (
	"app/entities"
	permissionsdto "app/service/dto/permissions_dto"
)

type RoleWithPermissionsDTO struct {
	RoleDTO     `tstype:",extends,required"`
	Permissions []permissionsdto.PermissionsDTO `json:"permissions" tstype:"PermissionsDTO[]"`
}

func ToRoleWithPermissionsDTO(r entities.Role) RoleWithPermissionsDTO {
	rdto := ToRoleDTO(r)
	return RoleWithPermissionsDTO{
		RoleDTO:     rdto,
		Permissions: []permissionsdto.PermissionsDTO{},
	}
}
