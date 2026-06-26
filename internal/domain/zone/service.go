package zone

import (
	"fmt"
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

func (s *service) GetAllZones() ([]dto.ZoneResponse, error) {
	zones, err := s.repo.GetAll()

	if err != nil {
		return nil, err
	}

	fmt.Println("Zones result", zones)

	res := make([]dto.ZoneResponse, len(zones))

	for i, zone := range zones {
		res[i] = *zone.toResponse()
	}

	return res, nil
}

func (s *service) GetZoneById(id uuid.UUID) (*dto.ZoneResponse, error) {
	zone, err := s.repo.GetById(id)

	if err != nil {
		return nil, err
	}

	return zone.toResponse(), nil
}
