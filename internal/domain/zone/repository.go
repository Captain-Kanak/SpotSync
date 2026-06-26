package zone

import "gorm.io/gorm"

type Repository interface {
	Create(zone *Zone) error
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
