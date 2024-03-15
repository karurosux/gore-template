package branch

import (
	"backend/base"
	"backend/branch/dto"
	"backend/branch/entity"

	"github.com/samber/do"
)

type Service struct {
	base.Service[entity.Branch, dto.BranchDTO]
}

func NewService(i *do.Injector) (*Service, error) {
	return &Service{
		Service: base.Service[entity.Branch, dto.BranchDTO]{
			Repository: do.MustInvoke[Repository](i),
			ToDTOFunc:  dto.ToBranchDTO,
			ToDTOsFunc: dto.ToBranchDTOs,
		},
	}, nil
}
