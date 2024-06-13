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
	ID   uuid.UUID `json:"id,omitempty"`
	Name string    `json:"name"`
}

type WriteCategory struct {
	Name  string `json:"name,omitempty" sql:"name"`
	Group string `json:"group,omitempty" sql:"group_id"`
}

type Category struct {
	ID    uuid.UUID `db:"id"`
	Name  string    `db:"name"`
	Group uuid.UUID `db:"group_id"`
}

type Spendings struct {
	ID        uuid.UUID
	Amount    big.Float
	Date      time.Time
	Comment   string
	Cateogory uuid.UUID
}

type DetailedCategory struct {
	ID    uuid.UUID
	Name  string
	Group Group
}

type DetailedSpendings struct {
	ID       uuid.UUID
	Amount   big.Float
	Date     time.Time
	Comment  string
	Category DetailedCategory
}
