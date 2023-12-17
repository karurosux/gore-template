package userdto

import "app/entities"

type UserWithPasswordDTO struct {
	UserDTO  `tstype:",extends,required"`
	Password string `json:"password"`
}

func ToUserWithPassword(u entities.User) UserWithPasswordDTO {
	ud := ToUserDto(u)
	return UserWithPasswordDTO{
		UserDTO:  ud,
		Password: u.Password,
	}
}
