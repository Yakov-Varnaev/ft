package converter

import (
	"github.com/Yakov-Varnaev/ft/internal/model"
	repoModel "github.com/Yakov-Varnaev/ft/internal/repository/category/model"
)

func ToRepoCategoryInfo(g *model.CategoryInfo) *repoModel.CategoryInfo {
	return &repoModel.CategoryInfo{
		Name:    g.Name,
		GroupID: g.GroupID,
	}
}

func FromRepoCategory(c *repoModel.Category) *model.Category {
	return &model.Category{
		ID:   c.ID,
		Name: c.Name,
		Group: model.Group{
			ID: c.Group.ID,
			GroupInfo: model.GroupInfo{
				Name: c.Group.Name,
			},
		},
	}
}
