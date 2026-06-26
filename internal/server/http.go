package server

import (
	"fmt"
	"spot-sync/internal/config"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"gorm.io/gorm"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.ErrBadRequest.Wrap(err)
	}

	return nil
}

func Start(env *config.Env, db *gorm.DB) {
	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Validator = &CustomValidator{validator: validator.New()}

	Routes(e, db)

	port := fmt.Sprintf(":%s", env.PORT)

	if err := e.Start(port); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
