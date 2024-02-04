package service

import (
	"app/entities"
	"app/repository"
	branchdto "app/service/dto/branch_dto"

	"github.com/samber/do"
)

type BranchService struct {
	BaseService[entities.Branch, branchdto.BranchDTO]
}

func NewBranchService(i *do.Injector) (*BranchService, error) {
	return &BranchService{
		BaseService: BaseService[entities.Branch, branchdto.BranchDTO]{
			repository: do.MustInvoke[repository.BranchRepository](i),
			toDTOFunc:  branchdto.ToBranchDTO,
			toDTOsFunc: branchdto.ToBranchDTOs,
		},
	}, nil
}
