package dto

import "github.com/google/uuid"

type DeleteBranchDto struct {
	ID uuid.UUID `param:"id" json:"id"`
}
