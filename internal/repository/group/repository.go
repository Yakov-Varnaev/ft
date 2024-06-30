package repository

import (
	"fmt"
	"strings"

	"github.com/Yakov-Varnaev/ft/internal/repository/group/model"
	"github.com/Yakov-Varnaev/ft/pkg/pagination"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) CheckNameExists(name string) (bool, error) {
	return r.Exists(map[string]interface{}{"name": name})
}

func (r *Repository) Exists(filters map[string]interface{}) (bool, error) {
	q := "SELECT id FROM groups WHERE "
	strFilters := []string{}
	idx := 1
	params := []interface{}{}
	// TODO: maybe it's a good idea to move this logic to separate function
	for k, v := range filters {
		strFilters = append(strFilters, fmt.Sprintf("%s = $%d", k, idx))
		idx++
		params = append(params, v)
	}
	if len(strFilters) > 1 {
		q = q + strings.Join(strFilters, " AND ")
	} else {
		q = q + strFilters[0]
	}

	var exists bool
	err := r.db.QueryRowx(
		fmt.Sprintf("SELECT EXISTS(%s)", q),
		params...,
	).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (r *Repository) Create(data *model.GroupInfo) (*model.Group, error) {
	var group model.Group
	err := r.db.QueryRowx(
		`INSERT INTO groups (name) VALUES ($1) RETURNING id, name`, data.Name,
	).Scan(&group.UUID, &group.Name)
	if err != nil {
		return nil, err
	}
	return &group, nil
}

func (r *Repository) List(pg pagination.Pagination) ([]*model.Group, int, error) {
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

func (r *Repository) Update(id string, data *model.GroupInfo) (*model.Group, error) {
	var group model.Group

	err := r.db.QueryRowx(
		"UPDATE groups SET name=$1 WHERE id=$2 RETURNING id, name",
		data.Name, id,
	).Scan(&group.UUID, &group.Name)
	if err != nil {
		return nil, err
	}
	return &group, nil
}

func (r *Repository) Delete(id string) error {
	_, err := r.db.Exec(
		"DELETE FROM groups where id=$1",
		id,
	)
	if err != nil {
		return err
	}
	return nil
}
