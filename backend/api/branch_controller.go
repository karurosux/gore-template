package api

import (
	constants "app/contants"
	"app/entities"
	mmidleware "app/middleware"
	"app/service"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/samber/do"
)

type BranchController struct {
	branchService *service.BranchService
	errorService  *service.ErrorService
}

func NewBranchController(i *do.Injector) (*BranchController, error) {
	return &BranchController{
		branchService: do.MustInvoke[*service.BranchService](i),
		errorService:  do.MustInvoke[*service.ErrorService](i),
	}, nil
}

func (bc *BranchController) RegisterRoutes(e *echo.Echo) {
	e.GET(constants.BranchEndpoint, mmidleware.PermissionCheck([]mmidleware.AllowedPermissions{
		{
			Cat: entities.BranchManagement,
			Val: mmidleware.READ,
		},
	}, bc.HandleGetAllBranches)).Name = "HandleGetAllBranches"
	e.GET(constants.BranchEndpoint+"/:id", mmidleware.PermissionCheck([]mmidleware.AllowedPermissions{
		{
			Cat: entities.BranchManagement,
			Val: mmidleware.READ,
		},
	}, bc.HandleGetBranchById)).Name = "HandleGetBranchById"
}

func (bc *BranchController) HandleGetAllBranches(c echo.Context) error {
	return c.JSON(200, bc.branchService.GetAllBranches())
}

func (bc *BranchController) HandleGetBranchById(c echo.Context) error {
	id := c.Param("id")
	pid, pidErr := uuid.Parse(id)

	if pidErr != nil {
		return echo.NewHTTPError(echo.ErrBadRequest.Code, pidErr.Error())
	}

	branch, err := bc.branchService.GetBranchId(pid)

	if err != nil {
		return bc.errorService.NotFoundError(err)
	}

	return c.JSON(200, branch)
}
