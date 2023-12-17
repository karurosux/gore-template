package service

import (
	"app/repository"
	branchdto "app/service/dto/branch_dto"

	"github.com/google/uuid"
	"github.com/samber/do"
)

type BranchService struct {
	branchRepository *repository.BranchRepository
}

func NewBranchService(i *do.Injector) (*BranchService, error) {
	return &BranchService{
		branchRepository: do.MustInvoke[*repository.BranchRepository](i),
	}, nil
}

func (bs *BranchService) GetAllBranches() []branchdto.BranchDTO {
	return branchdto.ToBranchDTOs(bs.branchRepository.GetAllBranches())
}

func (bs *BranchService) GetBranchId(id uuid.UUID) (branchdto.BranchDTO, error) {
	branch, err := bs.branchRepository.GetBranchById(id)
	if err != nil {
		return branchdto.BranchDTO{}, err
	}
	return branchdto.ToBranchDTO(branch), nil
}
