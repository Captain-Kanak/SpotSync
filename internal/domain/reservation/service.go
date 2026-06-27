package reservation

import (
	"spot-sync/internal/domain/reservation/dto"

	"github.com/google/uuid"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) ReserveSpot(req *dto.CreateRequest, userId uuid.UUID) (*dto.ReservationResponse, error) {
	var reservation = &Reservation{
		UserId:       userId,
		ZoneId:       req.ZoneId,
		LicensePlate: req.LicensePlate,
	}

	if err := s.repo.Create(reservation); err != nil {
		return nil, err
	}

	return reservation.toResponse(), nil
}
