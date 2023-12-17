package roledto

import (
	"app/entities"
	branchdto "app/service/dto/branch_dto"

	"github.com/samber/lo"
)

type RoleWithBranchDTO struct {
	RoleDTO `tstype:",extends,required"`
	Branch  branchdto.BranchDTO `json:"branch" tstype:"BranchDTO"`
}

func ToRoleWithBranchDTO(role entities.Role) RoleWithBranchDTO {
	return RoleWithBranchDTO{
		RoleDTO: ToRoleDTO(role),
		Branch:  branchdto.ToBranchDTO(role.Branch),
	}
}

func ToRoleWithBranchDTOs(roles []entities.Role) []RoleWithBranchDTO {
	return lo.Map(roles, func(role entities.Role, _ int) RoleWithBranchDTO {
		return ToRoleWithBranchDTO(role)
	})
}
