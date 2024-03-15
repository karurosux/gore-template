package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Branch struct {
	gorm.Model `tstype:"-"`
	ID         uuid.NullUUID  `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Name       string         `gorm:"type:varchar(255)" json:"name"`
	City       string         `gorm:"type:varchar(255)" json:"city"`
	State      string         `gorm:"type:varchar(255)" json:"province"`
	ZipCode    string         `gorm:"type:varchar(255)" json:"zipCode"`
	Country    string         `gorm:"type:varchar(255)" json:"country"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deletedAt"`

	// TODO: Set user entity
	// Users []User `gorm:"foreignKey:BranchId;references:ID" json:"users"`
}
