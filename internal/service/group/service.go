package service

import (
	"github.com/Yakov-Varnaev/ft/internal/model"
	"github.com/Yakov-Varnaev/ft/internal/repository"
	def "github.com/Yakov-Varnaev/ft/internal/service"
)

var _ def.GroupService = (*service)(nil)

type service struct {
	r repository.GroupRepository
}

func New(r repository.GroupRepository) *service {
	return &service{r}
}

func (s *service) Create(info *model.GroupInfo) (*model.Group, error) {
	group, err := s.r.Create(info)
	if err != nil {
		// TODO: wrap in httpError
		return nil, err
	}
	return group, nil
}

func (s *service) List() ([]*model.Group, error) {
	groups, err := s.r.List()
	if err != nil {
		return nil, err
	}
	return groups, nil
}
