package repository

import (
	"app/entities"

	"github.com/google/uuid"
	"github.com/samber/do"
	"gorm.io/gorm"
)

type BranchRepository struct {
	db *gorm.DB
}

func NewBranchRepository(i *do.Injector) (*BranchRepository, error) {
	return &BranchRepository{
		db: do.MustInvoke[*gorm.DB](i),
	}, nil
}

func (br *BranchRepository) GetAllBranches() []entities.Branch {
	var branches []entities.Branch
	br.db.Find(&branches)
	return branches
}

func (br *BranchRepository) GetBranchById(id uuid.UUID) (entities.Branch, error) {
	var branch entities.Branch
	err := br.db.First(&branch, id).Error
	return branch, err
}
