package api

import (
	"app/api/dto"
	userdto "app/api/dto/user_dto"
	constants "app/contants"
	"app/entities"
	"app/middleware"
	service "app/service"
	suserdto "app/service/dto/user_dto"

	"github.com/google/uuid"
	echo "github.com/labstack/echo/v4"
	"github.com/samber/do"
)

type UserController struct {
	userService  *service.UserService
	errorService *service.ErrorService
}

func NewUserController(i *do.Injector) (*UserController, error) {
	return &UserController{
		userService:  do.MustInvoke[*service.UserService](i),
		errorService: do.MustInvoke[*service.ErrorService](i),
	}, nil
}

func (uc *UserController) RegisterRoutes(e *echo.Echo) {
	e.GET(constants.UserEndpoint, middleware.PermissionCheck(
		[]middleware.AllowedPermissions{
			{
				Cat: entities.UserManagement,
				Val: middleware.READ,
			},
		},
		uc.HandleGetAllUsers,
	)).Name = "getAllUsers"
	e.GET(constants.UserEndpoint+"/me", uc.HandleGetMe).Name = "getMe"
	e.GET(constants.UserEndpoint+"/exist-by-email/:email", uc.HandleDoesEmailExists).Name = "existByEmail"
	e.DELETE(constants.UserEndpoint+"/:id", middleware.PermissionCheck(
		[]middleware.AllowedPermissions{
			{
				Cat: entities.UserManagement,
				Val: middleware.WRITE,
			},
		},
		uc.HandleDeleteById,
	)).Name = "deleteUserById"
	e.POST(constants.UserEndpoint, middleware.PermissionCheck([]middleware.AllowedPermissions{
		{
			Cat: entities.UserManagement,
			Val: middleware.WRITE,
		},
	}, uc.HandleCreateUser)).Name = "createUser"
	e.GET(constants.UserEndpoint+"/:id", middleware.PermissionCheck([]middleware.AllowedPermissions{
		{
			Cat: entities.UserManagement,
			Val: middleware.READ,
		},
	}, uc.HandleGetById)).Name = "getUserById"
}

func (uc *UserController) HandleGetAllUsers(c echo.Context) error {
	query := new(dto.PaginatedRequest)

	if err := c.Bind(query); err != nil {
		return uc.errorService.BadRequestError(err)
	}

	if err := c.Validate(query); err != nil {
		return uc.errorService.BadRequestError(err)
	}

	user := middleware.GetUserFromContext(c)
	res, err := uc.userService.GetPaginated(user, query.Page, query.Limit, query.Filter)

	if err != nil {
		return uc.errorService.InternalServerError(err)
	}

	return c.JSON(200, res)
}

func (uc *UserController) HandleGetById(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return uc.errorService.BadRequestError(echo.ErrBadRequest)
	}

	authUser := middleware.GetUserFromContext(c)
	uid := uuid.MustParse(id)
	res, err := uc.userService.GetById(authUser, uid)
	if err != nil {
		return uc.errorService.InternalServerError(err)
	}
	return c.JSON(200, res)
}

func (uc *UserController) HandleGetMe(c echo.Context) error {
	user := middleware.GetUserFromContext(c)
	return c.JSON(200, user)
}

func (uc *UserController) HandleDoesEmailExists(c echo.Context) error {
	email := c.Param("email")

	if email == "" {
		return uc.errorService.BadRequestError(echo.ErrBadRequest)
	}

	res, err := uc.userService.GetByEmailWithPassword(email)
	if err != nil {
		return uc.errorService.InternalServerError(err)
	}
	return c.JSON(200, res.Email == email)
}

func (uc *UserController) HandleDeleteById(c echo.Context) error {
	params := new(userdto.DeleteUserDto)

	if err := c.Bind(params); err != nil {
		return uc.errorService.BadRequestError(err)
	}

	user := middleware.GetUserFromContext(c)
	return uc.userService.DeleteById(user, params.ID)
}

func (uc *UserController) HandleCreateUser(c echo.Context) error {
	newUser := new(userdto.CreateUserDTO)

	if err := c.Bind(newUser); err != nil {
		return uc.errorService.BadRequestError(err)
	}

	if err := c.Validate(newUser); err != nil {
		return uc.errorService.BadRequestError(err)
	}

	authUser := middleware.GetUserFromContext(c)
	created, err := uc.userService.Create(authUser, suserdto.CreateUserDTO{
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Email:     newUser.Email,
		RoleId:    newUser.RoleId,
	})

	if err != nil {
		return uc.errorService.InternalServerError(err)
	}

	return c.JSON(200, created)
}
