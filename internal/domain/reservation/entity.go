package reservation

import (
	"spot-sync/internal/domain/reservation/dto"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReservationStatus string

const (
	Reserved  ReservationStatus = "ACTIVE"
	Completed ReservationStatus = "COMPLETED"
	Canceled  ReservationStatus = "CANCELED"
)

type Reservation struct {
	Id           uuid.UUID         `json:"id" gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	UserId       uuid.UUID         `json:"user_id" gorm:"type:uuid;not null"`
	ZoneId       uuid.UUID         `json:"zone_id" gorm:"type:uuid;not null"`
	LicensePlate string            `json:"license_plate" gorm:"type:varchar(255);not null"`
	Status       ReservationStatus `json:"status" gorm:"type:reservation_status;default:'ACTIVE';not null"`
	CreatedAt    time.Time         `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt    time.Time         `json:"updated_at" gorm:"type:timestamp"`
	DeletedAt    gorm.DeletedAt    `json:"-" gorm:"type:timestamp;index"`
}

func (r *Reservation) toResponse() *dto.ReservationResponse {
	return &dto.ReservationResponse{
		Id:           r.Id,
		UserId:       r.UserId,
		ZoneId:       r.ZoneId,
		LicensePlate: r.LicensePlate,
		Status:       string(r.Status),
		CreatedAt:    r.CreatedAt,
		UpdatedAt:    r.UpdatedAt,
	}
}
