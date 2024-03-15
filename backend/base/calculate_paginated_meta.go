package base

import (
	"math"
	"backend/base/dto"
)

func CalculatePaginationMeta(count int64, page int, limit int) dto.PaginatedMeta {
	lastPage := math.Ceil(float64(count) / float64(limit))
	var next int
	var prev int

	if page > 1 {
		prev = page - 1
	}

	if page < int(lastPage) {
		next = page + 1
	}

	return dto.PaginatedMeta{
		Total:       count,
		LastPage:    int32(lastPage),
		CurrentPage: int32(page),
		PerPage:     int32(limit),
		Prev:        int32(prev),
		Next:        int32(next),
	}
}
