package service

import (
	"github.com/Yakov-Varnaev/ft/internal/model"
	"github.com/Yakov-Varnaev/ft/pkg/pagination"
)

type GroupService interface {
	Create(info *model.GroupInfo) (*model.Group, error)
	List(pagination.Pagination) (*pagination.Page[*model.Group], error)
	Update(id string, info *model.GroupInfo) (*model.Group, error)
	Delete(id string) error
}

type CategoryService interface {
	Create(info *model.CategoryInfo) (*model.Category, error)
	List(pagination.Pagination) (*pagination.Page[*model.Category], error)
	Update(id string, info *model.CategoryInfo) (*model.Category, error)
	Delete(id string) error
}

type SpendingsService interface {
	Create(info *model.SpendingsInfo) (*model.Spendings, error)
	List(pagination.Pagination) (*pagination.Page[*model.Spendings], error)
	// Update(id string, info *model.SpendingsInfo) (*model.Spendings, error)
	// Delete(id string) error
}
