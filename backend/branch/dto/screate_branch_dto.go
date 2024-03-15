package dto

type SCreateBranchDTO struct {
	Name    string `json:"name"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zipcode"`
	Country string `json:"country"`
}
