package service

import (
	repository "github.com/Yakov-Varnaev/ft/internal/repository/group"
	"github.com/Yakov-Varnaev/ft/internal/service/group/model"
	"github.com/Yakov-Varnaev/ft/pkg/pagination"
	repoUtils "github.com/Yakov-Varnaev/ft/pkg/repository/utils"
	webErrors "github.com/Yakov-Varnaev/ft/pkg/web/errors"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type Service struct {
	repo *repository.Repository
}

// check if group with given id exists, returns NotFoundError
// if no such id exists.
func (s *Service) checkIdExists(id string) (bool, error) {
	exists, err := s.repo.Exists(repoUtils.Filters{"id": id})
	if err != nil {
		return false, &webErrors.InternalServerError{Err: err}
	}
	if !exists {
		return false, &webErrors.NotFoundError{
			Message: "Group with given id not found",
		}
	}
	return true, nil
}

func (s *Service) validateName(field validator.FieldLevel) bool {
	exists, err := s.repo.CheckNameExists(field.Field().String())
	if err != nil {
		return true
	}
	return !exists
}

func New(repo *repository.Repository) *Service {
	s := &Service{repo}
	validate.RegisterValidation("unique-name", s.validateName)
	return s
}

func (s *Service) Create(data *model.GroupInfo) (*model.Group, error) {
	if data == nil {
		panic("Data cannot be nil.")
	}
	if err := validate.Struct(data); err != nil {
		return nil, err
	}
	group, err := s.repo.Create(model.ToRepoGroupInfo(data))
	if err != nil {
		return nil, &webErrors.InternalServerError{Err: err}
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

func (s *Service) Update(id string, data *model.GroupInfo) (*model.Group, error) {
	if _, err := s.checkIdExists(id); err != nil {
		return nil, err
	}
	if err := validate.Struct(data); err != nil {
		return nil, err
	}
	group, err := s.repo.Update(id, model.ToRepoGroupInfo(data))
	if err != nil {
		return nil, err
	}
	return model.FromRepoGroup(group), nil
}

func (s *Service) Delete(id string) error {
	if _, err := s.checkIdExists(id); err != nil {
		return err
	}
	return s.repo.Delete(id)
}
