package converter

import (
	"github.com/Yakov-Varnaev/ft/internal/model"
	repoModel "github.com/Yakov-Varnaev/ft/internal/repository/group/model"
)

func ToRepoGroupInfo(g *model.GroupInfo) *repoModel.GroupInfo {
	return &repoModel.GroupInfo{
		Name: g.Name,
	}
}

func FromRepoGroup(g *repoModel.Group) *model.Group {
	return &model.Group{
		ID: g.ID,
		GroupInfo: model.GroupInfo{
			Name: g.Name,
		},
	}
}
