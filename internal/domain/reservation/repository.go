package reservation

import "gorm.io/gorm"

type Repository interface {
	Create(reservation *Reservation) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(reservation *Reservation) error {
	return r.db.Create(reservation).Error
}
