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

func getCategoryById(id uuid.UUID) (*models.Category, error) {
	return DB.GetById[models.Category](DB.CATEGORY_TABLE, id)
}

// TODO: use Categories.DB instead of GetDB
type Categories struct{}

func (s *Categories) validateGroup(id string) error {
	if id == "" {
		return &web.ValidationError{Message: "Group cannot be empty"}
	}
	groupId, err := uuid.Parse(id)
	if err != nil {
		return &web.ValidationError{Message: "Group must be a valid UUID"}
	}
	cnt, err := DB.GetDB().From(DB.GROUPS_TABLE).Where(goqu.C("id").Eq(groupId)).Count()
	if err != nil {
		return &web.InternalServerError{Message: err.Error()}
	}
	if cnt == 0 {
		return &web.ValidationError{Message: fmt.Sprintf("Group with id %s does not exist.", groupId)}
	}
	return nil
}

func (s *Categories) validateName(name string) error {
	if name == "" {
		return &web.ValidationError{Message: "Name cannot be empty"}
	}
	cnt, err := DB.GetDB().From(DB.CATEGORY_TABLE).Where(goqu.C("name").ILike(name)).Count()
	if err != nil {
		return &web.InternalServerError{Message: err.Error()}
	}
	if cnt > 0 {
		return &web.ValidationError{Message: fmt.Sprintf("Category with name %s already exists.", name)}
	}
	return nil
}

func (s *Categories) Create(data *models.WriteCategory) (*models.Category, error) {

	err := s.validateGroup(data.Group)
	if err != nil {
		return nil, err
	}

	category := &models.Category{
		ID:    uuid.New(),
		Name:  data.Name,
		Group: data.Group,
	}
	// create category
	_, err = DB.GetDB().Insert(DB.CATEGORY_TABLE).Rows(category).Executor().Exec()
	// return created cateogry
	return category, nil
}

func (s *Categories) List() (*[]models.DetailedCategory, error) {
	categories := make([]models.DetailedCategory, 0)
	err := DB.GetDB().From(DB.CATEGORY_TABLE).
		Select(
			goqu.I("categories.id").As(goqu.C("id")),
			goqu.I("categories.name").As(goqu.C("name")),
			goqu.I("group.id").As(goqu.C("group.id")),
			goqu.I("group.name").As(goqu.C("group.name")),
		).
		Join(goqu.T(DB.GROUPS_TABLE).As("group"), goqu.On(goqu.I("group_id").Eq(goqu.C("group.id")))).
		ScanStructs(&categories)
	if err != nil {
		slog.Error(err.Error())
		return nil, &web.InternalServerError{Message: err.Error()}
	}
	return &categories, nil
}

func (s *Categories) Update(id uuid.UUID, data *models.WriteCategory) (*models.Category, error) {
	category, err := getCategoryById(id)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, &web.NotFound{Message: "Category with given id not found"}
	}
	slog.Info("Category was found.")

	err = s.validateName(data.Name)
	if err != nil {
		return nil, err
	}
	slog.Info("Name is Valid")

	err = s.validateGroup(data.Group)
	if err != nil {
		return nil, err
	}
	slog.Info("Group is Valid")

	_, err = DB.GetDB().Update(DB.CATEGORY_TABLE).Set(data).Where(goqu.C("id").Eq(id)).Executor().Exec()
	if err != nil {
		return nil, &web.InternalServerError{Message: err.Error()}
	}
	slog.Info("Category updated successfully.")

	updatedCategory, err := getCategoryById(id)
	if err != nil {
		return nil, &web.InternalServerError{Message: err.Error()}
	}
	if updatedCategory == nil {
		panic("this shouldn't be right")
	}
	slog.Info("Return updated category")

	return updatedCategory, nil
}
