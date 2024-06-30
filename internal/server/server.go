package server

import (
	"github.com/Yakov-Varnaev/ft/internal/config"
	"github.com/Yakov-Varnaev/ft/internal/database"
	groupHandler "github.com/Yakov-Varnaev/ft/internal/handlers/group/handler"
	groupRepository "github.com/Yakov-Varnaev/ft/internal/repository/group"
	groupService "github.com/Yakov-Varnaev/ft/internal/service/group/service"
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	cfg, err := config.New()
	if err != nil {
		panic(err.Error())
	}
	db := database.New(cfg.Database)
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), ErrorHandler)

	apiGroup := r.Group("api")

	groupRepo := groupRepository.New(db)
	groupService := groupService.New(groupRepo)
	groupHandler := groupHandler.New(groupService)
	groupHandler.RegisterRoutes(apiGroup)

	return r
}
