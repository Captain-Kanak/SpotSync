package middleware

import (
	"net/http"
	"spot-sync/internal/auth"
	"spot-sync/internal/httpresponse"

	"github.com/labstack/echo/v5"
)

const (
	ContextUserKey = "user"
)

func AuthMiddleware(jwtService auth.JWTService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			cookie, err := c.Cookie("access_token")

			if err != nil {
				return c.JSON(http.StatusUnauthorized, httpresponse.Response{
					Success: false,
					Message: "Missing or invalid token",
				})
			}

			claims, err := jwtService.ValidateToken(cookie.Value)

			if err != nil {
				return c.JSON(http.StatusUnauthorized, httpresponse.Response{
					Success: false,
					Message: "Invalid or expired token",
				})
			}

			c.Set(ContextUserKey, claims)

			return next(c)
		}
	}
}
