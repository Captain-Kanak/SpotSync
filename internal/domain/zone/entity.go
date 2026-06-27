package zone

import (
	"spot-sync/internal/domain/zone/dto"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ZoneType string

const (
	GENERAL     ZoneType = "GENERAL"
	EV_CHARGING ZoneType = "EV_CHARGING"
	COVERED     ZoneType = "COVERED"
)

type Zone struct {
	Id            uuid.UUID      `json:"id" gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	Name          string         `json:"name" gorm:"type:varchar(255);not null"`
	Type          ZoneType       `json:"type" gorm:"type:zone_type;default:'GENERAL';not null"`
	TotalCapacity int            `json:"total_capacity" gorm:"type:int;not null"`
	PricePerHour  float64        `json:"price_per_hour" gorm:"type:float;not null"`
	CreatedAt     time.Time      `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt     time.Time      `json:"updated_at" gorm:"type:timestamp"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"type:timestamp;index"`
}

func (z *Zone) toResponse() *dto.ZoneResponse {
	return &dto.ZoneResponse{
		Id:             z.Id,
		Name:           z.Name,
		Type:           string(z.Type),
		TotalCapacity:  z.TotalCapacity,
		AvailableSpots: z.TotalCapacity,
		PricePerHour:   z.PricePerHour,
		CreatedAt:      z.CreatedAt,
	}
}
