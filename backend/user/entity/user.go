package entity

import (
	"time"

	branchEntity "backend/branch/entity"
	roleEntity "backend/role/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `tstype:"-"`
	ID         uuid.NullUUID  `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	FirstName  string         `gorm:"type:varchar(255)" json:"firstName"`
	LastName   string         `gorm:"type:varchar(255)" json:"lastName"`
	Email      string         `gorm:"type:varchar(255)" json:"email"`
	Password   string         `gorm:"type:varchar(255)" json:"password"`
	BranchId   uuid.NullUUID  `json:"branchId"`
	RoleId     uuid.NullUUID  `json:"roleId"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deletedAt"`

	Branch branchEntity.Branch `gorm:"foreignKey:BranchId;references:ID" json:"branch"`
	Role   roleEntity.Role     `gorm:"foreignKey:RoleId;references:ID" json:"role"`
}
