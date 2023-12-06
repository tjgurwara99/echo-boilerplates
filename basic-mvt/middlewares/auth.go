package middlewares

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/tjgurwara99/echo-boilerplates/db"
	"github.com/tjgurwara99/echo-boilerplates/models"
)

// SetCurrentUser sets the current user in the context if
// there exists current_user_id in the session.
func SetCurrentUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get("session", c)
		if err != nil {
			return err
		}
		userID := sess.Values["current_user_id"]
		if userID != nil {
			var user models.User
			err := db.DB.Where("id = ?", userID).First(&user).Error
			if err != nil {
				sess.Values["current_user_id"] = nil
				err = sess.Save(c.Request(), c.Response())
				if err != nil {
					log.Error(err)
				}
				return c.Redirect(http.StatusTemporaryRedirect, c.Request().URL.String())
			}
			c.Set("current_user", user)
		}
		return next(c)
	}
}

func Authorize(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("current_user")
		if user == nil {
			return c.Redirect(http.StatusTemporaryRedirect, "/login")
		}
		return next(c)
	}
}

func AdminOnly(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// user := c.Get("current_user").(models.User)
		// TODO: Check if user is admin
		return next(c)
	}
}
