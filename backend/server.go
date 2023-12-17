package main

import (
	"app/api"
	constants "app/contants"
	mmidleware "app/middleware"
	"app/model"
	"app/repository"
	"app/service"
	"app/utils"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	do "github.com/samber/do"
	"gorm.io/gorm"
)

type (
	Server struct {
		app       *echo.Echo
		container *do.Injector
	}

	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func (s *Server) Start() {
	s.configureORM()
	s.provideRepositories()
	s.provideServices()
	s.provideControllers()
	s.configMiddleware()
	s.registerControllers()
	s.setStaticRoutes()
	utils.Logger.LogInfo().Msg(s.app.Start(fmt.Sprintf(":%d", getPort())).Error())
}

func (s *Server) provideRepositories() {
	do.Provide(s.container, repository.NewPermissionsRepository)
	do.Provide(s.container, repository.NewUserRepository)
	do.Provide(s.container, repository.NewBranchRepository)
	do.Provide(s.container, repository.NewRoleRepository)
}

func (s *Server) provideServices() {
	do.Provide(s.container, mmidleware.NewSessionCheck)
	do.Provide(s.container, service.NewPermissionsService)
	do.Provide(s.container, service.NewUserService)
	do.Provide(s.container, service.NewRoleService)
	do.Provide(s.container, service.NewErrorService)
	do.Provide(s.container, service.NewPasswordService)
	do.Provide(s.container, service.NewAuthService)
	do.Provide(s.container, service.NewJwtService)
	do.Provide(s.container, service.NewBranchService)
}

func (s *Server) provideControllers() {
	do.Provide(s.container, api.NewUserController)
	do.Provide(s.container, api.NewBranchController)
	do.Provide(s.container, api.NewRoleController)
	do.Provide(s.container, api.NewAuthController)
	do.Provide(s.container, api.NewPermissionsController)
}

func (s *Server) registerControllers() {
	var controllers *[]model.Controller = &[]model.Controller{
		do.MustInvoke[*api.UserController](s.container),
		do.MustInvoke[*api.BranchController](s.container),
		do.MustInvoke[*api.RoleController](s.container),
		do.MustInvoke[*api.AuthController](s.container),
		do.MustInvoke[*api.PermissionsController](s.container),
	}

	for i := 0; i < len(*controllers); i++ {
		(*controllers)[i].RegisterRoutes(s.app)
	}
}

func (s *Server) configureORM() {
	db := utils.MigrateDb()

	do.Provide(s.container, func(i *do.Injector) (*gorm.DB, error) {
		return db, nil
	})
}

func (s *Server) configMiddleware() {
	sessionCheck := do.MustInvoke[*mmidleware.SessionCheck](s.container)
	s.app.Validator = &CustomValidator{validator: validator.New()}
	utils.NewLogger()
	s.app.Use(mmidleware.LoggingMiddleware)
	s.app.Use(mmidleware.AppendAppContextMiddleware)
	s.app.Use(middleware.Recover())
	s.app.Use(session.Middleware(sessions.NewCookieStore([]byte(os.Getenv("JWT_SECRET")))))
	s.app.Use(sessionCheck.SessionCheckMiddleware)
}

func (s *Server) setStaticRoutes() {
	s.app.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Skipper: nil,
		// Root directory from where the static content is served.
		Root: "static",
		// Index file for serving a directory.
		// Optional. Default value "index.html".
		Index: "index.html",
		// Enable HTML5 mode by forwarding all not-found requests to root so that
		// SPA (single-page application) can handle the routing.
		HTML5:      true,
		Browse:     false,
		IgnoreBase: false,
		Filesystem: nil,
	}))
}

func getPort() int {
	envPort := os.Getenv("PORT")

	if envPort != "" {
		port, _ := strconv.Atoi(envPort)
		return port
	}

	return constants.Port
}
