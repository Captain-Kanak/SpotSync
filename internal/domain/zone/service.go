package zone

import "spot-sync/internal/domain/zone/dto"

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
