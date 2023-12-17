package entities

import (
	"database/sql/driver"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleTypeVal string

const (
	Common     RoleTypeVal = "COMMON"
	SuperAdmin RoleTypeVal = "SUPER_ADMIN"
)

func (ct *RoleTypeVal) Scan(value interface{}) error {
	*ct = RoleTypeVal(value.(string))
	return nil
}

func (ct RoleTypeVal) Value() (driver.Value, error) {
	return string(ct), nil
}

type Role struct {
	gorm.Model `tstype:"-"`
	ID         uuid.NullUUID `gorm:"type:uuid;primary_key;default:gen_random_uuid();" json:"id"`
	Name       string        `gorm:"type:varchar(255)" json:"name"`
	RoleType   RoleTypeVal   `gorm:"type:varchar(255)" json:"roleType"`
	BranchId   uuid.NullUUID `gorm:"type:uuid;" json:"branchId"`

	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`

	Permissions []Permission `gorm:"foreignKey:RoleId;references:ID" json:"permissions"`
	Branch      Branch       `gorm:"foreignKey:BranchId;reference:ID" json:"branch"`
}
