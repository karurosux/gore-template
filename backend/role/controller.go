package role

import (
	"backend/base"
	constants "backend/base/contants"
	basedto "backend/base/dto"
	permissionentity "backend/permission/entity"
	"backend/role/dto"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/samber/do"
)

type Controller struct {
	roleService  *Service
	errorService *base.ErrorService
}

func NewController(i *do.Injector) (*Controller, error) {
	return &Controller{
		roleService:  do.MustInvoke[*Service](i),
		errorService: do.MustInvoke[*base.ErrorService](i),
	}, nil
}

func (rc *Controller) RegisterRoutes(e *echo.Echo) {
	writePermissions := base.PermissionCheck([]base.AllowedPermissions{
		{
			Cat: permissionentity.RoleManagement,
			Val: base.WRITE,
		},
	})
	readPermissions := base.PermissionCheck([]base.AllowedPermissions{
		{
			Cat: permissionentity.RoleManagement,
			Val: base.READ,
		},
	})
	e.GET(constants.RoleEndpoint, rc.HandleGetAllRoles, readPermissions).Name = "getAllRoles"
	e.GET(constants.RoleEndpoint+"/:id", rc.HandleGetRoleById, readPermissions).Name = "getRoleById"
	e.POST(constants.RoleEndpoint, rc.HandleCreateRole, writePermissions).Name = "createRole"
	e.DELETE(constants.RoleEndpoint+"/:id", rc.HandleDeleteRole, writePermissions).Name = "deleteRole"
}

func (rc *Controller) HandleGetAllRoles(c echo.Context) error {
	query := new(basedto.WithFilterRequestDTO)

	if err := c.Bind(query); err != nil {
		return rc.errorService.BadRequestError(err)
	}

	user := base.GetUserFromContext(c)
	res, err := rc.roleService.GetAllRoles(user, query.Filter)
	if err != nil {
		return rc.errorService.InternalServerError(err)
	}

	return c.JSON(200, res)
}

func (rc *Controller) HandleGetRoleById(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return rc.errorService.BadRequestError(echo.NewHTTPError(500, "Id is required."))
	}

	result, err := rc.roleService.GetRoleById(id)
	if err != nil {
		return echo.NewHTTPError(echo.ErrBadRequest.Code, err.Error())
	}

	return c.JSON(200, result)
}

func (rc *Controller) HandleCreateRole(c echo.Context) error {
	role := new(dto.CreateRoleDTO)

	if err := c.Bind(role); err != nil {
		return echo.NewHTTPError(echo.ErrBadRequest.Code, err.Error())
	}

	if err := c.Validate(role); err != nil {
		return echo.NewHTTPError(echo.ErrBadRequest.Code, err.Error())
	}

	user := base.GetUserFromContext(c)

	newRole, err := rc.roleService.CreateRole(dto.SCreateRoleDTO{
		Name:     role.Name,
		BranchId: user.BranchId.UUID,
	})
	if err != nil {
		return rc.errorService.InternalServerError(err)
	}

	return c.JSON(201, newRole)
}

func (rc *Controller) HandleDeleteRole(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return rc.errorService.BadRequestError(echo.NewHTTPError(500, "Id is required."))
	}

	uid := uuid.MustParse(id)
	user := base.GetUserFromContext(c)

	err := rc.roleService.DeleteById(user, uid)
	if err != nil {
		return rc.errorService.InternalServerError(err)
	}

	return c.NoContent(200)
}
