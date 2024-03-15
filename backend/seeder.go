package main

import (
	"backend/base"
	branchentity "backend/branch/entity"
	permissionentity "backend/permission/entity"
	roleentity "backend/role/entity"
	userentity "backend/user/entity"
	"log"

	"github.com/google/uuid"
)

func Seed() {
	log.Println("Loading environment variables...")

	log.Println("Seeding database...")
	db := MigrateDb()

	branchId := uuid.New()

	db.Create(&branchentity.Branch{
		ID:      uuid.NullUUID{UUID: branchId, Valid: true},
		Name:    "HQ",
		City:    "San Diego",
		State:   "CA",
		Country: "USA",
		ZipCode: "92101",
	})
	log.Printf("Branch %s seeded successfully.", branchId.String())

	roleId := uuid.New()
	db.Create(&roleentity.Role{
		ID:       uuid.NullUUID{UUID: roleId, Valid: true},
		Name:     "Super Admin",
		RoleType: roleentity.SuperAdmin,
		BranchId: uuid.NullUUID{UUID: branchId, Valid: true},
	})
	log.Printf("Role %s seeded successfully.", roleId.String())

	// Seed permissions
	db.Create(&permissionentity.Permission{
		RoleId:   roleId,
		Category: permissionentity.UserManagement,
		Write:    true,
		Read:     true,
	})
	db.Create(&permissionentity.Permission{
		RoleId:   roleId,
		Category: permissionentity.RoleManagement,
		Write:    true,
		Read:     true,
	})
	db.Create(&permissionentity.Permission{
		RoleId:   roleId,
		Category: permissionentity.CustomerManagement,
		Write:    true,
		Read:     true,
	})
	db.Create(&permissionentity.Permission{
		RoleId:   roleId,
		Category: permissionentity.BranchManagement,
		Write:    true,
		Read:     true,
	})
	log.Printf("Permissions seeded successfully.")

	hashedPassword, err := base.HashPassword("Pass123!")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Creating user with role id %s \n", roleId)

	db.Create(&userentity.User{
		FirstName: "Super",
		LastName:  "Admin",
		Email:     "admin@admin.com",
		Password:  hashedPassword,
		RoleId:    uuid.NullUUID{UUID: roleId, Valid: true},
		BranchId:  uuid.NullUUID{UUID: branchId, Valid: true},
	})

	log.Println("Database seeded successfully.")
}
