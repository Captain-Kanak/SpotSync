package zone

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func Routes(db *gorm.DB, api *echo.Group) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	api.POST("/zones", handler.CreateZone)
	api.GET("/zones", handler.GetAllZones)
	api.GET("/zones/:id", handler.GetZoneById)
}
