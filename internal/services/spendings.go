package services

import (
	db "github.com/Yakov-Varnaev/ft/internal"
	"github.com/Yakov-Varnaev/ft/internal/models"
)

func GetSpendings() (*[]*models.DetailedSpendings, error) {
	conn := db.GetDB()
	spendings := make([]*models.DetailedSpendings, 0)

	conn.From(db.SPENDINGS_TABLE).ScanStructs(spendings)

	return &spendings, nil
}
