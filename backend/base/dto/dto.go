package dto

type WithFilterRequestDTO struct {
	Filter string `query:"filter" json:"filter"`
}

type PaginatedRequest struct {
	WithFilterRequestDTO `tstype:",extends,required"`
	Page                 int `query:"page" validate:"required" json:"page"`
	Limit                int `query:"limit" validate:"required" json:"limit"`
}

type PaginatedMeta struct {
	Total       int64 `json:"total"`
	LastPage    int32 `json:"lastPage"`
	CurrentPage int32 `json:"currentPage"`
	PerPage     int32 `json:"perPage"`
	Prev        int32 `json:"prev"`
	Next        int32 `json:"next"`
}

type Paginated[T any] struct {
	Data []T           `json:"data"`
	Meta PaginatedMeta `json:"meta"`
}
