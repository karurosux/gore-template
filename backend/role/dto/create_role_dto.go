package dto

type CreateRoleDTO struct {
	Name string `json:"name" validate:"required"`
}
