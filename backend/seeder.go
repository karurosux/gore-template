package main

import (
	"app/entities"
	"app/service"
	"app/utils"
	"log"

	"github.com/google/uuid"
)

func Seed() {
	log.Println("Loading environment variables...")

	log.Println("Seeding database...")
	db := utils.MigrateDb()

	branchId := uuid.New()

	db.Create(&entities.Branch{
		ID:      uuid.NullUUID{UUID: branchId, Valid: true},
		Name:    "HQ",
		City:    "San Diego",
		State:   "CA",
		Country: "USA",
		ZipCode: "92101",
	})
	log.Printf("Branch %s seeded successfully.", branchId.String())

	roleId := uuid.New()
	db.Create(&entities.Role{
		ID:       uuid.NullUUID{UUID: roleId, Valid: true},
		Name:     "Super Admin",
		RoleType: entities.SuperAdmin,
		BranchId: uuid.NullUUID{UUID: branchId, Valid: true},
	})
	log.Printf("Role %s seeded successfully.", roleId.String())

	// Seed permissions
	db.Create(&entities.Permission{
		RoleId:   roleId,
		Category: entities.UserManagement,
		Write:    true,
		Read:     true,
	})
	db.Create(&entities.Permission{
		RoleId:   roleId,
		Category: entities.RoleManagement,
		Write:    true,
		Read:     true,
	})
	db.Create(&entities.Permission{
		RoleId:   roleId,
		Category: entities.CustomerManagement,
		Write:    true,
		Read:     true,
	})
	db.Create(&entities.Permission{
		RoleId:   roleId,
		Category: entities.BranchManagement,
		Write:    true,
		Read:     true,
	})
	log.Printf("Permissions seeded successfully.")

	hashedPassword, err := service.HashPassword("Pass123!")

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Creating user with role id %s \n", roleId)

	db.Create(&entities.User{
		FirstName: "Super",
		LastName:  "Admin",
		Email:     "admin@admin.com",
		Password:  hashedPassword,
		RoleId:    uuid.NullUUID{UUID: roleId, Valid: true},
		BranchId:  uuid.NullUUID{UUID: branchId, Valid: true},
	})

	log.Println("Database seeded successfully.")
}
