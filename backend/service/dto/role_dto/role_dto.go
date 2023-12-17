package roledto

import (
	"app/entities"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type RoleDTO struct {
	ID       uuid.NullUUID        `json:"id"`
	Name     string               `json:"name"`
	RoleType entities.RoleTypeVal `json:"roleType" tstype:"RoleTypeVal"`
}

func ToRoleDTO(role entities.Role) RoleDTO {
	return RoleDTO{
		ID:       role.ID,
		Name:     role.Name,
		RoleType: role.RoleType,
	}
}

func ToRoleDTOs(rs []entities.Role) []RoleDTO {
	return lo.Map[entities.Role](rs, func(r entities.Role, _ int) RoleDTO {
		return ToRoleDTO(r)
	})
}
