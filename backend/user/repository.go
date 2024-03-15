package user

import (
	"backend/base"
	"backend/user/entity"

	"github.com/samber/do"
	"gorm.io/gorm"
)

type Repository struct {
	base.Repository[entity.User]
}

func NewRepository(i *do.Injector) (Repository, error) {
	return Repository{
		Repository: base.Repository[entity.User]{
			Db: do.MustInvoke[*gorm.DB](i),
			FilterFunc: func(filter string) func(*gorm.DB) *gorm.DB {
				return func(d *gorm.DB) *gorm.DB {
					if filter != "" {
						return d.Where("first_name like ?", "%"+filter+"%").Or("last_name like ?", "%"+filter+"%").Or("email like ?", "%"+filter+"%")
					}

					return d
				}
			},
			PatchFunc: func(body entity.User) ([]string, entity.User) {
				return []string{"first_name", "last_name"}, entity.User{
					FirstName: body.FirstName,
					LastName:  body.LastName,
				}
			},
		},
	}, nil
}

func (ur Repository) FindByEmail(email string) (entity.User, error) {
	var user entity.User
	err := ur.Repository.Db.Where("email = ?", email).First(&user).Error
	return user, err
}
