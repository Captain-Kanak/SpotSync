package reservation

import (
	"errors"
	"spot-sync/internal/domain/zone"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	Create(reservation *Reservation) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

var (
	ErrZoneFull        = errors.New("zone is fully booked")
	ErrAlreadyReserved = errors.New("license plate already has an active reservation")
)

func (r *repository) Create(reservation *Reservation) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var z zone.Zone

		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where(&zone.Zone{Id: reservation.Id}).First(&z).Error; err != nil {
			return err
		}

		var activeCount int64

		if err := tx.Model(&Reservation{}).
			Where(&Reservation{ZoneId: z.Id, Status: ACTIVE}).
			Count(&activeCount).Error; err != nil {
			return err
		}

		if activeCount >= int64(z.TotalCapacity) {
			return ErrZoneFull
		}

		var existing int64

		if err := tx.Model(&Reservation{}).
			Where(&Reservation{LicensePlate: reservation.LicensePlate, Status: ACTIVE}).
			Count(&existing).Error; err != nil {
			return err
		}

		if existing > 0 {
			return ErrAlreadyReserved
		}

		return tx.Create(reservation).Error
	})
}
