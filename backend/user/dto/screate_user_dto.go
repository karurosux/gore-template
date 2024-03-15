package dto

import "github.com/google/uuid"

type SCreateUserDTO struct {
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	RoleId    uuid.UUID `json:"roleId"`
}
