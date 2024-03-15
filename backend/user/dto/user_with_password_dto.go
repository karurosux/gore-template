package dto

import (
	"backend/user/entity"
)

type UserWithPasswordDTO struct {
	UserDTO  `tstype:",extends,required"`
	Password string `json:"password"`
}

func ToUserWithPassword(u entity.User) UserWithPasswordDTO {
	ud := ToUserDTO(u)
	return UserWithPasswordDTO{
		UserDTO:  ud,
		Password: u.Password,
	}
}
