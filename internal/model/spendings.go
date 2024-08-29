package model

import (
	"math/big"
	"time"
)

type SpendingsInfo struct {
	Amount     *big.Int  `json:"amount,omitempty" validate:"required"`
	Date       time.Time `json:"date,omitempty" validate:"required"`
	Comment    string    `json:"comment,omitempty"`
	CategoryID string    `json:"category_id,omitempty" validate:"required,uuid4"`
}

type Spendings struct {
	ID       string    `json:"id,omitempty"`
	Amount   *big.Int  `json:"amount,omitempty"`
	Date     time.Time `json:"date,omitempty"`
	Comment  string    `json:"comment,omitempty"`
	Category Category  `json:"category,omitempty"`
}
