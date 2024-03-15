package user

import (
	"fmt"
	"backend/base"
	constants "backend/base/contants"
	basedto "backend/base/dto"
	permissionentity "backend/permission/entity"
	userdto "backend/user/dto"
	userentity "backend/user/entity"

	"github.com/google/uuid"
	echo "github.com/labstack/echo/v4"
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

func (uc *Controller) RegisterRoutes(e *echo.Echo) {
	readPermissions := base.PermissionCheck([]base.AllowedPermissions{
		{
			Cat: permissionentity.UserManagement,
			Val: base.READ,
		},
	})
	writePermissions := base.PermissionCheck([]base.AllowedPermissions{
		{
			Cat: permissionentity.UserManagement,
			Val: base.WRITE,
		},
	})
	e.GET(constants.UserEndpoint, uc.HandleGetAllUsers, readPermissions).Name = "getAllUsers"
	e.GET(constants.UserEndpoint+"/me", uc.HandleGetMe).Name = "getMe"
	e.GET(constants.UserEndpoint+"/exist-by-email/:email", uc.HandleDoesEmailExists).Name = "existByEmail"
	e.DELETE(constants.UserEndpoint+"/:id", uc.HandleDeleteById, writePermissions).Name = "deleteUserById"
	e.POST(constants.UserEndpoint, uc.HandleCreateUser, writePermissions).Name = "createUser"
	e.GET(constants.UserEndpoint+"/:id", uc.HandleGetById, readPermissions).Name = "getUserById"
}

func (uc *Controller) HandleGetAllUsers(c echo.Context) error {
	query := new(basedto.PaginatedRequest)

	if err := c.Bind(query); err != nil {
		return uc.errorService.BadRequestError(err)
	}

	if err := c.Validate(query); err != nil {
		return uc.errorService.BadRequestError(err)
	}

	user := base.GetUserFromContext(c)
	res, err := uc.service.FindPagedWithRoleAndPermissions(user, query.Page, query.Limit, query.Filter)
	if err != nil {
		return uc.errorService.InternalServerError(err)
	}

	cnt, err := uc.service.CountPagedByUserBranchWithFilter(user, query.Filter)
	if err != nil {
		return uc.errorService.InternalServerError(err)
	}

	return c.JSON(200, basedto.Paginated[userdto.UserWithRoleAndPermissions]{
		Data: res,
		Meta: base.CalculatePaginationMeta(cnt, query.Page, query.Limit),
	})
}

func (uc *Controller) HandleGetById(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return uc.errorService.BadRequestError(echo.ErrBadRequest)
	}

	authUser := base.GetUserFromContext(c)
	uid := uuid.MustParse(id)
	res, err := uc.service.FindByUserBranchAndId(authUser, uid)
	if err != nil {
		return uc.errorService.InternalServerError(err)
	}
	return c.JSON(200, res)
}

func (uc *Controller) HandleGetMe(c echo.Context) error {
	user := base.GetUserFromContext(c)
	return c.JSON(200, user)
}

func (uc *Controller) HandleDoesEmailExists(c echo.Context) error {
	email := c.Param("email")

	if email == "" {
		return uc.errorService.BadRequestError(echo.ErrBadRequest)
	}

	fmt.Print("This is something that we are writing...")

	res, err := uc.service.FindByEmailWithPassword(email)
	if err != nil {
		return uc.errorService.InternalServerError(err)
	}
	return c.JSON(200, res.Email == email)
}

func (uc *Controller) HandleDeleteById(c echo.Context) error {
	params := new(userdto.DeleteUserDto)

	if err := c.Bind(params); err != nil {
		return uc.errorService.BadRequestError(err)
	}

	user := base.GetUserFromContext(c)
	err := uc.service.DeleteByUserBranchAndId(user, params.ID)
	if err != nil {
		return uc.errorService.BadRequestError(err)
	}

	return c.NoContent(200)
}

func (uc *Controller) HandleCreateUser(c echo.Context) error {
	newUser := new(userdto.CreateUserDTO)

	if err := c.Bind(newUser); err != nil {
		return uc.errorService.BadRequestError(err)
	}

	if err := c.Validate(newUser); err != nil {
		return uc.errorService.BadRequestError(err)
	}

	authUser := base.GetUserFromContext(c)
	created, err := uc.service.Create(userentity.User{
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Email:     newUser.Email,
		Password:  newUser.Password,
		BranchId:  authUser.BranchId,
		RoleId:    uuid.NullUUID{UUID: newUser.RoleId, Valid: true},
	})
	if err != nil {
		return uc.errorService.InternalServerError(err)
	}

	return c.JSON(200, created)
}
