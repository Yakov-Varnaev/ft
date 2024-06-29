package model

import repoModel "github.com/Yakov-Varnaev/ft/internal/repository/group/model"

func ToRepoGroupInfo(g *GroupInfo) *repoModel.GroupInfo {
	return &repoModel.GroupInfo{
		Name: g.Name,
	}
}

func FromRepoGroupInfo(g *repoModel.GroupInfo) *GroupInfo {
	return &GroupInfo{
		Name: g.Name,
	}
}

func ToRepoGroup(g *Group) *repoModel.Group {
	return &repoModel.Group{
		UUID: g.UUID,
		GroupInfo: repoModel.GroupInfo{
			Name: g.Name,
		},
	}
}

func FromRepoGroup(g *repoModel.Group) *Group {
	return &Group{
		UUID: g.UUID,
		GroupInfo: GroupInfo{
			Name: g.Name,
		},
	}
}
