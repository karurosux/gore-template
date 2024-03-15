package branch

import (
	"backend/base"
	"backend/branch/entity"

	"github.com/samber/do"
	"gorm.io/gorm"
)

type Repository struct {
	base.Repository[entity.Branch]
}

func NewRepository(i *do.Injector) (Repository, error) {
	return Repository{
		Repository: base.Repository[entity.Branch]{
			Db: do.MustInvoke[*gorm.DB](i),
			FilterFunc: func(filter string) func(*gorm.DB) *gorm.DB {
				return func(d *gorm.DB) *gorm.DB {
					if filter != "" {
						return d.Where("name like ?", "%"+filter+"%").Or("city like ?", "%"+filter+"%").Or("state like ?", "%"+filter+"%")
					}

					return d
				}
			},
			PatchFunc: func(body entity.Branch) ([]string, entity.Branch) {
				return []string{
						"name",
					}, entity.Branch{
						Name: body.Name,
					}
			},
		},
	}, nil
}
