package dto

import (
	"github.com/google/uuid"
)

type SCreateRoleDTO struct {
	Name     string
	BranchId uuid.UUID
}
