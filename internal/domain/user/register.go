package user

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func Routes(db *gorm.DB, api *echo.Group) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	api.POST("/auth/register", handler.RegisterUser)
	api.POST("/auth/login", handler.LoginUser)
}
