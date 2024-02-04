package userdto

import "github.com/google/uuid"

type CreateUserDTO struct {
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	RoleId    uuid.UUID `json:"roleId"`
}
