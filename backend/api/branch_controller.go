package api

import (
	"app/api/dto"
	abranchdto "app/api/dto/branch_dto"
	constants "app/contants"
	"app/entities"
	mmidleware "app/middleware"
	"app/model"
	"app/service"
	branchdto "app/service/dto/branch_dto"
	"app/utils"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/samber/do"
)

type BranchController struct {
	service      *service.BranchService
	errorService *service.ErrorService
}

func NewBranchController(i *do.Injector) (*BranchController, error) {
	return &BranchController{
		service:      do.MustInvoke[*service.BranchService](i),
		errorService: do.MustInvoke[*service.ErrorService](i),
	}, nil
}

func (bc *BranchController) RegisterRoutes(e *echo.Echo) {
	readPermissions := []mmidleware.AllowedPermissions{
		{
			Cat: entities.BranchManagement,
			Val: mmidleware.READ,
		},
	}
	writePermissions := []mmidleware.AllowedPermissions{
		{
			Cat: entities.BranchManagement,
			Val: mmidleware.WRITE,
		},
	}
	e.GET(constants.BranchEndpoint, bc.HandleGetPaginatedBranches, mmidleware.SuperAdminOnly(), mmidleware.PermissionCheck(readPermissions)).Name = "HandleGetAllBranches"
	e.GET(constants.BranchEndpoint+"/:id", bc.HandleGetBranchById, mmidleware.SuperAdminOnly(), mmidleware.PermissionCheck(readPermissions)).Name = "HandleGetBranchById"
	e.DELETE(constants.BranchEndpoint+"/:id", bc.HandleDeleteBranch, mmidleware.SuperAdminOnly(), mmidleware.PermissionCheck(writePermissions)).Name = "deleteBranch"
	e.POST(constants.BranchEndpoint, bc.HandleCreateBranch, mmidleware.SuperAdminOnly(), mmidleware.PermissionCheck(writePermissions)).Name = "createBranch"
}

func (bc *BranchController) HandleGetAllBranches(c echo.Context) error {
	r, err := bc.service.FindAll()
	if err != nil {
		return bc.errorService.InternalServerError(err)
	}
	return c.JSON(200, r)
}

func (bc *BranchController) HandleGetBranchById(c echo.Context) error {
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

func (bc *BranchController) HandleGetPaginatedBranches(c echo.Context) error {
	q := new(dto.PaginatedRequest)
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

	return c.JSON(200, model.Paginated[branchdto.BranchDTO]{
		Data: r,
		Meta: utils.CalculatePaginationMeta(cnt, q.Page, q.Limit),
	})
}

func (bc *BranchController) HandleDeleteBranch(c echo.Context) error {
	params := new(abranchdto.DeleteBranchDto)

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

func (bc *BranchController) HandleCreateBranch(c echo.Context) error {
	b := new(abranchdto.CreateBranchDTO)

	if err := c.Bind(b); err != nil {
		return bc.errorService.BadRequestError(err)
	}

	if err := c.Validate(b); err != nil {
		return bc.errorService.BadRequestError(err)
	}

	nb, err := bc.service.Create(entities.Branch{
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
