package zone

import (
	"spot-sync/internal/domain/zone/dto"

	"github.com/google/uuid"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) CreateZone(req dto.CreateRequest) (*dto.ZoneResponse, error) {
	zone := &Zone{
		Name:          req.Name,
		Type:          ZoneType(req.Type),
		TotalCapacity: req.TotalCapacity,
		PricePerHour:  req.PricePerHour,
	}

	if err := s.repo.Create(zone); err != nil {
		return nil, err
	}

	return zone.toResponse(), nil
}

func (s *service) GetAllZones() ([]dto.ZoneWithAvailability, error) {
	zones, err := s.repo.FindAllWithAvailability()
	if err != nil {
		return nil, err
	}

	res := make([]dto.ZoneWithAvailability, 0, len(zones))
	for _, z := range zones {
		res = append(res, dto.ZoneWithAvailability{
			Id:             z.Id,
			Name:           z.Name,
			Type:           z.Type,
			TotalCapacity:  z.TotalCapacity,
			AvailableSpots: z.AvailableSpots,
			PricePerHour:   z.PricePerHour,
			CreatedAt:      z.CreatedAt,
		})
	}

	return res, nil
}

func (s *service) GetZoneById(id uuid.UUID) (*dto.ZoneWithAvailability, error) {
	zone, err := s.repo.FindByIdWithAvailability(id)

	if err != nil {
		return nil, err
	}

	res := &dto.ZoneWithAvailability{
		Id:             zone.Id,
		Name:           zone.Name,
		Type:           zone.Type,
		TotalCapacity:  zone.TotalCapacity,
		AvailableSpots: zone.AvailableSpots,
		PricePerHour:   zone.PricePerHour,
		CreatedAt:      zone.CreatedAt,
	}

	return res, nil
}
