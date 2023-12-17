package permissionsdto

import (
	"app/entities"

	"github.com/google/uuid"
)

type CreatePermissionDTO struct {
	Category entities.PermissionCategoryVal `json:"category" validate:"required"`
	Read     bool                           `json:"read"`
	Write    bool                           `json:"write"`
	RoleId   uuid.UUID                      `json:"roleId" validate:"required"`
}
