package dto

import (
	"backend/permission/entity"

	"github.com/google/uuid"
)

type SCreatePermissionDTO struct {
	Category entity.PermissionCategoryVal
	Read     bool
	Write    bool
	RoleId   uuid.UUID
}
