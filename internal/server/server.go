package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

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
	s.engine.Use(errorHandler)
	apiGroup := s.engine.Group("/api")
	{
		v1 := apiGroup.Group("/v1")
		{
			groupGroup := v1.Group("/groups")
			groupGroup.GET("", s.h.groupHandler.List)
			groupGroup.POST("", s.h.groupHandler.Create)
			groupGroup.PUT("/:id", s.h.groupHandler.Update)
			groupGroup.DELETE("/:id", s.h.groupHandler.Delete)
		}
		{
			categoryGroup := v1.Group("/categories")
			categoryGroup.GET("", s.h.categoryHandler.List)
			categoryGroup.POST("", s.h.categoryHandler.Create)
			categoryGroup.PUT("/:id", s.h.categoryHandler.Update)
			categoryGroup.DELETE("/:id", s.h.categoryHandler.Delete)
		}
		{
			spendingsGroup := v1.Group("/spendings")
			spendingsGroup.POST("", s.h.spendingsHandler.Create)
			spendingsGroup.GET("", s.h.spendingsHandler.List)
		}
	}
}
