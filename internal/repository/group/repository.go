package repository

import (
	"fmt"

	"github.com/Yakov-Varnaev/ft/internal/model"
	def "github.com/Yakov-Varnaev/ft/internal/repository"
	"github.com/Yakov-Varnaev/ft/internal/repository/group/converter"
	repoModel "github.com/Yakov-Varnaev/ft/internal/repository/group/model"

	"github.com/Yakov-Varnaev/ft/pkg/repository/utils"
	"github.com/jmoiron/sqlx"
)

var _ def.GroupRepository = (*repository)(nil)

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *repository {
	return &repository{db}
}

const createQuery string = `INSERT INTO groups (name) VALUES ($1) RETURNING id, name`

func (r *repository) Exists(filters utils.Filters) (bool, error) {
	q := "SELECT id FROM groups WHERE "

	whereQ, params := filters.Prepare() // I really wany simple sql builder...
	q = q + whereQ

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

func (r *repository) Create(info *model.GroupInfo) (*model.Group, error) {
	var group repoModel.Group
	data := converter.ToRepoGroupInfo(info)
	err := r.db.QueryRowx(createQuery, data.Name).Scan(&group.ID, &group.Name)
	if err != nil {
		return nil, err
	}
	return converter.FromRepoGroup(&group), nil
}

func (r *repository) List() ([]*model.Group, error) {
	// TODO: add pagination
	// err := r.db.QueryRowx(
	// 	"SELECT COUNT(*) from groups",
	// ).Scan(&count)
	// if err != nil {
	// 	return nil, 0, err
	// }
	rows, err := r.db.Queryx(
		"SELECT id, name FROM groups LIMIT $1 OFFSET $2",
		100, 0,
	)
	if err != nil {
		return nil, err
	}
	groups := make([]*model.Group, 0)
	for rows.Next() {
		var group repoModel.Group
		err = rows.Scan(&group.ID, &group.Name)
		if err != nil {
			return nil, err
		}
		groups = append(groups, converter.FromRepoGroup(&group))
	}
	return groups, nil
}

const updateQuery = `UPDATE groups SET name = $1 WHERE id = $2 RETURNING id, name`

func (r *repository) Update(id string, info *model.GroupInfo) (*model.Group, error) {
	row := r.db.QueryRow(updateQuery, info.Name, id)

	var group repoModel.Group
	if err := row.Scan(&group.ID, &group.Name); err != nil {
		return nil, err
	}

	return converter.FromRepoGroup(&group), nil
}
