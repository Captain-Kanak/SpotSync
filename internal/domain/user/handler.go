package user

import (
	"fmt"
	"net/http"
	"spot-sync/internal/domain/user/dto"
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
	var req dto.CreateRequest

	if err := c.Bind(&req); err != nil {
		fmt.Println(err)

		return c.JSON(http.StatusBadRequest, httpresponse.Response{
			Success: false,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	if err := c.Validate(&req); err != nil {
		fmt.Println(err.Error())

		return c.JSON(http.StatusBadRequest, httpresponse.Response{
			Success: false,
			Message: "Validation failed!",
			Error:   err.Error(),
		})
	}

	res, err := h.service.RegisterUser(req)

	if err != nil {
		fmt.Println(err)

		return c.JSON(http.StatusBadRequest, httpresponse.Response{
			Success: false,
			Message: "User registration failed",
		})
	}

	return c.JSON(http.StatusCreated, httpresponse.Response{
		Success: true,
		Message: "User registered successfully",
		Data:    res,
	})
}

func (h *handler) LoginUser(c *echo.Context) error {
	return c.JSON(http.StatusOK, httpresponse.Response{
		Success: true,
		Message: "Login successful",
	})
}
