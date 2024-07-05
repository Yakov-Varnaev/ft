package repository

import (
	"fmt"

	"github.com/Yakov-Varnaev/ft/internal/repository/category/model"
	"github.com/Yakov-Varnaev/ft/pkg/pagination"
	repoUtils "github.com/Yakov-Varnaev/ft/pkg/repository/utils"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) Exists(filters repoUtils.Filters) (bool, error) {
	q := "SELECT id from categories WHERE "

	whereQ, params := filters.Prepare()
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

func (r *Repository) Create(data *model.CategoryInfo) (*model.Category, error) {
	var group model.Category
	err := r.db.QueryRowx(
		"INSERT INTO categories (name, group_id) VALUES ($1, $2) RETURNING id, name, group_id",
		data.Name, data.GroupId,
	).Scan(&group.UUID, &group.Name, &group.GroupId)
	if err != nil {
		return nil, err
	}
	return &group, nil
}

func (r *Repository) List(pg pagination.Pagination) ([]*model.Category, int, error) {
	q := "SELECT id, name, group_id FROM categories"
	return repoUtils.List[model.Category](r.db, "categories", q, pg)
}

func (r *Repository) Update(id string, data *model.CategoryInfo) (*model.Category, error) {
	var category model.Category

	err := r.db.QueryRowx(
		"UPDATE categories SET name=$1, group_id=$2 WHERE id=$3 RETURNING id, name, group_id",
		data.Name, data.GroupId, id,
	).Scan(&category.UUID, &category.Name, &category.GroupId)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *Repository) Delete(id string) error {
	_, err := r.db.Exec(
		"DELETE FROM categories WHERE id=$1",
		id,
	)
	if err != nil {
		return err
	}
	return nil
}
