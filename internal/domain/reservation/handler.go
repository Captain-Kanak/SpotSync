package reservation

import (
	"errors"
	"fmt"
	"net/http"
	"spot-sync/internal/auth"
	"spot-sync/internal/domain/reservation/dto"
	"spot-sync/internal/httpresponse"
	"spot-sync/internal/middleware"

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

	claims := c.Get(middleware.ContextUserKey).(*auth.JWTClaims)

	res, err := h.service.ReserveSpot(&req, claims.Id)

	if err != nil {
		switch {
		case errors.Is(err, ErrZoneFull):
			return c.JSON(http.StatusConflict, httpresponse.Response{
				Success: false,
				Message: "This zone is fully booked",
			})

		case errors.Is(err, ErrAlreadyReserved):
			return c.JSON(http.StatusConflict, httpresponse.Response{
				Success: false,
				Message: "This vehicle already has an active reservation",
			})

		default:
			return c.JSON(http.StatusInternalServerError, httpresponse.Response{
				Success: false,
				Message: "Failed to reserve spot",
			})
		}
	}

	return c.JSON(http.StatusOK, httpresponse.Response{
		Success: true,
		Message: "Spot reserved successfully",
		Data:    res,
	})
}
