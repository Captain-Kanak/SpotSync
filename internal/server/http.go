package server

import (
	"errors"
	"fmt"
	"spot-sync/internal/config"
	"strings"

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
		if ve, ok := err.(validator.ValidationErrors); ok {
			return echo.ErrBadRequest.Wrap(errors.New(formatValidationErrors(ve)))
		}
		return echo.ErrBadRequest.Wrap(err)
	}

	return nil
}

func formatValidationErrors(ve validator.ValidationErrors) string {
	var messages []string

	for _, e := range ve {
		switch {
		case e.Field() == "Type" && e.Tag() == "oneof":
			messages = append(messages, fmt.Sprintf(
				"Type must be one of: GENERAL, EV_CHARGING, COVERED (uppercase required, e.g. use EV_CHARGING instead of %s)",
				e.Value(),
			))
		default:
			messages = append(messages, fmt.Sprintf("%s failed on '%s' validation", e.Field(), e.Tag()))
		}
	}

	return strings.Join(messages, "; ")
}

func Start(env *config.Env, db *gorm.DB) {
	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Validator = &CustomValidator{validator: validator.New()}

	Routes(e, db, env)

	port := fmt.Sprintf(":%s", env.PORT)

	if err := e.Start(port); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
