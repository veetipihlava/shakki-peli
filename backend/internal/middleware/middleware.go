package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/veetipihlava/shakki-peli/internal/database"
)

const UserContextName string = "user"
const DatabaseContextName string = "database-service"
const RedisContextName string = "redis"
const userIDCookieName string = "user-id-cookie"

// Middleware to pass a variable to context.
func WithContext(variableName string, variable interface{}) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(variableName, variable)
			return next(c)
		}
	}
}

// Middleware to provide handlers with the database user from the json web token.
func UseUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		db := c.Get(DatabaseContextName).(*database.DatabaseService)

		// tmp cookie implementation
		cookie, err := c.Cookie(userIDCookieName)
		if err != nil {
			if err == http.ErrNoCookie {
				user, err := db.CreateUser("non-authenticated-user")
				cookie := &http.Cookie{
					Name:     userIDCookieName,
					Value:    strconv.FormatInt(user.ID, 10),
					MaxAge:   int(time.Hour * 24 * 365),
					Path:     "/",
					Secure:   true,
					HttpOnly: true,
				}

				if err != nil {
					return err
				}

				c.SetCookie(cookie)
			}

			return err
		}

		userID, err := strconv.ParseInt(cookie.Value, 10, 64)
		if err != nil {
			return err
		}

		user, err := db.ReadUser(userID)
		if err != nil {
			return err
		}

		c.Set(UserContextName, user)
		return next(c)
	}
}
