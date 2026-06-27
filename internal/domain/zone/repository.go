package zone

import (
	"spot-sync/internal/domain/zone/dto"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Create(zone *Zone) error
	FindAllWithAvailability() ([]dto.ZoneWithAvailability, error)
	FindByIdWithAvailability(id uuid.UUID) (*dto.ZoneWithAvailability, error)
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

	err := r.db.Table("zones AS z").
		Select(`
			z.id,
			z.name,
			z.type,
			z.total_capacity,
			z.total_capacity - COALESCE((
				SELECT COUNT(*) FROM reservations r
				WHERE r.zone_id = z.id
				AND r.status = ?
				AND r.deleted_at IS NULL
			), 0) AS available_spots,
			z.price_per_hour,
			z.created_at
		`, "ACTIVE").
		Find(&zones).Error

	if err != nil {
		return nil, err
	}

	return zones, nil
}

func (r *repository) FindByIdWithAvailability(id uuid.UUID) (*dto.ZoneWithAvailability, error) {
	var zone dto.ZoneWithAvailability

	err := r.db.Table("zones AS z").
		Select(`
			z.id,
			z.name,
			z.type,
			z.total_capacity,
			z.total_capacity - COALESCE((
				SELECT COUNT(*) FROM reservations r
				WHERE r.zone_id = z.id
				AND r.status = ?
				AND r.deleted_at IS NULL
			), 0) AS available_spots,
			z.price_per_hour,
			z.created_at
		`, "ACTIVE").
		Where("z.id = ?", id).
		First(&zone).Error

	if err != nil {
		return nil, err
	}

	return &zone, nil
}
