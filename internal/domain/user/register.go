package user

import (
	"spot-sync/internal/httpresponse"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func Routes(db *gorm.DB, api *echo.Group) {

	api.POST("/auth/register", func(c *echo.Context) error {
		return c.JSON(200, httpresponse.Response{
			Success: true,
			Message: "Register route working",
		})
	})
}
