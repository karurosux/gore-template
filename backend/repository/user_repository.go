package repository

import (
	"app/entities"
	userdto "app/service/dto/user_dto"
	"app/utils"

	"github.com/google/uuid"
	"github.com/samber/do"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(i *do.Injector) (*UserRepository, error) {
	return &UserRepository{
		db: do.MustInvoke[*gorm.DB](i),
	}, nil
}

func (ur *UserRepository) GetByEmail(email string) (entities.User, error) {
	var user entities.User
	err := ur.db.Where("email = ?", email).First(&user).Error
	return user, err
}

func (ur *UserRepository) GetPageByUser(user userdto.UserWithRoleAndPermissions, page int, limit int, filter string) ([]entities.User, error) {
	var users []entities.User
	r := ur.db.Table("users").Scopes(
		withUserFilter(filter),
		utils.AsPage(page, limit),
		utils.ForUserBranch(user),
	).Preload("Role").Find(&users)
	return users, r.Error
}

func (ur *UserRepository) GetCountByUser(user userdto.UserWithRoleAndPermissions, filter string) (int64, error) {
	var count int64
	r := ur.db.Table("users").Scopes(withUserFilter(filter), utils.ForUserBranch(user)).Count(&count)
	return count, r.Error
}

func (ur *UserRepository) CreateUser(user entities.User) (entities.User, error) {
	err := ur.db.Table("users").Create(&user).Error
	return user, err
}

func (ur *UserRepository) DeleteById(user userdto.UserWithRoleAndPermissions, id uuid.UUID) error {
	return ur.db.Table("users").Scopes(utils.ForUserBranch(user)).Delete(&entities.User{
		ID: uuid.NullUUID{UUID: id, Valid: true},
	}).Error
}

func (ur *UserRepository) GetById(user userdto.UserWithRoleAndPermissions, id uuid.UUID) (entities.User, error) {
	var foundUser entities.User
	err := ur.db.Table("users").Scopes(utils.ForUserBranch(user)).Where("id = ?", id).First(&foundUser).Error
	return foundUser, err
}

func withUserFilter(filter string) func(*gorm.DB) *gorm.DB {
	return func(d *gorm.DB) *gorm.DB {
		if filter != "" {
			return d.Where("first_name like ?", "%"+filter+"%").Or("last_name like ?", "%"+filter+"%").Or("email like ?", "%"+filter+"%")
		}

		return d
	}
}
