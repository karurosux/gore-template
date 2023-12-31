package userdto

import (
	"app/entities"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type UserDTO struct {
	ID        uuid.NullUUID `json:"id"`
	FirstName string        `json:"firstName"`
	LastName  string        `json:"lastName"`
	Email     string        `json:"email"`
	BranchId  uuid.NullUUID `json:"branchId"`
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"`
}

func ToUserDto(user entities.User) UserDTO {
	return UserDTO{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		BranchId:  user.BranchId,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToUserDTOs(u []entities.User) []UserDTO {
	userDTOs := lo.Map(u, func(cu entities.User, idx int) UserDTO {
		return ToUserDto(cu)
	})
	return userDTOs
}
