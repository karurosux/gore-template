package main

import (
	"backend/base/utils"
	branchentity "backend/branch/entity"
	permissionentity "backend/permission/entity"
	roleentity "backend/role/entity"
	userentity "backend/user/entity"

	"gorm.io/gorm"
)

func MigrateDb() *gorm.DB {
	db := utils.NewDb()

	db.AutoMigrate(branchentity.Branch{})
	db.AutoMigrate(permissionentity.Permission{})
	db.AutoMigrate(roleentity.Role{})
	db.AutoMigrate(userentity.User{})

	return db
}
