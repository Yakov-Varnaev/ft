package model

import "github.com/Yakov-Varnaev/ft/pkg/repository/utils"

type CategoryInfo struct {
	Name    string `db:"name"`
	GroupId string `db:"group_id"`
}

type Category struct {
	UUID string `db:"id"`
	CategoryInfo
}

func (c Category) FromRow(src utils.Scanner) (*Category, error) {
	if err := src.Scan(&c.UUID, &c.Name, &c.GroupId); err != nil {
		return nil, err
	}
	return &c, nil
}
