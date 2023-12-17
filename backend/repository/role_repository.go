package repository

import (
	"app/entities"
	userdto "app/service/dto/user_dto"
	"app/utils"

	"github.com/google/uuid"
	"github.com/samber/do"
	"gorm.io/gorm"
)

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(i *do.Injector) (*RoleRepository, error) {
	return &RoleRepository{
		db: do.MustInvoke[*gorm.DB](i),
	}, nil
}

func (rr *RoleRepository) GetAllRoles(user userdto.UserWithRoleAndPermissions, filter string) ([]entities.Role, error) {
	var roles []entities.Role
	err := rr.db.Table("roles").Scopes(
		withRoleFilter(filter),
		utils.ForUserBranch(user),
	).Preload("Branch").Find(&roles).Error
	return roles, err
}

func (rr *RoleRepository) GetRoleById(id string) (entities.Role, error) {
	var role entities.Role
	err := rr.db.Where("id = ?", id).Find(&role).Error
	return role, err
}

func (rr *RoleRepository) GetRoleWithPermisionsById(id uuid.NullUUID) (*entities.Role, error) {
	role := &entities.Role{
		ID: id,
	}
	err := rr.db.Model(&entities.Role{}).Find(role).Error
	return role, err
}

func (rr *RoleRepository) CreateRole(role entities.Role) (entities.Role, error) {
	err := rr.db.Create(&role).Error
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

func (rr *RoleRepository) DeleteById(user userdto.UserWithRoleAndPermissions, id uuid.UUID) error {
	return rr.db.Table("roles").Scopes(
		utils.ForUserBranch(user),
	).Delete(&entities.Role{
		ID: uuid.NullUUID{UUID: id, Valid: true},
	}).Error
}
