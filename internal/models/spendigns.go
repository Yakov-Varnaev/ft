package models

import (
	"math/big"
	"time"

	"github.com/google/uuid"
)

type WriteGroup struct {
	Name string `json:"name"`
}

type Group struct {
	ID   uuid.UUID `json:"id,omitempty" db:"id"`
	Name string    `json:"name,omitempty" db:"name"`
}

type WriteCategory struct {
	Name  string `json:"name,omitempty" db:"name"`
	Group string `json:"group,omitempty" db:"group_id"`
}

type Category struct {
	ID    uuid.UUID `db:"id" json:"id,omitempty"`
	Name  string    `db:"name" json:"name,omitempty"`
	Group string    `db:"group_id" json:"group,omitempty"`
}

type Spendings struct {
	ID        uuid.UUID
	Amount    big.Float
	Date      time.Time
	Comment   string
	Cateogory uuid.UUID
}

type DetailedCategory struct {
	ID    uuid.UUID `db:"id" json:"id,omitempty"`
	Name  string    `db:"name" json:"name,omitempty"`
	Group Group     `db:"group" json:"group,omitempty"`
}

type DetailedSpendings struct {
	ID       uuid.UUID
	Amount   big.Float
	Date     time.Time
	Comment  string
	Category DetailedCategory
}
