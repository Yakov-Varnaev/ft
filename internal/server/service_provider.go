package server

import (
	categoryRepository "github.com/Yakov-Varnaev/ft/internal/repository/category"
	groupRepository "github.com/Yakov-Varnaev/ft/internal/repository/group"
	spendingsRepository "github.com/Yakov-Varnaev/ft/internal/repository/spendings"
	categoryService "github.com/Yakov-Varnaev/ft/internal/service/category"
	groupService "github.com/Yakov-Varnaev/ft/internal/service/group"
	spendingsService "github.com/Yakov-Varnaev/ft/internal/service/spendings"
	"github.com/jmoiron/sqlx"
)

type handlerProvider struct {
	groupHandler     *groupHandler
	categoryHandler  *categoryHandler
	spendingsHandler *spendingsHandler
}

func newServiceProvider(db *sqlx.DB) *handlerProvider {
	groupRepo := groupRepository.New(db)
	groupService := groupService.New(groupRepo)
	groupHandler := NewGroupHandler(groupService)

	categoryRepo := categoryRepository.New(db)
	categoryService := categoryService.New(categoryRepo)
	categoryHandler := NewCategoryHandler(categoryService)

	spendingsRepo := spendingsRepository.New(db)
	spendingsServcie := spendingsService.New(spendingsRepo)
	spendingsHandler := NewSpendingsHandler(spendingsServcie)

	return &handlerProvider{
		groupHandler:     groupHandler,
		categoryHandler:  categoryHandler,
		spendingsHandler: spendingsHandler,
	}
}
