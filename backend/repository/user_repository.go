package repository

import (
	"app/entities"

	"github.com/samber/do"
	"gorm.io/gorm"
)

type UserRepository struct {
	BaseRepository[entities.User]
}

func NewUserRepository(i *do.Injector) (UserRepository, error) {
	return UserRepository{
		BaseRepository: BaseRepository[entities.User]{
			db: do.MustInvoke[*gorm.DB](i),
			filterFunc: func(filter string) func(*gorm.DB) *gorm.DB {
				return func(d *gorm.DB) *gorm.DB {
					if filter != "" {
						return d.Where("first_name like ?", "%"+filter+"%").Or("last_name like ?", "%"+filter+"%").Or("email like ?", "%"+filter+"%")
					}

					return d
				}
			},
			patchFunc: func(body entities.User) ([]string, entities.User) {
				return []string{"first_name", "last_name"}, entities.User{
					FirstName: body.FirstName,
					LastName:  body.LastName,
				}
			},
		},
	}, nil
}

func (ur UserRepository) FindByEmail(email string) (entities.User, error) {
	var user entities.User
	err := ur.db.Where("email = ?", email).First(&user).Error
	return user, err
}
