package permission

import (
	"backend/base"
	constants "backend/base/contants"
	"backend/permission/dto"
	"backend/permission/entity"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/samber/do"
)

type Controller struct {
	permissionsService *Service
	errorService       *base.ErrorService
}

func NewController(i *do.Injector) (*Controller, error) {
	return &Controller{
		permissionsService: do.MustInvoke[*Service](i),
		errorService:       do.MustInvoke[*base.ErrorService](i),
	}, nil
}

func (pc *Controller) RegisterRoutes(e *echo.Echo) {
	readPermissions := base.PermissionCheck([]base.AllowedPermissions{
		{
			Cat: entity.RoleManagement,
			Val: base.READ,
		},
	})
	writePermissions := base.PermissionCheck([]base.AllowedPermissions{
		{
			Cat: entity.RoleManagement,
			Val: base.WRITE,
		},
	})
	e.GET(constants.PermissionsEndpoint, pc.HandleGetPermissionsByRole, readPermissions).Name = "getPermissionsByRole"
	e.GET(constants.PermissionsEndpoint+"/categories", pc.HandleGetPermissionCategories, readPermissions).Name = "getPermissionsCategories"
	e.POST(constants.PermissionsEndpoint, pc.HandleCreatePermission, writePermissions).Name = "createPermission"
	e.PATCH(constants.PermissionsEndpoint+"/:id", pc.HandlePatchPermission, writePermissions).Name = "updatePermission"
	e.DELETE(constants.PermissionsEndpoint+"/:id", pc.HandleDeleteById, writePermissions).Name = "deleteById"
}

func (pc *Controller) HandleGetPermissionsByRole(c echo.Context) error {
	query := new(dto.GetByRoleIdDTO)

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

func (pc *Controller) HandleGetPermissionCategories(c echo.Context) error {
	return c.JSON(200, []entity.PermissionCategoryVal{
		entity.BranchManagement,
		entity.CustomerManagement,
		entity.RoleManagement,
		entity.UserManagement,
	})
}

func (pc *Controller) HandleCreatePermission(c echo.Context) error {
	body := new(dto.CreatePermissionDTO)

	if err := c.Bind(body); err != nil {
		return pc.errorService.BadRequestError(err)
	}

	if err := c.Validate(body); err != nil {
		return pc.errorService.BadRequestError(err)
	}

	permission, err := pc.permissionsService.Create(dto.SCreatePermissionDTO{
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

func (pc *Controller) HandleDeleteById(c echo.Context) error {
	path := new(dto.PermissionByIdDTO)

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

func (pc *Controller) HandlePatchPermission(c echo.Context) error {
	req := new(dto.PatchPermissionDTO)

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
