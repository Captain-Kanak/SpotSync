package dto

type CreateRequest struct {
	Name          string  `json:"name" validate:"required"`
	Type          string  `json:"type"`
	TotalCapacity int     `json:"total_capacity" validate:"required"`
	PricePerHour  float64 `json:"price_per_hour" validate:"required"`
}
