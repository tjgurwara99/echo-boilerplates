package main

import (
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tjgurwara99/echo-boilerplates/db"
	"github.com/tjgurwara99/echo-boilerplates/middlewares"
	"github.com/tjgurwara99/echo-boilerplates/views"
)

func setRoutes(e *echo.Echo) *echo.Echo {
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})
	fs := echo.MustSubFS(static, "static")
	e.StaticFS("/static", fs)

	e.GET("/login", views.LoginPage)
	e.POST("/login", views.Login)
	e.GET("/logout", views.Logout)
	e.GET("/signup", views.SignupPage)
	e.POST("/signup", views.Signup)
	e.GET("/", views.HomePage)
	return e
}

func setMiddlewares(e *echo.Echo) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(session.MiddlewareWithConfig(session.Config{
		Store: db.SessionStore,
	}), middlewares.SetCurrentUser)
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup:    "form:_csrf,header:X-CSRF-Token",
		CookieMaxAge:   3600,
		CookieSecure:   true,
		CookieHTTPOnly: true,
	}))
	return e
}
