package repository

import (
	"github.com/Yakov-Varnaev/ft/internal/repository/group/model"
	"github.com/Yakov-Varnaev/ft/pkg/pagination"
	"github.com/jmoiron/sqlx"
)

type Repositroy struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repositroy {
	return &Repositroy{db}
}

func (r *Repositroy) Create(data *model.GroupInfo) (*model.Group, error) {
	var group model.Group
	err := r.db.QueryRowx(
		`INSERT INTO groups (name) VALUES ($1) RETURNING id, name`, data.Name,
	).Scan(&group.UUID, &group.Name)
	if err != nil {
		return nil, err
	}
	return &group, nil
}

func (r *Repositroy) List(pg pagination.Pagination) ([]*model.Group, int, error) {
	var count int
	err := r.db.QueryRowx(
		"SELECT COUNT(*) from groups",
	).Scan(&count)
	if err != nil {
		return nil, 0, err
	}
	rows, err := r.db.Queryx(
		"SELECT id, name FROM groups LIMIT $1 OFFSET $2",
		pg.Limit, pg.Offset,
	)
	if err != nil {
		return nil, 0, err
	}
	groups := make([]*model.Group, 0)
	for rows.Next() {
		var group model.Group
		err = rows.Scan(&group.UUID, &group.Name)
		if err != nil {
			return nil, 0, err
		}
		groups = append(groups, &group)
	}
	return groups, count, nil
}
