package server

import (
	categoryRepository "github.com/Yakov-Varnaev/ft/internal/repository/category"
	groupRepository "github.com/Yakov-Varnaev/ft/internal/repository/group"
	categoryService "github.com/Yakov-Varnaev/ft/internal/service/category"
	groupService "github.com/Yakov-Varnaev/ft/internal/service/group"
	"github.com/jmoiron/sqlx"
)

type handlerProvider struct {
	groupHandler    *groupHandler
	categoryHandler *categoryHandler
}

func newServiceProvider(db *sqlx.DB) *handlerProvider {
	groupRepo := groupRepository.New(db)
	groupService := groupService.New(groupRepo)
	groupHandler := NewGroupHandler(groupService)

	categoryRepo := categoryRepository.New(db)
	categoryService := categoryService.New(categoryRepo)
	categoryHandler := NewCategoryHandler(categoryService)

	return &handlerProvider{
		groupHandler:    groupHandler,
		categoryHandler: categoryHandler,
	}
}
