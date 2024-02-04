package utils

import (
	"math"
	"app/model"
)

func CalculatePaginationMeta(count int64, page int, limit int) model.PaginatedMeta {
	lastPage := math.Ceil(float64(count) / float64(limit))
	var next int
	var prev int

	if page > 1 {
		prev = page - 1
	}

	if page < int(lastPage) {
		next = page + 1
	}

	return model.PaginatedMeta{
		Total:       count,
		LastPage:    int32(lastPage),
		CurrentPage: int32(page),
		PerPage:     int32(limit),
		Prev:        int32(prev),
		Next:        int32(next),
	}
}
