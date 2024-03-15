package role

import (
	"backend/base"
	"backend/role/entity"
	userdto "backend/user/dto"

	"github.com/google/uuid"
	"github.com/samber/do"
	"gorm.io/gorm"
)

type Repository struct {
	base.Repository[entity.Role]
}

func NewRepository(i *do.Injector) (*Repository, error) {
	return &Repository{
		Repository: base.Repository[entity.Role]{
			Db: do.MustInvoke[*gorm.DB](i),
			FilterFunc: func(filter string) func(*gorm.DB) *gorm.DB {
				return func(db *gorm.DB) *gorm.DB {
					if filter != "" {
						return db.Where("name like ?", "%"+filter+"%")
					}
					return db
				}
			},
			PatchFunc: func(body entity.Role) ([]string, entity.Role) {
				return []string{}, entity.Role{}
			},
		},
	}, nil
}

func (rr *Repository) GetAllRoles(user userdto.UserWithRoleAndPermissions, filter string) ([]entity.Role, error) {
	var roles []entity.Role
	err := rr.Db.Table("roles").Scopes(
		withRoleFilter(filter),
		base.ForUserBranch(user),
	).Preload("Branch").Find(&roles).Error
	return roles, err
}

func (rr *Repository) GetRoleById(id string) (entity.Role, error) {
	var role entity.Role
	err := rr.Db.Where("id = ?", id).Find(&role).Error
	return role, err
}

func (rr *Repository) GetRoleWithPermisionsById(id uuid.NullUUID) (*entity.Role, error) {
	role := &entity.Role{
		ID: id,
	}
	err := rr.Db.Model(&entity.Role{}).Find(role).Error
	return role, err
}

func (rr *Repository) CreateRole(role entity.Role) (entity.Role, error) {
	err := rr.Db.Create(&role).Error
	return role, err
}

func withRoleFilter(filter string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if filter != "" {
			return db.Where("name like ?", "%"+filter+"%")
		}
		return db
	}
}

func (rr *Repository) DeleteById(user userdto.UserWithRoleAndPermissions, id uuid.UUID) error {
	return rr.Db.Table("roles").Scopes(
		base.ForUserBranch(user),
	).Delete(&entity.Role{
		ID: uuid.NullUUID{UUID: id, Valid: true},
	}).Error
}
