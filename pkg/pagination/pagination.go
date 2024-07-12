package pagination

import (
	"fmt"
	"strconv"

	"github.com/Yakov-Varnaev/ft/pkg/web/errors"
	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Limit  int
	Offset int
}

func getStrKeyFromContext(c *gin.Context, key string, def ...string) (string, error) {
	value := c.Query(key)
	if value == "" && len(def) > 0 {
		return def[0], nil
	}
	if value == "" {
		return "", fmt.Errorf("Key %s not found in context", key)
	}
	return value, nil
}

func getIntKeyFromContext(c *gin.Context, key string, def ...string) (int, error) {
	value, err := getStrKeyFromContext(c, key, def...)
	if err != nil {
		return 0, nil
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("Cannot convert '%v' to int.", value)
	}
	return intValue, nil
}

func NewFromContext(c *gin.Context) (Pagination, error) {
	pg := Pagination{}
	offset, err := getIntKeyFromContext(c, "offset", "0")
	if err != nil {
		return pg, &errors.BadRequest{Message: err.Error()}
	}
	limit, err := getIntKeyFromContext(c, "limit", "10")
	if err != nil {
		return pg, &errors.BadRequest{Message: err.Error()}
	}
	pg.Limit = limit
	pg.Offset = offset
	return pg, nil
}

type model interface{}

type Page[m model] struct {
	Total int `json:"total"`
	Data  []m `json:"data"`
}
