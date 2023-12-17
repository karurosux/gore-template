package roledto

type CreateRoleDTO struct {
	Name string `json:"name" validate:"required"`
}
