package dto

import (
	"backend/permission/entity"

	"github.com/google/uuid"
)

type CreatePermissionDTO struct {
	Category entity.PermissionCategoryVal `json:"category" validate:"required"`
	Read     bool                         `json:"read"`
	Write    bool                         `json:"write"`
	RoleId   uuid.UUID                    `json:"roleId" validate:"required"`
}
