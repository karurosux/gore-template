package branchdto

type CreateBranchDTO struct {
	Name    string `json:"name" validate:"required"`
	City    string `json:"city" validate:"required"`
	State   string `json:"state" validate:"required"`
	Country string `json:"country" validate:"required"`
	ZipCode string `json:"zipcode" validate:"required"`
}
