package dto

type PaginatedRequest struct {
	WithFilterRequestDTO `tstype:",extends,required"`
	Page                 int `query:"page" validate:"required" json:"page"`
	Limit                int `query:"limit" validate:"required" json:"limit"`
}
