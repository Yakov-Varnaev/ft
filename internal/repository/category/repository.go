package repository

import (
	"fmt"

	"github.com/Yakov-Varnaev/ft/internal/model"
	def "github.com/Yakov-Varnaev/ft/internal/repository"
	"github.com/Yakov-Varnaev/ft/internal/repository/category/converter"
	repoModel "github.com/Yakov-Varnaev/ft/internal/repository/category/model"
	pg "github.com/Yakov-Varnaev/ft/pkg/pagination"
	"github.com/Yakov-Varnaev/ft/pkg/repository/utils"
	"github.com/jmoiron/sqlx"
)

var _ def.CategoryRepository = (*repository)(nil)

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *repository {
	return &repository{db}
}

func (r *repository) Exists(filters utils.Filters) (bool, error) {
	q := "SELECT id FROM categories WHERE "
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

const createQuery = `INSERT INTO categories (name, group_id) VALUES ($1, $2) RETURNING id`

func (r *repository) Create(info *model.CategoryInfo) (*model.Category, error) {

	data := converter.ToRepoCategoryInfo(info)
	var id string
	// TODO: May be it worth to create more difficult query with join and returning
	err := r.db.QueryRowx(
		createQuery, data.Name, data.GroupID,
	).Scan(&id)
	if err != nil {
		return nil, err
	}

	var category repoModel.Category

	err = r.db.QueryRowx(
		`
		SELECT categories.id, categories.name, groups.id, groups.name
		FROM categories
		JOIN groups on categories.group_id = groups.id
		WHERE categories.id = $1
		`, id,
	).Scan(&category.ID, &category.Name, &category.Group.ID, &category.Group.Name)
	if err != nil {
		return nil, err
	}

	return converter.FromRepoCategory(&category), nil
}

func (r *repository) List(p pg.Pagination) ([]*model.Category, int, error) {

	var count int
	err := r.db.QueryRowx(
		"SELECT COUNT(*) FROM categories",
	).Scan(&count)
	if err != nil {
		return nil, 0, err
	}

	rows, err := r.db.Queryx(
		`
		SELECT categories.id, categories.name, groups.id, groups.name
		FROM categories
		JOIN groups on categories.group_id = groups.id
		LIMIT $1 OFFSET $2
		`, p.Limit, p.Offset,
	)
	categories := make([]*model.Category, 0)
	for rows.Next() {
		var c repoModel.Category
		// TODO: move scan logic to the model
		err = rows.Scan(&c.ID, &c.Name, &c.Group.ID, &c.Group.Name)
		if err != nil {
			return nil, 0, err
		}
		categories = append(categories, converter.FromRepoCategory(&c))
	}
	return categories, count, nil
}

const deleteQuery = `DELETE FROM categories WHERE id = $1`

func (r *repository) Delete(id string) error {
	_, err := r.db.Exec(deleteQuery, id)
	if err != nil {
		return err
	}
	return nil
}
