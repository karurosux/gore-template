package dto

import (
	"backend/permission/entity"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type PermissionsDTO struct {
	ID       uuid.UUID                    `json:"id"`
	Category entity.PermissionCategoryVal `json:"category" tstype:"PermissionCategoryVal"`
	Write    bool                         `json:"write"`
	Read     bool                         `json:"read"`
}

func ToPermissionDTO(p entity.Permission) PermissionsDTO {
	return PermissionsDTO{
		ID:       p.ID.UUID,
		Category: p.Category,
		Write:    p.Write,
		Read:     p.Read,
	}
}

func ToPermissionDTOs(p []entity.Permission) []PermissionsDTO {
	permissions := lo.Map[entity.Permission](p, func(permission entity.Permission, _ int) PermissionsDTO {
		return ToPermissionDTO(permission)
	})
	return permissions
}
