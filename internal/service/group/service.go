package service

import (
	"github.com/Yakov-Varnaev/ft/internal/model"
	"github.com/Yakov-Varnaev/ft/internal/repository"
	def "github.com/Yakov-Varnaev/ft/internal/service"
	"github.com/Yakov-Varnaev/ft/pkg/repository/utils"
	webErrors "github.com/Yakov-Varnaev/ft/pkg/web/errors"
	"github.com/go-playground/validator/v10"
)

var _ def.GroupService = (*service)(nil)

type service struct {
	r repository.GroupRepository
	v *validator.Validate
}

func New(r repository.GroupRepository) *service {
	validate := validator.New()
	s := &service{r, validate}
	validate.RegisterValidation("unique-name", s.checkNameIsUnique)
	return s
}

func (s *service) checkNameIsUnique(field validator.FieldLevel) bool {
	exists, err := s.r.Exists(utils.Filters{"name": field.Field().String()})
	if err != nil {
		return false
	}
	return !exists
}

func (s *service) Create(info *model.GroupInfo) (*model.Group, error) {
	if err := s.v.Struct(info); err != nil {
		return nil, err
	}
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

func (s *service) Update(id string, info *model.GroupInfo) (*model.Group, error) {
	exists, err := s.r.Exists(utils.Filters{"id": id})
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, &webErrors.NotFoundError{
			Message: "Group with given id not found",
		}
	}
	updatedGroup, err := s.r.Update(id, info)
	if err != nil {
		return nil, err
	}
	return updatedGroup, nil
}
