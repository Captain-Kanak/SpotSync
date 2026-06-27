package reservation

import (
	"errors"
	"spot-sync/internal/domain/reservation/dto"
	"spot-sync/internal/domain/user"

	"github.com/google/uuid"
	"gorm.io/gorm"
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

func (s *service) GetMyReservations(userId uuid.UUID) ([]dto.MyReservationResponse, error) {
	reservations, err := s.repo.FindByUserId(userId)

	if err != nil {
		return nil, err
	}

	res := make([]dto.MyReservationResponse, 0, len(reservations))

	for _, r := range reservations {
		res = append(res, dto.MyReservationResponse{
			Id:           r.Id,
			LicensePlate: r.LicensePlate,
			Status:       string(r.Status),
			Zone: dto.ReservationZoneInfo{
				Id:   r.Zone.Id,
				Name: r.Zone.Name,
				Type: string(r.Zone.Type),
			},
			CreatedAt: r.CreatedAt,
		})
	}

	return res, nil
}

func (s *service) CancelReservation(id uuid.UUID, userId uuid.UUID, role user.UserRole) error {
	reservation, err := s.repo.FindById(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrReservationNotFound
		}

		return err
	}

	if role != user.ADMIN && reservation.UserId != userId {
		return ErrNotOwner
	}

	if reservation.Status != ACTIVE {
		return ErrInvalidStatusTransition
	}

	return s.repo.UpdateStatus(id, CANCELED)
}
