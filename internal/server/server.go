package server

import (
	"github.com/Yakov-Varnaev/ft/internal/config"
	"github.com/Yakov-Varnaev/ft/internal/database"
	catHandler "github.com/Yakov-Varnaev/ft/internal/handlers/categories"
	groupHandler "github.com/Yakov-Varnaev/ft/internal/handlers/group"
	catRepo "github.com/Yakov-Varnaev/ft/internal/repository/category"
	groupRepository "github.com/Yakov-Varnaev/ft/internal/repository/group"
	catService "github.com/Yakov-Varnaev/ft/internal/service/category"
	groupService "github.com/Yakov-Varnaev/ft/internal/service/group"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	db     *sqlx.DB
	Engine *gin.Engine
}

func New(cfg config.Config) *Server {
	db := database.New(cfg.Database)
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), ErrorHandler)

	apiGroup := r.Group("api")

	groupRepo := groupRepository.New(db)
	groupService := groupService.New(groupRepo)
	groupHandler := groupHandler.New(groupService)
	groupHandler.RegisterRoutes(apiGroup)

	catRepo := catRepo.New(db)
	catService := catService.New(catRepo, groupRepo)
	catHandler := catHandler.New(catService)
	catHandler.RegisterRoutes(apiGroup)

	return &Server{
		db:     db,
		Engine: r,
	}
}

func (s *Server) Run() error {
	return s.Engine.Run()
}

func (s *Server) Close() error {
	return s.db.Close()
}
