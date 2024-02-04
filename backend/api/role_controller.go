package api

import (
	"app/api/dto"
	roledto "app/api/dto/role_dto"
	constants "app/contants"
	"app/entities"
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
	readPermissions := []mmidleware.AllowedPermissions{
		{
			Cat: entities.RoleManagement,
			Val: mmidleware.READ,
		},
	}
	writePermissions := []mmidleware.AllowedPermissions{
		{
			Cat: entities.RoleManagement,
			Val: mmidleware.WRITE,
		},
	}
	e.GET(constants.RoleEndpoint, rc.HandleGetAllRoles, mmidleware.PermissionCheck(readPermissions)).Name = "getAllRoles"
	e.GET(constants.RoleEndpoint+"/:id", rc.HandleGetRoleById, mmidleware.PermissionCheck(readPermissions)).Name = "getRoleById"
	e.POST(constants.RoleEndpoint, rc.HandleCreateRole, mmidleware.PermissionCheck(writePermissions)).Name = "createRole"
	e.DELETE(constants.RoleEndpoint+"/:id", rc.HandleDeleteRole, mmidleware.PermissionCheck(writePermissions)).Name = "deleteRole"
}

func (rc *RoleController) HandleGetAllRoles(c echo.Context) error {
	query := new(dto.WithFilterRequestDTO)

	if err := c.Bind(query); err != nil {
		return rc.errorService.BadRequestError(err)
	}

	user := mmidleware.GetUserFromContext(c)
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

	user := mmidleware.GetUserFromContext(c)

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
	user := mmidleware.GetUserFromContext(c)

	err := rc.roleService.DeleteById(user, uid)
	if err != nil {
		return rc.errorService.InternalServerError(err)
	}

	return c.NoContent(200)
}
