package converter

import (
	"github.com/Yakov-Varnaev/ft/internal/model"
	repoModel "github.com/Yakov-Varnaev/ft/internal/repository/spendings/model"
)

func ToRepoSpendingsInfo(s *model.SpendingsInfo) *repoModel.SpendingsInfo {
	return &repoModel.SpendingsInfo{
		Amount:     s.Amount,
		Date:       s.Date,
		Comment:    s.Comment,
		CategoryID: s.CategoryID,
	}
}

func FromRepoSpendings(c *repoModel.Spendings) *model.Spendings {
	return &model.Spendings{
		ID:      c.ID,
		Amount:  c.Amount,
		Date:    c.Date,
		Comment: c.Comment,
		Category: model.Category{
			ID:   c.Category.ID,
			Name: c.Category.Name,
			Group: model.Group{
				ID: c.Category.Group.ID,
				GroupInfo: model.GroupInfo{
					Name: c.Category.Group.Name,
				},
			},
		},
	}
}
