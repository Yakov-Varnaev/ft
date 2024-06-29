package pagination

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Model interface{}

type Page[T Model] struct {
	Data  []T `json:"data"`
	Total int `json:"total"`
}

type Pagination struct {
	Limit  int
	Offset int
}

func getStrParam(c *gin.Context, key string, def ...string) (string, error) {
	value := c.Query(key)
	if value == "" {
		if len(def) > 0 {
			return def[0], nil
		} else {
			return "", errors.New("param not found")
		}
	}
	return value, nil
}

func getIntParam(c *gin.Context, key string, def ...string) (int, error) {
	value, err := getStrParam(c, key, def...)
	if err != nil {
		return 0, err
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}
	return intValue, nil
}

func NewFromRequest(c *gin.Context) (Pagination, error) {
	pg := Pagination{}
	var err error
	pg.Limit, err = getIntParam(c, "limit", "10")
	if err != nil {
		return pg, err
	}

	pg.Offset, err = getIntParam(c, "offset", "0")
	if err != nil {
		return pg, err
	}
	return pg, nil
}
