package repository

import (
	"app/entities"

	"github.com/samber/do"
	"gorm.io/gorm"
)

type BranchRepository struct {
	BaseRepository[entities.Branch]
}

func NewBranchRepository(i *do.Injector) (BranchRepository, error) {
	return BranchRepository{
		BaseRepository: BaseRepository[entities.Branch]{
			db: do.MustInvoke[*gorm.DB](i),
			filterFunc: func(filter string) func(*gorm.DB) *gorm.DB {
				return func(d *gorm.DB) *gorm.DB {
					if filter != "" {
						return d.Where("name like ?", "%"+filter+"%").Or("city like ?", "%"+filter+"%").Or("state like ?", "%"+filter+"%")
					}

					return d
				}
			},
			patchFunc: func(body entities.Branch) ([]string, entities.Branch) {
				return []string{
						"name",
					}, entities.Branch{
						Name: body.Name,
					}
			},
		},
	}, nil
}
