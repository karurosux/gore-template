package userdto

import "github.com/google/uuid"

type DeleteUserDto struct {
	ID uuid.UUID `param:"id" json:"id"`
}
