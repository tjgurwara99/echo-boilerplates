package views

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tjgurwara99/echo-boilerplates/basic-mvt/db"
	"github.com/tjgurwara99/echo-boilerplates/basic-mvt/models"
	"github.com/tjgurwara99/echo-boilerplates/basic-mvt/templates"
	"github.com/tjgurwara99/echo-boilerplates/basic-mvt/templates/auth"
	"golang.org/x/crypto/bcrypt"
)

func LoginPage(c echo.Context) error {
	csrf, ok := c.Get(middleware.DefaultCSRFConfig.ContextKey).(string)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "csrf error")
	}
	t := auth.LoginView(csrf)
	base := templates.Base(t)
	return base.Render(c.Request().Context(), c.Response())
}

func Login(c echo.Context) error {
	var user models.User
	email := c.FormValue("email")
	password := c.FormValue("password")
	if email == "" {
		return c.JSON(http.StatusBadRequest, models.Error{
			Message: "Email is required",
		})
	}
	if password == "" {
		return c.JSON(http.StatusBadRequest, models.Error{
			Message: "Password is required",
		})
	}
	err := db.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return err
	}
	session, err := session.Get("session", c)
	if err != nil {
		return err
	}
	session.Values["current_user_id"] = user.ID.String()
	err = session.Save(c.Request(), c.Response())
	if err != nil {
		return err
	}
	c.Response().Header().Add("HX-Redirect", "/")
	return c.HTML(http.StatusTemporaryRedirect, "")
}

func Logout(c echo.Context) error {
	session, err := session.Get("session", c)
	if err != nil {
		return err
	}
	session.Values["current_user_id"] = nil
	err = session.Save(c.Request(), c.Response())
	if err != nil {
		return err
	}
	c.Response().Header().Add("HX-Redirect", "/")
	return c.HTML(http.StatusTemporaryRedirect, "")
}

func SignupPage(c echo.Context) error {
	csrf, ok := c.Get(middleware.DefaultCSRFConfig.ContextKey).(string)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "csrf error")
	}
	t := auth.SignUpView(csrf)
	base := templates.Base(t)
	return base.Render(c.Request().Context(), c.Response())
}

func Signup(c echo.Context) error {
	var user models.User
	email := c.FormValue("email")
	password := c.FormValue("password")
	passwordConfirmation := c.FormValue("password_confirmation")
	// TODO: Instead of returning a JSON error, render the signup page with the errors
	if email == "" {
		return c.JSON(http.StatusBadRequest, models.Error{
			Message: "Email is required",
		})
	}
	if password == "" {
		return c.JSON(http.StatusBadRequest, models.Error{
			Message: "Password is required",
		})
	}
	if passwordConfirmation == "" {
		return c.JSON(http.StatusBadRequest, models.Error{
			Message: "Password confirmation is required",
		})
	}
	if password != passwordConfirmation {
		return c.JSON(http.StatusBadRequest, models.Error{
			Message: "Password and password confirmation do not match",
		})
	}
	err := db.DB.Where("email = ?", email).First(&user).Error
	if err == nil {
		return c.JSON(http.StatusBadRequest, models.Error{
			Message: "User already exists",
		})
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user = models.User{
		Email:        email,
		PasswordHash: string(hashedPassword),
	}
	err = db.DB.Create(&user).Error
	if err != nil {
		return err
	}
	session, err := session.Get("session", c)
	if err != nil {
		return err
	}
	session.Values["current_user_id"] = user.ID.String()
	err = session.Save(c.Request(), c.Response())
	if err != nil {
		return err
	}
	c.Response().Header().Add("HX-Redirect", "/")
	return c.HTML(http.StatusTemporaryRedirect, "")
}
