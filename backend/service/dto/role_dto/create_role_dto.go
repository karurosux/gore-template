package roledto

import (
	"github.com/google/uuid"
)

type CreateRoleDTO struct {
	Name     string
	BranchId uuid.UUID
}
