package dto

import "github.com/google/uuid"

type PermissionByIdDTO struct {
	ID uuid.UUID `param:"id" validate:"required"`
}
