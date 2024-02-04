package customerdto

import (
	"app/entities"
	"time"

	"github.com/samber/lo"
)

type CustomerDTO struct {
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Birthdate time.Time `json:"birthdate"`
}

func ToCustomerDTO(c entities.Customer) CustomerDTO {
	return CustomerDTO{
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Email:     c.Email,
		Birthdate: c.Birthdate,
	}
}

func ToCustomerDTOs(cs []entities.Customer) []CustomerDTO {
	return lo.Map(cs, func(c entities.Customer, _ int) CustomerDTO {
		return ToCustomerDTO(c)
	})
}
