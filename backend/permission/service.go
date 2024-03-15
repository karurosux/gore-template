package permission

import (
	"backend/permission/dto"
	"backend/permission/entity"

	"github.com/google/uuid"
	"github.com/samber/do"
)

type Service struct {
	permissionsRepository *Repository
}

func NewService(i *do.Injector) (*Service, error) {
	return &Service{
		permissionsRepository: do.MustInvoke[*Repository](i),
	}, nil
}

func (ps *Service) GetPermissionsByRoleId(id uuid.NullUUID) ([]dto.PermissionsDTO, error) {
	permission, err := ps.permissionsRepository.GetPermissionsByRoleId(id)
	if err != nil {
		return nil, err
	}
	return dto.ToPermissionDTOs(permission), nil
}

func (ps *Service) Create(permision dto.SCreatePermissionDTO) (dto.PermissionsDTO, error) {
	permission, err := ps.permissionsRepository.Create(entity.Permission{
		RoleId:   permision.RoleId,
		Category: permision.Category,
		Read:     permision.Read,
		Write:    permision.Write,
	})
	if err != nil {
		return dto.PermissionsDTO{}, err
	}

	return dto.ToPermissionDTO(permission), nil
}

func (ps *Service) DeleteById(id uuid.UUID) error {
	return ps.permissionsRepository.DeleteById(id)
}

func (ps *Service) PatchById(id uuid.UUID, write bool, read bool) error {
	return ps.permissionsRepository.PatchById(id, write, read)
}
