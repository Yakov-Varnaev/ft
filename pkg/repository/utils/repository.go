package utils

import (
	"fmt"
	"strings"
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
