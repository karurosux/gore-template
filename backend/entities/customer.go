package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model `tstype:"-"`
	ID         uuid.NullUUID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	FirstName  string        `json:"firstName" gorm:"type:varchar(255)"`
	LastName   string        `json:"lastName" gorm:"type:varchar(255)"`
	Email      string        `json:"email" gorm:"type:varchar(255)"`
	Birthdate  time.Time     `json:"birthdate"`
	BranchId   uuid.UUID     `json:"branchId"`
	CreatorId  uuid.UUID     `json:"creatorId"`

	Branch  Branch `json:"branch" gorm:"foreignKey:BranchId;references:ID"`
	Creator User   `json:"creator" gorm:"foreignKey:CreatorId;references:ID"`
}
