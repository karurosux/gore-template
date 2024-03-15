package dto

import (
	"backend/branch/entity"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type BranchDTO struct {
	ID        uuid.NullUUID `json:"id"`
	Name      string        `json:"name"`
	City      string        `json:"city"`
	State     string        `json:"state"`
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"`
}

func ToBranchDTO(branch entity.Branch) BranchDTO {
	return BranchDTO{
		ID:        branch.ID,
		Name:      branch.Name,
		City:      branch.City,
		State:     branch.State,
		CreatedAt: branch.CreatedAt,
		UpdatedAt: branch.UpdatedAt,
	}
}

func ToBranchDTOs(branches []entity.Branch) []BranchDTO {
	return lo.Map(branches, func(branch entity.Branch, idx int) BranchDTO {
		return ToBranchDTO(branch)
	})
}
