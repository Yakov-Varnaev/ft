package repository

import (
	"github.com/Yakov-Varnaev/ft/internal/model"
	pg "github.com/Yakov-Varnaev/ft/pkg/pagination"
	"github.com/Yakov-Varnaev/ft/pkg/repository/utils"
)

type GroupRepository interface {
	Create(info *model.GroupInfo) (*model.Group, error)
	List(pg pg.Pagination) ([]*model.Group, int, error)
	Exists(filters utils.Filters) (bool, error)
	Update(id string, info *model.GroupInfo) (*model.Group, error)
	Delete(id string) error
}

type CategoryRepository interface {
	Create(info *model.CategoryInfo) (*model.Category, error)
	List(pg pg.Pagination) ([]*model.Category, int, error)
	Exists(filters utils.Filters) (bool, error)
	Update(id string, info *model.CategoryInfo) (*model.Category, error)
	Delete(id string) error
}

type SpendingsRepository interface {
	Create(info *model.SpendingsInfo) (*model.Spendings, error)
	List(pg pg.Pagination) ([]*model.Spendings, int, error)
	Exists(filters utils.Filters) (bool, error)
	// Update(id string, info *model.SpendingsInfo) (*model.Spendings, error)
	// Delete(id string) error
}
