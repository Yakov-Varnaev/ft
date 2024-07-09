package repository

import (
	"github.com/Yakov-Varnaev/ft/internal/model"
	"github.com/Yakov-Varnaev/ft/pkg/repository/utils"
)

type GroupRepository interface {
	Create(info *model.GroupInfo) (*model.Group, error)
	List() ([]*model.Group, error)
	Exists(filters utils.Filters) (bool, error)
}
