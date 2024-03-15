package main

import (
	"fmt"
	"backend/base"
	constants "backend/base/contants"
	mmidleware "backend/base/middleware"
	utils "backend/base/utils"
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
	RegisterRepositories(s.container)
	do.Provide(s.container, mmidleware.NewSessionCheck)
	RegisterServices(s.container)
	RegisterControllers(s.container)
	s.configMiddleware()
	RegisterRoutes(s.app, s.container)
	s.setStaticRoutes()
	utils.Logger.LogInfo().Msg(s.app.Start(fmt.Sprintf(":%d", getPort())).Error())
}

func (s *Server) configureORM() {
	db := MigrateDb()

	do.Provide(s.container, func(i *do.Injector) (*gorm.DB, error) {
		return db, nil
	})
}

func (s *Server) configMiddleware() {
	sessionCheck := do.MustInvoke[*mmidleware.SessionCheck](s.container)
	s.app.Validator = &CustomValidator{validator: validator.New()}
	utils.NewLogger()
	s.app.Use(mmidleware.LoggingMiddleware)
	s.app.Use(base.AppendBackendContextMiddleware)
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
