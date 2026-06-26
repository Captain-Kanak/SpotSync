package user

import (
	"spot-sync/internal/domain/user/dto"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) RegisterUser(req dto.CreateRequest) (*dto.UserResponse, error) {
	user := &User{
		Name:  req.Name,
		Email: req.Email,
	}

	if err := user.hashPassword(req.Password); err != nil {
		return nil, err
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user.toResponse(), nil
}
