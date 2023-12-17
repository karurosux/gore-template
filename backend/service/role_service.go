package service

import (
	"app/entities"
	"app/repository"
	roledto "app/service/dto/role_dto"
	userdto "app/service/dto/user_dto"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/samber/do"
)

type RoleService struct {
	roleRepository     *repository.RoleRepository
	permissionsService *PermissionsService
}

func NewRoleService(i *do.Injector) (*RoleService, error) {
	return &RoleService{
		roleRepository:     do.MustInvoke[*repository.RoleRepository](i),
		permissionsService: do.MustInvoke[*PermissionsService](i),
	}, nil
}

func (rs *RoleService) GetAllRoles(user userdto.UserWithRoleAndPermissions, filter string) ([]roledto.RoleWithBranchDTO, error) {
	roles, err := rs.roleRepository.GetAllRoles(user, filter)
	if err != nil {
		return []roledto.RoleWithBranchDTO{}, err
	}
	return roledto.ToRoleWithBranchDTOs(roles), nil
}

func (rs *RoleService) CreateRole(role roledto.CreateRoleDTO) (roledto.RoleDTO, error) {
	newRole, err := rs.roleRepository.CreateRole(entities.Role{
		Name:     role.Name,
		BranchId: uuid.NullUUID{UUID: role.BranchId, Valid: true},
		RoleType: entities.Common,
	})
	if err != nil {
		return roledto.RoleDTO{}, err
	}
	return roledto.ToRoleDTO(newRole), nil
}

func (rs *RoleService) GetRoleWithPermissionsById(id uuid.NullUUID) (roledto.RoleWithPermissionsDTO, error) {
	role, err := rs.roleRepository.GetRoleById(id.UUID.String())

	if err != nil {
		return roledto.RoleWithPermissionsDTO{}, err
	}

	if (role.ID == uuid.NullUUID{}) {
		return roledto.RoleWithPermissionsDTO{}, echo.NewHTTPError(echo.ErrNotFound.Code, "Role not found")
	}

	rolep := roledto.ToRoleWithPermissionsDTO(role)
	permissions, err := rs.permissionsService.GetPermissionsByRoleId(id)

	if err != nil {
		return roledto.RoleWithPermissionsDTO{}, err
	}

	rolep.Permissions = permissions

	return rolep, nil
}

func (rs *RoleService) GetRoleById(id string) (roledto.RoleDTO, error) {
	role, err := rs.roleRepository.GetRoleById(id)
	return roledto.ToRoleDTO(role), err
}

func (rs *RoleService) DeleteById(user userdto.UserWithRoleAndPermissions, id uuid.UUID) error {
	return rs.roleRepository.DeleteById(user, id)
}
