package api

import (
	"app/api/dto"
	constants "app/contants"
	"app/entities"
	"app/middleware"
	"app/service"

	"github.com/labstack/echo/v4"
	"github.com/samber/do"
)

type CustomerController struct {
	customerService *service.CustomerService
	errorService    *service.ErrorService
}

func NewCustomerController(i *do.Injector) (*CustomerController, error) {
	return &CustomerController{
		customerService: do.MustInvoke[*service.CustomerService](i),
		errorService:    do.MustInvoke[*service.ErrorService](i),
	}, nil
}

func (cc *CustomerController) RegisterRoutes(e *echo.Echo) {
	readPermissions := []middleware.AllowedPermissions{
		{
			Cat: entities.CustomerManagement,
			Val: middleware.READ,
		},
	}
	e.GET(constants.CustomersEndpoint, cc.HandleGetPagedByUser, middleware.PermissionCheck(readPermissions))
}

func (cc *CustomerController) HandleGetPagedByUser(c echo.Context) error {
	p := new(dto.PaginatedRequest)

	if err := c.Bind(p); err != nil {
		return cc.errorService.BadRequestError(err)
	}

	if err := c.Validate(p); err != nil {
		return cc.errorService.BadRequestError(err)
	}

	u := middleware.GetUserFromContext(c)

	cs, err := cc.customerService.GetPageByUser(u, p.Page, p.Limit, p.Filter)
	if err != nil {
		return cc.errorService.InternalServerError(err)
	}

	return c.JSON(200, cs)
}
