package dto

import (
	branchdto "backend/branch/dto"
	"backend/role/entity"

	"github.com/samber/lo"
)

type RoleWithBranchDTO struct {
	RoleDTO `tstype:",extends,required"`
	Branch  branchdto.BranchDTO `json:"branch" tstype:"BranchDTO"`
}

func ToRoleWithBranchDTO(role entity.Role) RoleWithBranchDTO {
	return RoleWithBranchDTO{
		RoleDTO: ToRoleDTO(role),
		Branch:  branchdto.ToBranchDTO(role.Branch),
	}
}

func ToRoleWithBranchDTOs(roles []entity.Role) []RoleWithBranchDTO {
	return lo.Map(roles, func(role entity.Role, _ int) RoleWithBranchDTO {
		return ToRoleWithBranchDTO(role)
	})
}
