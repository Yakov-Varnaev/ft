package service

import (
	"github.com/Yakov-Varnaev/ft/internal/model"
	"github.com/Yakov-Varnaev/ft/internal/repository"
	def "github.com/Yakov-Varnaev/ft/internal/service"
	pg "github.com/Yakov-Varnaev/ft/pkg/pagination"
	"github.com/Yakov-Varnaev/ft/pkg/repository/utils"
	// webErrors "github.com/Yakov-Varnaev/ft/pkg/web/errors"
	"github.com/go-playground/validator/v10"
)

var _ def.SpendingsService = (*service)(nil)

type service struct {
	r repository.SpendingsRepository
	v *validator.Validate
}

func New(repo repository.SpendingsRepository) *service {
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

func (s *service) Create(info *model.SpendingsInfo) (*model.Spendings, error) {
	if err := s.v.Struct(info); err != nil {
		return nil, err
	}
	sp, err := s.r.Create(info)
	if err != nil {
		return nil, err
	}
	return sp, nil
}

func (s *service) List(p pg.Pagination) (*pg.Page[*model.Spendings], error) {
	spendings, total, err := s.r.List(p)
	if err != nil {
		return nil, err
	}
	return &pg.Page[*model.Spendings]{
		Total: total,
		Data:  spendings,
	}, nil
}

//
// func (s *service) Update(id string, info *model.CategoryInfo) (*model.Category, error) {
// 	exists, err := s.r.Exists(utils.Filters{"id": id})
// 	if err != nil {
// 		return nil, err
// 	}
// 	if !exists {
// 		return nil, &webErrors.NotFoundError{
// 			Message: "Category with given id not found",
// 		}
// 	}
//
// 	if err := s.v.Struct(info); err != nil {
// 		return nil, &webErrors.BadRequest{Message: err.Error()}
// 	}
// 	updatedCategory, err := s.r.Update(id, info)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return updatedCategory, nil
// }
//
// func (s *service) Delete(id string) error {
// 	exists, err := s.r.Exists(utils.Filters{"id": id})
// 	if err != nil {
// 		return err
// 	}
// 	if !exists {
// 		return &webErrors.NotFoundError{
// 			Message: "Category with given id not found",
// 		}
// 	}
// 	err = s.r.Delete(id)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
