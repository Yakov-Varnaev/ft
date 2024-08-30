package server

import (
	"github.com/Yakov-Varnaev/ft/internal/handlers"
	groupRepository "github.com/Yakov-Varnaev/ft/internal/repository/group"
	groupService "github.com/Yakov-Varnaev/ft/internal/service/group"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type handlerProvider struct {
	groupHandler *handlers.GroupHandler
}

func newServiceProvider(db *sqlx.DB) *handlerProvider {
	groupRepo := groupRepository.New(db)
	groupService := groupService.New(groupRepo)
	groupHandler := handlers.NewGroupHandler(groupService)

	return &handlerProvider{
		groupHandler: groupHandler,
	}
}

type Server struct {
	h      *handlerProvider
	engine *gin.Engine
}

func NewServer(db *sqlx.DB) *Server {
	serviceProvider := newServiceProvider(db)
	return &Server{h: serviceProvider, engine: gin.Default()}
}

func (s *Server) Run() error {
	s.RegisterRoutes()
	return s.engine.Run()
}

func (s *Server) RegisterRoutes() {
	s.engine.Use(handlers.ErrorHandler)
	apiGroup := s.engine.Group("/api")
	{
		v1 := apiGroup.Group("/v1")
		{
			groupGroup := v1.Group("/groups")
			groupGroup.GET("", s.h.groupHandler.List)
			groupGroup.POST("", s.h.groupHandler.Create)
		}
	}
}
