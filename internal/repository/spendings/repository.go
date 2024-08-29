package repository

import (
	"fmt"
	"math/big"

	"github.com/Yakov-Varnaev/ft/internal/model"
	def "github.com/Yakov-Varnaev/ft/internal/repository"
	"github.com/Yakov-Varnaev/ft/internal/repository/spendings/converter"
	repoModel "github.com/Yakov-Varnaev/ft/internal/repository/spendings/model"
	pg "github.com/Yakov-Varnaev/ft/pkg/pagination"
	"github.com/Yakov-Varnaev/ft/pkg/repository/utils"
	"github.com/jmoiron/sqlx"
)

var _ def.SpendingsRepository = (*repository)(nil)

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *repository {
	return &repository{db}
}

func (r *repository) Exists(filters utils.Filters) (bool, error) {
	q := "SELECT id FROM spendings WHERE "
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

const createQuery = `
INSERT INTO spendings (amount, date, comment, category_id)
VALUES ($1, $2, $3, $4) RETURNING id
`

func (r *repository) Create(info *model.SpendingsInfo) (*model.Spendings, error) {
	data := converter.ToRepoSpendingsInfo(info)
	var id string
	amount := data.Amount.Int64()
	// TODO: May be it worth to create more difficult query with join and returning
	err := r.db.QueryRowx(
		createQuery, amount, data.Date, data.Comment, data.CategoryID,
	).Scan(&id)
	if err != nil {
		return nil, err
	}

	var spendings repoModel.Spendings
	var amountres int

	err = r.db.QueryRowx(
		`
		SELECT (
			spendings.id,
			spendings.amount,
			spendings.date,
			spendings.comment,
			groups.id,
			groups.name,
			categories.id,
			categories.name
		)
		FROM spendings
		JOIN categories on categories.id = spendings.category_id
		JOIN groups on categories.group_id = groups.id
		WHERE spendings.id = $1
		`, id,
	).Scan(
		&spendings.ID,
		&amountres,
		// &spendings.Amount,
		&spendings.Date,
		&spendings.Comment,
		&spendings.Category.Group.ID,
		&spendings.Category.Group.Name,
		&spendings.Category.ID,
		&spendings.Category.Name,
	)
	if err != nil {
		return nil, err
	}

	return converter.FromRepoSpendings(&spendings), nil
}

func (r *repository) List(p pg.Pagination) ([]*model.Spendings, int, error) {
	var count int
	err := r.db.QueryRowx(
		"SELECT COUNT(*) FROM spendings",
	).Scan(&count)
	if err != nil {
		return nil, 0, err
	}

	rows, err := r.db.Queryx(
		`
		SELECT spendings.id, spendings.amount, spendings.date, spendings.comment, groups.id, groups.name, categories.id, categories.name
		FROM spendings
		JOIN categories on categories.id = spendings.category_id
		JOIN groups on categories.group_id = groups.id
		LIMIT $1 OFFSET $2
		`, p.Limit, p.Offset,
	)
	spendings := make([]*model.Spendings, 0)
	var amount float64
	if err != nil {
		return nil, 0, err
	}
	for rows.Next() {
		var s repoModel.Spendings
		err = rows.Scan(
			&s.ID,
			&amount,
			&s.Date,
			&s.Comment,
			&s.Category.Group.ID,
			&s.Category.Group.Name,
			&s.Category.ID,
			&s.Category.Name,
		)
		// TODO: fix this shit...
		s.Amount = big.NewInt(int64(amount))
		if err != nil {
			return nil, 0, err
		}
		spendings = append(spendings, converter.FromRepoSpendings(&s))
	}
	return spendings, count, nil
}

// const updateQuery = `
// 	UPDATE categories SET name = $1, group_id = $2 WHERE id = $3 RETURNING id
// `
//
// func (r *repository) Update(id string, info *model.SpendingsInfo) (*model.Spendings, error) {
// 	data := converter.ToRepoSpendingsInfo(info)
// 	err := r.db.QueryRowx(
// 		updateQuery, data.Name, data.GroupID, id,
// 	).Scan(&id)
//
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	var category repoModel.Category
//
// 	err = r.db.QueryRowx(
// 		`
// 		SELECT categories.id, categories.name, groups.id, groups.name
// 		FROM categories
// 		JOIN groups on categories.group_id = groups.id
// 		WHERE categories.id = $1
// 		`, id,
// 	).Scan(&category.ID, &category.Name, &category.Group.ID, &category.Group.Name)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return converter.FromRepoCategory(&category), nil
// }
//
// const deleteQuery = `DELETE FROM categories WHERE id = $1`
//
// func (r *repository) Delete(id string) error {
// 	_, err := r.db.Exec(deleteQuery, id)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
