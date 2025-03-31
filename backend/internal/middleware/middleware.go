package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/veetipihlava/shakki-peli/internal/database"
)

var UserIDContextName = "user-id"
var DatabaseContextName = "database-service"
var RedisContextName = "redis"

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
		// a tmp implementation
		db := c.Get(DatabaseContextName).(*database.DatabaseService)
		user, err := db.CreateUser("shared-user")
		if err != nil {
			return err
		}

		c.Set(UserIDContextName, user)
		return next(c)
	}
}
