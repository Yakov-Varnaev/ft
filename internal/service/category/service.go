package service

import (
	"github.com/Yakov-Varnaev/ft/internal/model"
	"github.com/Yakov-Varnaev/ft/internal/repository"
	def "github.com/Yakov-Varnaev/ft/internal/service"
	pg "github.com/Yakov-Varnaev/ft/pkg/pagination"
	"github.com/Yakov-Varnaev/ft/pkg/repository/utils"
	webErrors "github.com/Yakov-Varnaev/ft/pkg/web/errors"
	"github.com/go-playground/validator/v10"
)

var _ def.CategoryService = (*service)(nil)

type service struct {
	r repository.CategoryRepository
	v *validator.Validate
}

func New(repo repository.CategoryRepository) *service {
	validate := validator.New()
	srv := &service{repo, validate}
	validate.RegisterValidation("unique-name", srv.checkNameIsUnique)
	return srv
}

func (s *service) checkNameIsUnique(field validator.FieldLevel) bool {
	exists, err := s.r.Exists(utils.Filters{"name": field.Field().String()})
	if err != nil {
		return false
	}
	return !exists
}

func (s *service) Create(info *model.CategoryInfo) (*model.Category, error) {
	if err := s.v.Struct(info); err != nil {
		return nil, err
	}
	category, err := s.r.Create(info)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (s *service) List(p pg.Pagination) (*pg.Page[*model.Category], error) {
	groups, total, err := s.r.List(p)
	if err != nil {
		return nil, err
	}
	return &pg.Page[*model.Category]{
		Total: total,
		Data:  groups,
	}, nil

}

func (s *service) Delete(id string) error {
	exists, err := s.r.Exists(utils.Filters{"id": id})
	if err != nil {
		return err
	}
	if !exists {
		return &webErrors.NotFoundError{
			Message: "Group with given id not found",
		}
	}
	err = s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
