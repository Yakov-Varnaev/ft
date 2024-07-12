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
