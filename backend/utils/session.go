package utils

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func getSession(c echo.Context) *sessions.Session {
	sess, _ := session.Get("session", c)
	return sess
}

func GetSessionToken(c echo.Context) string {
	sess := getSession(c)

	if sess.Values["token"] == nil {
		return ""
	}

	return sess.Values["token"].(string)
}

func SetSessionToken(c echo.Context, token string) {
	sess := getSession(c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["token"] = token
	sess.Save(c.Request(), c.Response())
}

func ClearSessionToken(c echo.Context) {
	sess := getSession(c)
	sess.Values["token"] = nil
	sess.Save(c.Request(), c.Response())
}
