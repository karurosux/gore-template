package userdto

import "github.com/google/uuid"

type CreateUserDTO struct {
	FirstName string    `json:"firstName" validate:"required"`
	LastName  string    `json:"lastName" validate:"required"`
	Email     string    `json:"email" validate:"email,required"`
	Password  string    `json:"password" validate:"required"`
	RoleId    uuid.UUID `json:"roleId" validate:"required"`
}
