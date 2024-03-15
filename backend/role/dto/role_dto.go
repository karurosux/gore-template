package dto

import (
	"backend/role/entity"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type RoleDTO struct {
	ID       uuid.NullUUID      `json:"id"`
	Name     string             `json:"name"`
	RoleType entity.RoleTypeVal `json:"roleType" tstype:"RoleTypeVal"`
}

func ToRoleDTO(role entity.Role) RoleDTO {
	return RoleDTO{
		ID:       role.ID,
		Name:     role.Name,
		RoleType: role.RoleType,
	}
}

func ToRoleDTOs(rs []entity.Role) []RoleDTO {
	return lo.Map[entity.Role](rs, func(r entity.Role, _ int) RoleDTO {
		return ToRoleDTO(r)
	})
}
