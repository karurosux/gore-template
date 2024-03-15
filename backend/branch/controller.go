package branch

import (
	"backend/base"
	constants "backend/base/contants"
	basedto "backend/base/dto"
	"backend/branch/dto"
	"backend/branch/entity"
	permissionentity "backend/permission/entity"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/samber/do"
)

type Controller struct {
	service      *Service
	errorService *base.ErrorService
}

func NewController(i *do.Injector) (*Controller, error) {
	return &Controller{
		service:      do.MustInvoke[*Service](i),
		errorService: do.MustInvoke[*base.ErrorService](i),
	}, nil
}

func (bc *Controller) RegisterRoutes(e *echo.Echo) {
	readPermissions := base.PermissionCheck([]base.AllowedPermissions{
		{
			Cat: permissionentity.BranchManagement,
			Val: base.READ,
		},
	})
	writePermissions := base.PermissionCheck([]base.AllowedPermissions{
		{
			Cat: permissionentity.BranchManagement,
			Val: base.WRITE,
		},
	})
	e.GET(constants.BranchEndpoint, bc.HandleGetPaginatedBranches, base.SuperAdminOnly(), readPermissions).Name = "HandleGetAllBranches"
	e.GET(constants.BranchEndpoint+"/:id", bc.HandleGetBranchById, base.SuperAdminOnly(), readPermissions).Name = "HandleGetBranchById"
	e.DELETE(constants.BranchEndpoint+"/:id", bc.HandleDeleteBranch, base.SuperAdminOnly(), writePermissions).Name = "deleteBranch"
	e.POST(constants.BranchEndpoint, bc.HandleCreateBranch, base.SuperAdminOnly(), writePermissions).Name = "createBranch"
}

func (bc *Controller) HandleGetAllBranches(c echo.Context) error {
	r, err := bc.service.FindAll()
	if err != nil {
		return bc.errorService.InternalServerError(err)
	}
	return c.JSON(200, r)
}

func (bc *Controller) HandleGetBranchById(c echo.Context) error {
	id := c.Param("id")
	pid, pidErr := uuid.Parse(id)

	if pidErr != nil {
		return echo.NewHTTPError(echo.ErrBadRequest.Code, pidErr.Error())
	}

	branch, err := bc.service.FindById(pid)
	if err != nil {
		return bc.errorService.NotFoundError(err)
	}

	return c.JSON(200, branch)
}

func (bc *Controller) HandleGetPaginatedBranches(c echo.Context) error {
	q := new(basedto.PaginatedRequest)
	if err := c.Bind(q); err != nil {
		return bc.errorService.BadRequestError(err)
	}

	if err := c.Validate(q); err != nil {
		return bc.errorService.BadRequestError(err)
	}

	r, err := bc.service.FindPagedWithFilter(q.Page, q.Limit, q.Filter)
	if err != nil {
		return bc.errorService.InternalServerError(err)
	}

	cnt, err := bc.service.CountPagedWithFilter(q.Filter)
	if err != nil {
		return bc.errorService.BadRequestError(err)
	}

	return c.JSON(200, basedto.Paginated[dto.BranchDTO]{
		Data: r,
		Meta: base.CalculatePaginationMeta(cnt, q.Page, q.Limit),
	})
}

func (bc *Controller) HandleDeleteBranch(c echo.Context) error {
	params := new(dto.DeleteBranchDto)

	if err := c.Bind(params); err != nil {
		return bc.errorService.BadRequestError(err)
	}

	if err := c.Validate(params); err != nil {
		return bc.errorService.BadRequestError(err)
	}

	err := bc.service.DeleteById(params.ID)
	if err != nil {
		return bc.errorService.InternalServerError(err)
	}

	return c.NoContent(200)
}

func (bc *Controller) HandleCreateBranch(c echo.Context) error {
	b := new(dto.CreateBranchDTO)

	if err := c.Bind(b); err != nil {
		return bc.errorService.BadRequestError(err)
	}

	if err := c.Validate(b); err != nil {
		return bc.errorService.BadRequestError(err)
	}

	nb, err := bc.service.Create(entity.Branch{
		Name:    b.Name,
		City:    b.City,
		State:   b.State,
		Country: b.Country,
		ZipCode: b.ZipCode,
	})
	if err != nil {
		return bc.errorService.InternalServerError(err)
	}

	return c.JSON(200, nb)
}
