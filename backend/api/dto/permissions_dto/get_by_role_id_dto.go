package permissionsdto

import "github.com/google/uuid"

type GetByRoleIdDTO struct {
	RoleId uuid.UUID `query:"roleId" validate:"required" json:"roleId"`
}
