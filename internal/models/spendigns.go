package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type WriteGroup struct {
	Name string `json:"name,omitempty" db:"name"`
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
	ID       uuid.UUID       `db:"id"`
	Amount   decimal.Decimal `db:"amount"`
	Date     time.Time       `db:"date"`
	Comment  string          `db:"comment"`
	Category uuid.UUID       `db:"category_id"`
}

type WriteSpendings struct {
	Amount   decimal.Decimal `db:"amount" json:"amount,omitempty"`
	Date     time.Time       `db:"date" json:"date,omitempty"`
	Comment  string          `db:"comment" json:"comment,omitempty"`
	Category uuid.UUID       `db:"category_id" json:"category,omitempty"`
}

type DetailedCategory struct {
	ID    uuid.UUID `db:"id" json:"id,omitempty"`
	Name  string    `db:"name" json:"name,omitempty"`
	Group Group     `db:"group" json:"group,omitempty"`
}

type DetailedSpendings struct {
	ID       uuid.UUID        `db:"id" json:"id,omitempty"`
	Amount   decimal.Decimal  `db:"amount" json:"amount,omitempty"`
	Date     time.Time        `db:"date" json:"date,omitempty"`
	Comment  string           `db:"comment" json:"comment,omitempty"`
	Category DetailedCategory `db:"category" json:"category,omitempty"`
}
