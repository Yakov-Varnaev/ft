package utils

import (
	"reflect"
	"testing"
)

func TestFilters_Prepare(t *testing.T) {
	tests := []struct {
		name  string
		f     Filters
		want  string
		want1 []interface{}
	}{
		{
			name:  "One filter",
			f:     Filters{"key": "value"},
			want:  "key = $1",
			want1: []interface{}{"value"},
		},
		{
			name:  "Two filters",
			f:     Filters{"key1": "value1", "key2": 1},
			want:  "key1 = $1 AND key2 = $2",
			want1: []interface{}{"value1", 1},
		},
		{
			name:  "Three filters",
			f:     Filters{"key1": "value1", "key2": 1, "key3": true},
			want:  "key1 = $1 AND key2 = $2 AND key3 = $3",
			want1: []interface{}{"value1", 1, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.f.Prepare()
			if got != tt.want {
				t.Errorf("Filters.Prepare() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Filters.Prepare() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
