package reservation

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func Routes(db *gorm.DB, api *echo.Group) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	api.POST("/reservations", handler.ReserveSpot)
}
