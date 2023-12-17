package repository

import (
	"app/entities"

	"github.com/google/uuid"
	"github.com/samber/do"
	"gorm.io/gorm"
)

type PermissionsRepository struct {
	db *gorm.DB
}

func NewPermissionsRepository(i *do.Injector) (*PermissionsRepository, error) {
	return &PermissionsRepository{
		db: do.MustInvoke[*gorm.DB](i),
	}, nil
}

func (pr *PermissionsRepository) GetPermissionsByRoleId(id uuid.NullUUID) ([]entities.Permission, error) {
	permissions := []entities.Permission{}
	err := pr.db.Where("role_id = ?", id).Order("category DESC").Find(&permissions).Error
	return permissions, err
}

func (pr *PermissionsRepository) Create(permission entities.Permission) (entities.Permission, error) {
	err := pr.db.Create(&permission).Error
	return permission, err
}

func (pr *PermissionsRepository) DeleteById(id uuid.UUID) error {
	return pr.db.Delete(&entities.Permission{
		ID: uuid.NullUUID{UUID: id, Valid: true},
	}).Error
}

func (pr *PermissionsRepository) PatchById(id uuid.UUID, write bool, read bool) error {
	return pr.db.Model(&entities.Permission{}).Select("read", "write").Where("id = ?", id).UpdateColumns(entities.Permission{
		Read:  read,
		Write: write,
	}).Error
}
