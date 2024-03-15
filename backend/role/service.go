package role

import (
	"backend/permission"
	roledto "backend/role/dto"
	"backend/role/entity"
	"backend/user/dto"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/samber/do"
)

type Service struct {
	roleRepository     *Repository
	permissionsService *permission.Service
}

func NewService(i *do.Injector) (*Service, error) {
	return &Service{
		roleRepository:     do.MustInvoke[*Repository](i),
		permissionsService: do.MustInvoke[*permission.Service](i),
	}, nil
}

func (rs *Service) GetAllRoles(user dto.UserWithRoleAndPermissions, filter string) ([]roledto.RoleWithBranchDTO, error) {
	roles, err := rs.roleRepository.GetAllRoles(user, filter)
	if err != nil {
		return []roledto.RoleWithBranchDTO{}, err
	}
	return roledto.ToRoleWithBranchDTOs(roles), nil
}

func (rs *Service) CreateRole(role roledto.SCreateRoleDTO) (roledto.RoleDTO, error) {
	newRole, err := rs.roleRepository.CreateRole(entity.Role{
		Name:     role.Name,
		BranchId: uuid.NullUUID{UUID: role.BranchId, Valid: true},
		RoleType: entity.Common,
	})
	if err != nil {
		return roledto.RoleDTO{}, err
	}
	return roledto.ToRoleDTO(newRole), nil
}

func (rs *Service) GetRoleWithPermissionsById(id uuid.NullUUID) (roledto.RoleWithPermissionsDTO, error) {
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

func (rs *Service) GetRoleById(id string) (roledto.RoleDTO, error) {
	role, err := rs.roleRepository.GetRoleById(id)
	return roledto.ToRoleDTO(role), err
}

func (rs *Service) DeleteById(user dto.UserWithRoleAndPermissions, id uuid.UUID) error {
	return rs.roleRepository.DeleteById(user, id)
}
