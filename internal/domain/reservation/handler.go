package reservation

import (
	"fmt"
	"net/http"
	"spot-sync/internal/domain/reservation/dto"
	"spot-sync/internal/httpresponse"

	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

type handler struct {
	service *service
}

func NewHandler(service *service) *handler {
	return &handler{service}
}

func (h *handler) ReserveSpot(c *echo.Context) error {
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

	res, err := h.service.ReserveSpot(&req, c.Get("user_id").(uuid.UUID))

	if err != nil {
		fmt.Println(err)

		return c.JSON(http.StatusBadRequest, httpresponse.Response{
			Success: false,
			Message: "Failed to reserve spot",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, httpresponse.Response{
		Success: true,
		Message: "Spot reserved successfully",
		Data:    res,
	})
}
