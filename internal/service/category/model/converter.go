package model

import repoModel "github.com/Yakov-Varnaev/ft/internal/repository/category/model"

func ToRepoCategoryInfo(cat *CategoryInfo) *repoModel.CategoryInfo {
	return &repoModel.CategoryInfo{
		Name:    cat.Name,
		GroupId: cat.GroupId,
	}
}

func FromRepoCategoryInfo(cat *repoModel.CategoryInfo) *CategoryInfo {
	return &CategoryInfo{
		Name:    cat.Name,
		GroupId: cat.GroupId,
	}
}

func ToRepoCategory(cat *Category) *repoModel.Category {
	return &repoModel.Category{
		UUID: cat.UUID,
		CategoryInfo: repoModel.CategoryInfo{
			Name:    cat.Name,
			GroupId: cat.GroupId,
		},
	}
}

func FromRepoCategory(cat *repoModel.Category) *Category {
	return &Category{
		UUID: cat.UUID,
		CategoryInfo: CategoryInfo{
			Name:    cat.Name,
			GroupId: cat.GroupId,
		},
	}
}
