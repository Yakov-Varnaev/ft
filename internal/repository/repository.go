package repository

import "github.com/Yakov-Varnaev/ft/internal/model"

type GroupRepository interface {
	Create(info *model.GroupInfo) (*model.Group, error)
	List() ([]*model.Group, error)
}
