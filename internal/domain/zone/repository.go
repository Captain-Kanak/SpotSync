package zone

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Create(zone *Zone) error
	GetAll() ([]Zone, error)
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

func (r *repository) GetAll() ([]Zone, error) {
	zones := []Zone{}

	if err := r.db.Find(&zones).Error; err != nil {
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
