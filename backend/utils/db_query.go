package utils

import (
	"app/entities"
	userdto "app/service/dto/user_dto"

	"gorm.io/gorm"
)

/**
* Adds pagination to a db query from gorm.
**/
func AsPage(page int, limit int) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Limit(limit).Offset((page - 1) * limit)
	}
}

func ForUserBranch(user userdto.UserWithRoleAndPermissions) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		if user.Role.RoleType == entities.SuperAdmin {
			return db
		}

		return db.Where("branch_id = ?", user.BranchId)
	}
}
