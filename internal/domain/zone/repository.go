package zone

import "gorm.io/gorm"

type Repository interface {
	Create(zone *Zone) error
	GetAll() ([]Zone, error)
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
