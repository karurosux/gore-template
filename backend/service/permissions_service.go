package service

import (
	"app/entities"
	"app/repository"
	permissionsdto "app/service/dto/permissions_dto"

	"github.com/google/uuid"
	"github.com/samber/do"
)

type PermissionsService struct {
	permissionsRepository *repository.PermissionsRepository
}

func NewPermissionsService(i *do.Injector) (*PermissionsService, error) {
	return &PermissionsService{
		permissionsRepository: do.MustInvoke[*repository.PermissionsRepository](i),
	}, nil
}

func (ps *PermissionsService) GetPermissionsByRoleId(id uuid.NullUUID) ([]permissionsdto.PermissionsDTO, error) {
	permission, err := ps.permissionsRepository.GetPermissionsByRoleId(id)
	if err != nil {
		return nil, err
	}
	return permissionsdto.ToPermissionDTOs(permission), nil
}

func (ps *PermissionsService) Create(permision permissionsdto.CreatePermissionDTO) (permissionsdto.PermissionsDTO, error) {
	permission, err := ps.permissionsRepository.Create(entities.Permission{
		RoleId:   permision.RoleId,
		Category: permision.Category,
		Read:     permision.Read,
		Write:    permision.Write,
	})

	if err != nil {
		return permissionsdto.PermissionsDTO{}, err
	}

	return permissionsdto.ToPermissionDTO(permission), nil
}

func (ps *PermissionsService) DeleteById(id uuid.UUID) error {
	return ps.permissionsRepository.DeleteById(id)
}

func (ps *PermissionsService) PatchById(id uuid.UUID, write bool, read bool) error {
	return ps.permissionsRepository.PatchById(id, write, read)
}
