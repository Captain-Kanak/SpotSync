package dto

import (
	"time"

	"github.com/google/uuid"
)

type ReservationResponse struct {
	Id           uuid.UUID `json:"id"`
	UserId       uuid.UUID `json:"user_id"`
	ZoneId       uuid.UUID `json:"zone_id"`
	LicensePlate string    `json:"license_plate"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type MyReservationResponse struct {
	Id           uuid.UUID           `json:"id"`
	LicensePlate string              `json:"license_plate"`
	Status       string              `json:"status"`
	Zone         ReservationZoneInfo `json:"zone"`
	CreatedAt    time.Time           `json:"created_at"`
}

type ReservationZoneInfo struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Type string    `json:"type"`
}

type ZoneWithAvailability struct {
	Id             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Type           string    `json:"type"`
	TotalCapacity  int       `json:"total_capacity"`
	AvailableSpots int       `json:"available_spots"`
	PricePerHour   float64   `json:"price_per_hour"`
	CreatedAt      time.Time `json:"created_at"`
}
