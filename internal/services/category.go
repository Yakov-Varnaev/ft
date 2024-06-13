package services

import (
	"fmt"
	"log/slog"

	DB "github.com/Yakov-Varnaev/ft/internal"
	"github.com/Yakov-Varnaev/ft/internal/models"
	"github.com/Yakov-Varnaev/ft/internal/web"
	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
)

type Categories struct{}

func (s *Categories) Create(data *models.WriteCategory) (*models.Category, error) {
	if data.Name == "" {
		return nil, &web.ValidationError{Message: "Name cannot be empty"}
	}
	if data.Group == "" {
		return nil, &web.ValidationError{Message: "Group cannot be empty"}
	}
	groupId, err := uuid.Parse(data.Group)
	if err != nil {
		return nil, &web.ValidationError{Message: "Group must be a valid UUID"}
	}
	db := DB.GetDB()
	cnt, err := db.From(DB.CATEGORY_TABLE).Where(goqu.C("name").ILike(data.Name)).Count()
	if err != nil {
		return nil, &web.InternalServerError{Message: err.Error()}
	}
	if cnt > 0 {
		return nil, &web.ValidationError{Message: fmt.Sprintf("Category with name %s already exists.", data.Name)}
	}
	cnt, err = db.From(DB.GROUPS_TABLE).Where(goqu.C("id").Eq(groupId)).Count()
	if err != nil {
		return nil, &web.InternalServerError{Message: err.Error()}
	}
	if cnt == 0 {
		return nil, &web.ValidationError{Message: fmt.Sprintf("Group with id %s does not exist.", groupId)}
	}

	category := &models.Category{
		ID:    uuid.New(),
		Name:  data.Name,
		Group: groupId,
	}
	// create category
	_, err = db.Insert(DB.CATEGORY_TABLE).Rows(category).Executor().Exec()
	// return created cateogry
	return category, nil
}

func (s *Categories) List() (*[]models.Category, error) {
	db := DB.GetDB()
	categories := make([]models.Category, 0)
	err := db.From(DB.CATEGORY_TABLE).ScanStructs(&categories)
	if err != nil {
		slog.Error("Fuck.")
		return nil, &web.InternalServerError{Message: err.Error()}
	}
	return &categories, nil
}
