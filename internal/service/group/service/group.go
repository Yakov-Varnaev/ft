package service

import (
	"fmt"

	repository "github.com/Yakov-Varnaev/ft/internal/repository/group"
	"github.com/Yakov-Varnaev/ft/internal/service/group/model"
	"github.com/Yakov-Varnaev/ft/pkg/pagination"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type Service struct {
	repo *repository.Repositroy
}

func (s *Service) validateName(field validator.FieldLevel) bool {
	exists, err := s.repo.CheckNameExists(field.Field().String())
	if err != nil {
		fmt.Println(err.Error())
		return true
	}
	return !exists
}

func New(repo *repository.Repositroy) *Service {
	s := &Service{repo}
	validate.RegisterValidation("unique-name", s.validateName)
	return s
}

func (s *Service) Create(data *model.GroupInfo) (*model.Group, error) {
	if data == nil {
		return nil, fmt.Errorf("data cannot be nil")
	}
	err := validate.Struct(data)
	if err != nil {
		return nil, err
	}
	group, err := s.repo.Create(model.ToRepoGroupInfo(data))
	if err != nil {
		return nil, err
	}
	return model.FromRepoGroup(group), nil
}

func (s *Service) List(pg pagination.Pagination) (*pagination.Page[*model.Group], error) {
	groups, count, err := s.repo.List(pg)
	if err != nil {
		return nil, err
	}

	serviceGroups := make([]*model.Group, len(groups))
	for i, group := range groups {
		serviceGroups[i] = model.FromRepoGroup(group)
	}
	return &pagination.Page[*model.Group]{Data: serviceGroups, Total: count}, nil
}
