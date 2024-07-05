package utils

import (
	"fmt"
	"strings"

	"github.com/Yakov-Varnaev/ft/pkg/pagination"
	"github.com/jmoiron/sqlx"
)

type Filters map[string]interface{}

func (f Filters) Prepare() (string, []interface{}) {
	strFilters := []string{}
	idx := 1
	params := []interface{}{}
	for k, v := range f {
		strFilters = append(strFilters, fmt.Sprintf("%s = $%d", k, idx))
		idx++
		params = append(params, v)
	}
	var q string
	if len(strFilters) > 1 {
		q = strings.Join(strFilters, " AND ")
	} else {
		q = strFilters[0]
	}
	return q, params
}

type Scanner interface {
	Scan(dest ...any) error
}

type DBModel[T any] interface {
	FromRow(src Scanner) (*T, error)
}

func List[dbModel DBModel[dbModel]](
	db *sqlx.DB, table, query string, pg pagination.Pagination,
) ([]*dbModel, int, error) {
	var count int
	err := db.QueryRowx(
		// TODO: there can be some sort of filtration
		"SELECT COUNT(*) FROM " + table,
	).Scan(&count)
	if err != nil {
		return nil, 0, err
	}
	rows, err := db.Queryx(
		query+" LIMIT $1 OFFSET $2",
		pg.Limit, pg.Offset,
	)
	if err != nil {
		return nil, 0, err
	}
	objects := make([]*dbModel, 0)
	for rows.Next() {
		var obj dbModel
		ptr, err := obj.FromRow(rows)
		if err != nil {
			return nil, 0, err
		}
		objects = append(objects, ptr)
	}
	return objects, count, nil

}
