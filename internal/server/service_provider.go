package server

import (
	repository "github.com/Yakov-Varnaev/ft/internal/repository/group"
	groupService "github.com/Yakov-Varnaev/ft/internal/service/group"
	"github.com/jmoiron/sqlx"
)

type handlerProvider struct {
	groupHandler *groupHandler
}

func newServiceProvider(db *sqlx.DB) *handlerProvider {
	groupRepo := repository.New(db)
	groupService := groupService.New(groupRepo)
	groupHandler := NewGroupHandler(groupService)
	return &handlerProvider{
		groupHandler: groupHandler,
	}
}
