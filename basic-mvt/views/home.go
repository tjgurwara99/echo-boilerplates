package views

import (
	"github.com/labstack/echo/v4"
	"github.com/tjgurwara99/echo-boilerplates/models"
	"github.com/tjgurwara99/echo-boilerplates/templates"
)

func HomePage(c echo.Context) error {
	user, ok := c.Get("current_user").(models.User)
	if !ok {
		t := templates.SimplePage(nil)
		base := templates.Base(t)
		return base.Render(c.Request().Context(), c.Response())
	}
	t := templates.SimplePage(&user)
	base := templates.Base(t)
	return base.Render(c.Request().Context(), c.Response())
}
