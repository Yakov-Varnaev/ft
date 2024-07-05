package service

import (
	repository "github.com/Yakov-Varnaev/ft/internal/repository/category"
	groupRepository "github.com/Yakov-Varnaev/ft/internal/repository/group"
	"github.com/Yakov-Varnaev/ft/internal/service/category/model"
	"github.com/Yakov-Varnaev/ft/pkg/pagination"
	"github.com/Yakov-Varnaev/ft/pkg/repository/utils"
	webErrors "github.com/Yakov-Varnaev/ft/pkg/web/errors"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type Service struct {
	repo      *repository.Repository
	groupRepo *groupRepository.Repository
}

func New(repo *repository.Repository, groupRepo *groupRepository.Repository) *Service {
	s := &Service{repo, groupRepo}
	validate.RegisterValidation("unique-name", s.validateName)
	validate.RegisterValidation("group-exists", s.validateGroupId)
	return s
}

func (s *Service) checkIdExists(id string) (bool, error) {
	exists, err := s.repo.Exists(utils.Filters{"id": id})
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
	exists, err := s.repo.Exists(
		utils.Filters{"name": field.Field().String()},
	)
	if err != nil {
		return false
	}
	return !exists
}

func (s *Service) validateGroupId(field validator.FieldLevel) bool {
	exists, err := s.groupRepo.Exists(utils.Filters{"id": field.Field().String()})
	if err != nil {
		return false
	}
	return exists
}

func (s *Service) Create(data *model.CategoryInfo) (*model.Category, error) {
	if data == nil {
		panic("Data cannot be nil.")
	}
	if err := validate.Struct(data); err != nil {
		return nil, err
	}
	category, err := s.repo.Create(model.ToRepoCategoryInfo(data))
	if err != nil {
		return nil, &webErrors.InternalServerError{Err: err}
	}
	return model.FromRepoCategory(category), nil

}

func (s *Service) List(pg pagination.Pagination) (*pagination.Page[*model.Category], error) {
	categories, count, err := s.repo.List(pg)
	if err != nil {
		return nil, err
	}

	serviceCategories := make([]*model.Category, len(categories))
	for i, category := range categories {
		serviceCategories[i] = model.FromRepoCategory(category)
	}
	return &pagination.Page[*model.Category]{Data: serviceCategories, Total: count}, nil
}

func (s *Service) Update(id string, data *model.CategoryInfo) (*model.Category, error) {
	if _, err := s.checkIdExists(id); err != nil {
		return nil, err
	}
	if err := validate.Struct(data); err != nil {
		return nil, err
	}
	category, err := s.repo.Update(id, model.ToRepoCategoryInfo(data))
	if err != nil {
		return nil, err
	}
	return model.FromRepoCategory(category), nil
}

func (s *Service) Delete(id string) error {
	if _, err := s.checkIdExists(id); err != nil {
		return err
	}
	return s.repo.Delete(id)
}
