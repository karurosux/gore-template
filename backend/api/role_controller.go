package api

import (
	"app/api/dto"
	roledto "app/api/dto/role_dto"
	constants "app/contants"
	"app/entities"
	"app/middleware"
	mmidleware "app/middleware"
	"app/service"
	sroledto "app/service/dto/role_dto"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/samber/do"
)

type RoleController struct {
	roleService  *service.RoleService
	errorService *service.ErrorService
}

func NewRoleController(i *do.Injector) (*RoleController, error) {
	return &RoleController{
		roleService:  do.MustInvoke[*service.RoleService](i),
		errorService: do.MustInvoke[*service.ErrorService](i),
	}, nil
}

func (rc *RoleController) RegisterRoutes(e *echo.Echo) {
	e.GET(constants.RoleEndpoint, mmidleware.PermissionCheck([]mmidleware.AllowedPermissions{
		{
			Cat: entities.RoleManagement,
			Val: mmidleware.READ,
		},
	}, rc.HandleGetAllRoles)).Name = "getAllRoles"
	e.GET(constants.RoleEndpoint+"/:id", mmidleware.PermissionCheck([]mmidleware.AllowedPermissions{
		{
			Cat: entities.RoleManagement,
			Val: mmidleware.READ,
		},
	}, rc.HandleGetRoleById)).Name = "getRoleById"
	e.POST(constants.RoleEndpoint, mmidleware.PermissionCheck([]mmidleware.AllowedPermissions{
		{
			Cat: entities.RoleManagement,
			Val: mmidleware.WRITE,
		},
	}, rc.HandleCreateRole)).Name = "createRole"
	e.DELETE(constants.RoleEndpoint+"/:id", mmidleware.PermissionCheck([]mmidleware.AllowedPermissions{
		{
			Cat: entities.RoleManagement,
			Val: mmidleware.WRITE,
		},
	}, rc.HandleDeleteRole)).Name = "deleteRole"
}

func (rc *RoleController) HandleGetAllRoles(c echo.Context) error {
	query := new(dto.WithFilterRequestDTO)

	if err := c.Bind(query); err != nil {
		return rc.errorService.BadRequestError(err)
	}

	user := middleware.GetUserFromContext(c)
	res, err := rc.roleService.GetAllRoles(user, query.Filter)

	if err != nil {
		return rc.errorService.InternalServerError(err)
	}

	return c.JSON(200, res)
}

func (rc *RoleController) HandleGetRoleById(c echo.Context) error {
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

func (rc *RoleController) HandleCreateRole(c echo.Context) error {
	role := new(roledto.CreateRoleDTO)

	if err := c.Bind(role); err != nil {
		return echo.NewHTTPError(echo.ErrBadRequest.Code, err.Error())
	}

	if err := c.Validate(role); err != nil {
		return echo.NewHTTPError(echo.ErrBadRequest.Code, err.Error())
	}

	user := middleware.GetUserFromContext(c)

	newRole, err := rc.roleService.CreateRole(sroledto.CreateRoleDTO{
		Name:     role.Name,
		BranchId: user.BranchId.UUID,
	})

	if err != nil {
		return rc.errorService.InternalServerError(err)
	}

	return c.JSON(201, newRole)
}

func (rc *RoleController) HandleDeleteRole(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return rc.errorService.BadRequestError(echo.NewHTTPError(500, "Id is required."))
	}

	uid := uuid.MustParse(id)
	user := middleware.GetUserFromContext(c)

	err := rc.roleService.DeleteById(user, uid)

	if err != nil {
		return rc.errorService.InternalServerError(err)
	}

	return c.NoContent(200)
}
