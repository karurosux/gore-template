package permissionsdto

import (
	"app/entities"

	"github.com/google/uuid"
)

type CreatePermissionDTO struct {
	Category entities.PermissionCategoryVal
	Read     bool
	Write    bool
	RoleId   uuid.UUID
}
