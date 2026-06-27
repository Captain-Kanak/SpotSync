package zone

import (
	"spot-sync/internal/domain/zone/dto"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Create(zone *Zone) error
	FindAllWithAvailability() ([]dto.ZoneWithAvailability, error)
	GetById(id uuid.UUID) (*Zone, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(zone *Zone) error {
	return r.db.Create(zone).Error
}

func (r *repository) FindAllWithAvailability() ([]dto.ZoneWithAvailability, error) {
	var zones []dto.ZoneWithAvailability

	err := r.db.Model(&Zone{}).
		Select(`
			parking_zones.id,
			parking_zones.name,
			parking_zones.type,
			parking_zones.total_capacity,
			parking_zones.total_capacity - COALESCE((
				SELECT COUNT(*) FROM reservations
				WHERE reservations.zone_id = parking_zones.id
				AND reservations.status = ?
				AND reservations.deleted_at IS NULL
			), 0) AS available_spots,
			parking_zones.price_per_hour,
			parking_zones.created_at
		`, "ACTIVE").
		Find(&zones).Error

	if err != nil {
		return nil, err
	}

	return zones, nil
}

func (r *repository) GetById(id uuid.UUID) (*Zone, error) {
	zone := &Zone{}

	if err := r.db.Where(&Zone{Id: id}).First(&zone).Error; err != nil {
		return nil, err
	}

	return zone, nil
}
