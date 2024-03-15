package permission

import (
	"backend/permission/entity"

	"github.com/google/uuid"
	"github.com/samber/do"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(i *do.Injector) (*Repository, error) {
	return &Repository{
		db: do.MustInvoke[*gorm.DB](i),
	}, nil
}

func (pr *Repository) GetPermissionsByRoleId(id uuid.NullUUID) ([]entity.Permission, error) {
	permissions := []entity.Permission{}
	err := pr.db.Where("role_id = ?", id).Order("category DESC").Find(&permissions).Error
	return permissions, err
}

func (pr *Repository) Create(permission entity.Permission) (entity.Permission, error) {
	err := pr.db.Create(&permission).Error
	return permission, err
}

func (pr *Repository) DeleteById(id uuid.UUID) error {
	return pr.db.Delete(&entity.Permission{
		ID: uuid.NullUUID{UUID: id, Valid: true},
	}).Error
}

func (pr *Repository) PatchById(id uuid.UUID, write bool, read bool) error {
	return pr.db.Model(&entity.Permission{}).Select("read", "write").Where("id = ?", id).UpdateColumns(entity.Permission{
		Read:  read,
		Write: write,
	}).Error
}
