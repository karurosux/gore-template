package permissionsdto

import (
	"app/entities"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type PermissionsDTO struct {
	ID       uuid.UUID                      `json:"id"`
	Category entities.PermissionCategoryVal `json:"category" tstype:"PermissionCategoryVal"`
	Write    bool                           `json:"write"`
	Read     bool                           `json:"read"`
}

func ToPermissionDTO(p entities.Permission) PermissionsDTO {
	return PermissionsDTO{
		ID:       p.ID.UUID,
		Category: p.Category,
		Write:    p.Write,
		Read:     p.Read,
	}
}

func ToPermissionDTOs(p []entities.Permission) []PermissionsDTO {
	permissions := lo.Map[entities.Permission](p, func(permission entities.Permission, _ int) PermissionsDTO {
		return ToPermissionDTO(permission)
	})
	return permissions
}
