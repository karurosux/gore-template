package api

import (
	permissionsdto "app/api/dto/permissions_dto"
	constants "app/contants"
	"app/entities"
	"app/middleware"
	"app/service"
	spermissionsdto "app/service/dto/permissions_dto"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/samber/do"
)

type PermissionsController struct {
	permissionsService *service.PermissionsService
	errorService       *service.ErrorService
}

func NewPermissionsController(i *do.Injector) (*PermissionsController, error) {
	return &PermissionsController{
		permissionsService: do.MustInvoke[*service.PermissionsService](i),
		errorService:       do.MustInvoke[*service.ErrorService](i),
	}, nil
}

func (pc *PermissionsController) RegisterRoutes(e *echo.Echo) {
	readPermissions := []middleware.AllowedPermissions{
		{
			Cat: entities.RoleManagement,
			Val: middleware.READ,
		},
	}
	writePermissions := []middleware.AllowedPermissions{
		{
			Cat: entities.RoleManagement,
			Val: middleware.WRITE,
		},
	}
	e.GET(constants.PermissionsEndpoint, pc.HandleGetPermissionsByRole, middleware.PermissionCheck(readPermissions)).Name = "getPermissionsByRole"
	e.GET(constants.PermissionsEndpoint+"/categories", pc.HandleGetPermissionCategories, middleware.PermissionCheck(readPermissions)).Name = "getPermissionsCategories"
	e.POST(constants.PermissionsEndpoint, pc.HandleCreatePermission, middleware.PermissionCheck(writePermissions)).Name = "createPermission"
	e.PATCH(constants.PermissionsEndpoint+"/:id", pc.HandlePatchPermission, middleware.PermissionCheck(writePermissions)).Name = "updatePermission"
	e.DELETE(constants.PermissionsEndpoint+"/:id", pc.HandleDeleteById, middleware.PermissionCheck(writePermissions)).Name = "deleteById"
}

func (pc *PermissionsController) HandleGetPermissionsByRole(c echo.Context) error {
	query := new(permissionsdto.GetByRoleIdDTO)

	if err := c.Bind(query); err != nil {
		return pc.errorService.BadRequestError(err)
	}

	if err := c.Validate(query); err != nil {
		return pc.errorService.BadRequestError(err)
	}

	permissions, err := pc.permissionsService.GetPermissionsByRoleId(uuid.NullUUID{
		UUID:  query.RoleId,
		Valid: true,
	})
	if err != nil {
		return pc.errorService.InternalServerError(err)
	}

	return c.JSON(200, permissions)
}

func (pc *PermissionsController) HandleGetPermissionCategories(c echo.Context) error {
	return c.JSON(200, []entities.PermissionCategoryVal{
		entities.BranchManagement,
		entities.CustomerManagement,
		entities.RoleManagement,
		entities.UserManagement,
		entities.SurveyManagement,
	})
}

func (pc *PermissionsController) HandleCreatePermission(c echo.Context) error {
	body := new(permissionsdto.CreatePermissionDTO)

	if err := c.Bind(body); err != nil {
		return pc.errorService.BadRequestError(err)
	}

	if err := c.Validate(body); err != nil {
		return pc.errorService.BadRequestError(err)
	}

	permission, err := pc.permissionsService.Create(spermissionsdto.CreatePermissionDTO{
		Category: body.Category,
		Write:    body.Write,
		Read:     body.Read,
		RoleId:   body.RoleId,
	})
	if err != nil {
		return pc.errorService.InternalServerError(err)
	}

	return c.JSON(200, permission)
}

func (pc *PermissionsController) HandleDeleteById(c echo.Context) error {
	path := new(permissionsdto.PermissionByIdDTO)

	if err := c.Bind(path); err != nil {
		return pc.errorService.BadRequestError(err)
	}

	if err := c.Validate(path); err != nil {
		return pc.errorService.BadRequestError(err)
	}

	if err := pc.permissionsService.DeleteById(path.ID); err != nil {
		return pc.errorService.InternalServerError(err)
	}

	return c.NoContent(200)
}

func (pc *PermissionsController) HandlePatchPermission(c echo.Context) error {
	req := new(permissionsdto.PatchPermissionDTO)

	if err := c.Bind(req); err != nil {
		return pc.errorService.BadRequestError(err)
	}

	if err := c.Validate(req); err != nil {
		return pc.errorService.BadRequestError(err)
	}

	err := pc.permissionsService.PatchById(req.ID, req.Write, req.Read)
	if err != nil {
		return pc.errorService.InternalServerError(err)
	}

	return c.NoContent(200)
}
