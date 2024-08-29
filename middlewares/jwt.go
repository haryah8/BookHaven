package middlewares

import (
	"BookHaven/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// JWT returns a middleware function for JWT authorization
func JWT(logger *logrus.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"status":  "Error",
					"message": "Missing authorization header",
				})
			}

			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenStr == authHeader {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"status":  "Error",
					"message": "Invalid authorization header format",
				})
			}

			claims, err := utils.ValidateToken(tokenStr)
			if err != nil {
				logger.Error("Invalid token: ", err)
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"status":  "Error",
					"message": "Invalid token",
				})
			}
			fmt.Println(claims.UserId)
			// Attach the claims to the context
			c.Set("user", claims)
			return next(c)
		}
	}
}
