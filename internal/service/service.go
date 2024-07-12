package service

import "github.com/Yakov-Varnaev/ft/internal/model"

type GroupService interface {
	Create(info *model.GroupInfo) (*model.Group, error)
	List() ([]*model.Group, error)
	Update(id string, info *model.GroupInfo) (*model.Group, error)
	Delete(id string) error
}
