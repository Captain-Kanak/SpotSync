package dto

type CreateRequest struct {
	Name          string  `json:"name" validate:"required"`
	Type          string  `json:"type"`
	TotalCapacity int     `json:"total_capacity" validate:"required"`
	PricePerHour  float64 `json:"price_per_hour" validate:"required"`
}

type UpdateRequest struct {
	Name          string  `json:"name" validate:"omitempty"`
	Type          string  `json:"type" validate:"omitempty,oneof=GENERAL EV_CHARGING COVERED"`
	TotalCapacity int     `json:"total_capacity" validate:"omitempty,gt=0"`
	PricePerHour  float64 `json:"price_per_hour" validate:"omitempty,gt=0"`
}
