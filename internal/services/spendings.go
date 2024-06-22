package services

import (
	"log/slog"

	DB "github.com/Yakov-Varnaev/ft/internal"
	"github.com/Yakov-Varnaev/ft/internal/models"
	"github.com/Yakov-Varnaev/ft/internal/web"
	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type detailedQueryProcessor struct{}

func (p *detailedQueryProcessor) Process(query *goqu.SelectDataset) *goqu.SelectDataset {
	return query.
		Select(
			goqu.I("spendings.id").As(goqu.C("id")),
			goqu.I("spendings.amount").As(goqu.C("amount")),
			goqu.I("spendings.date").As(goqu.C("date")),
			goqu.I("spendings.comment").As(goqu.C("comment")),
			// category
			goqu.I("category.id").As(goqu.C("category.id")),
			goqu.I("category.name").As(goqu.C("category.name")),
			// group
			goqu.I("group.id").As(goqu.C("category.group.id")),
			goqu.I("group.name").As(goqu.C("category.group.name")),
		).
		Join(goqu.T(DB.CATEGORY_TABLE).As("category"), goqu.On(goqu.I("category.id").Eq(goqu.C("category_id")))).
		Join(goqu.T(DB.GROUPS_TABLE).As("group"), goqu.On(goqu.I("group.id").Eq(goqu.C("group_id"))))
}

type Spendigns struct{}

func (s *Spendigns) Create(data *models.WriteSpendings) (*models.DetailedSpendings, error) {
	if data.Amount == decimal.NewFromInt(0) {
		return nil, &web.ValidationError{Message: "Amount cannot be 0"}
	}
	exists, err := DB.Exists(DB.CATEGORY_TABLE, &idFilter{data.Category})
	if err != nil {
		return nil, &web.ValidationError{Message: "Fail to query category: " + err.Error()}
	}
	if !exists {
		return nil, &web.ValidationError{Message: "Category does not exists"}
	}
	slog.Info("category exists this is fine")
	newSpending := models.Spendings{
		ID:       uuid.New(),
		Comment:  data.Comment,
		Date:     data.Date,
		Amount:   data.Amount,
		Category: data.Category,
	}
	_, err = DB.Create[models.Spendings, models.Spendings](DB.SPENDINGS_TABLE, &newSpending)
	if err != nil {
		return nil, err
	}
	slog.Info("New spend created")
	spend, err := DB.GetById[models.DetailedSpendings](DB.SPENDINGS_TABLE, newSpending.ID, &detailedQueryProcessor{})
	if err != nil {
		return nil, err
	}
	slog.Info("Retrieved new spend")
	return spend, nil
}

func (s *Spendigns) List() (*[]models.DetailedSpendings, error) {
	// TODO: add pagination and date filtering
	spendings, err := DB.List[models.DetailedSpendings](DB.SPENDINGS_TABLE, &detailedQueryProcessor{})
	if err != nil {
		return nil, err
	}
	return spendings, nil
}

type idFilter struct {
	ID uuid.UUID
}

func (p *idFilter) Process(query *goqu.SelectDataset) *goqu.SelectDataset {
	return query.Where(goqu.C("id").Eq(p.ID))
}

func (s *Spendigns) Update(id string, data *models.WriteSpendings) (*models.DetailedSpendings, error) {
	// check if spending exists
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	spending, err := DB.GetById[models.Spendings](DB.SPENDINGS_TABLE, uid, nil)
	if err != nil {
		return nil, err
	}
	if spending == nil {
		return nil, &web.NotFound{Message: "Spending not found"}
	}
	// check if new category exists
	if spending.Category != data.Category {
		newCategory, err := DB.GetById[models.Category](DB.CATEGORY_TABLE, data.Category, nil)
		if err != nil {
			return nil, err
		}
		if newCategory == nil {
			return nil, &web.NotFound{Message: "Category does not exist"}
		}
	}
	// check if new data is valid
	if data.Amount == decimal.NewFromInt(0) {
		return nil, &web.ValidationError{Message: "Amount cannot be null"}
	}
	// update
	return nil, nil
}

func (s *Spendigns) Delete(id string) (*uuid.UUID, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	deletedID, err := DB.DeleteById(DB.SPENDINGS_TABLE, uid)
	if err != nil {
		return nil, err
	}
	return deletedID, nil
}
