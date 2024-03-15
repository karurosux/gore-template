package entity

import (
	"database/sql/driver"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PermissionCategoryVal string

func (ct *PermissionCategoryVal) Scan(value interface{}) error {
	*ct = PermissionCategoryVal(value.(string))
	return nil
}

func (ct PermissionCategoryVal) Value() (driver.Value, error) {
	return string(ct), nil
}

const (
	UserManagement     PermissionCategoryVal = "USER_MANAGEMENT"
	RoleManagement     PermissionCategoryVal = "ROLE_MANAGEMENT"
	CustomerManagement PermissionCategoryVal = "CUSTOMER_MANAGEMENT"
	BranchManagement   PermissionCategoryVal = "BRANCH_MANAGEMENT"
)

type Permission struct {
	gorm.Model `tstype:"-"`
	ID         uuid.NullUUID         `gorm:"type:uuid;primary_key;default:gen_random_uuid();" json:"id"`
	RoleId     uuid.UUID             `gorm:"type:uuid;" json:"roleId"`
	Category   PermissionCategoryVal `gorm:"type:varchar(255)" json:"category"`
	Write      bool                  `gorm:"type:boolean" json:"write"`
	Read       bool                  `gorm:"type:boolean" json:"read"`
}
