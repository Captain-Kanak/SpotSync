package server

import (
	"net/http"
	"spot-sync/internal/httpresponse"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func Routes(e *echo.Echo, db *gorm.DB) {
	db.AutoMigrate()

	e.GET("/", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, httpresponse.Response{
			Success: true,
			Message: "Server is running...",
		})
	})
}
