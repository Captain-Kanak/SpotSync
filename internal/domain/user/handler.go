package user

import (
	"net/http"
	"spot-sync/internal/httpresponse"

	"github.com/labstack/echo/v5"
)

type handler struct {
	service *service
}

func NewHandler(service *service) *handler {
	return &handler{service}
}

func (h *handler) RegisterUser(c *echo.Context) error {
	return c.JSON(http.StatusCreated, httpresponse.Response{
		Success: true,
		Message: "User registered successfully",
	})
}

func (h *handler) LoginUser(c *echo.Context) error {
	return c.JSON(http.StatusOK, httpresponse.Response{
		Success: true,
		Message: "Login successful",
	})
}
