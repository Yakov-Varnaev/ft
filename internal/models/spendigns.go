package models

import (
	"math/big"
	"time"
)

type Spendings struct {
	Amount    big.Float
	Date      time.Time
	Comment   string
	Cateogory int
}
