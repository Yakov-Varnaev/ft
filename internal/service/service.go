package service

import "github.com/Yakov-Varnaev/ft/internal/model"

type GroupService interface {
	Create(info *model.GroupInfo) (*model.Group, error)
	List() ([]*model.Group, error)
}
