package repository

import (
	"app/entities"
	userdto "app/service/dto/user_dto"
	"app/utils"

	"github.com/samber/do"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(i *do.Injector) (*CustomerRepository, error) {
	return &CustomerRepository{
		db: do.MustInvoke[*gorm.DB](i),
	}, nil
}

func (cr *CustomerRepository) GetAllCustomers(user userdto.UserWithRoleAndPermissions) ([]entities.Customer, error) {
	var customers []entities.Customer
	err := cr.db.Table("customers").Scopes(utils.ForUserBranch(user)).Find(&customers).Error
	return customers, err
}

func (cr *CustomerRepository) GetPageByUser(user userdto.UserWithRoleAndPermissions, page int, limit int, filter string) ([]entities.Customer, error) {
	var cs []entities.Customer
	err := cr.db.Table("customers").Scopes(withCustomerFilter(filter), utils.AsPage(page, limit), utils.ForUserBranch(user)).Find(&cs).Error
	return cs, err
}

func (cr *CustomerRepository) GetCountByUser(user userdto.UserWithRoleAndPermissions, filter string) (int64, error) {
	var count int64
	error := cr.db.Table("customers").Scopes(withCustomerFilter(filter), utils.ForUserBranch(user)).Count(&count).Error
	return count, error
}

func withCustomerFilter(filter string) func(*gorm.DB) *gorm.DB {
	return func(d *gorm.DB) *gorm.DB {
		if filter != "" {
			return d.Where("first_name like ?", "%"+filter+"%").Or("last_name like ?", "%"+filter+"%").Or("email like ?", "%"+filter+"%")
		}

		return d
	}
}
