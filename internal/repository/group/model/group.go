package model

import "github.com/Yakov-Varnaev/ft/pkg/repository/utils"

type GroupInfo struct {
	Name string `db:"name"`
}

type Group struct {
	GroupInfo `db:"group_info"`
	UUID      string `db:"id"`
}

func (g Group) FromRow(src utils.Scanner) (*Group, error) {

	if err := src.Scan(&g.UUID, &g.Name); err != nil {
		return nil, err
	}
	return &g, nil
}
