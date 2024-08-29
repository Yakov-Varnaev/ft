package model

import (
	"math/big"
	"time"
)

type Group struct {
	ID   string
	Name string
}

type Category struct {
	ID    string
	Name  string
	Group Group
}

type SpendingsInfo struct {
	Amount     *big.Int
	Date       time.Time
	Comment    string
	CategoryID string
}

type Spendings struct {
	ID       string
	Amount   *big.Int
	Date     time.Time
	Comment  string
	Category Category
}
