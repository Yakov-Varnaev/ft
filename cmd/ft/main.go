package main

import (
	"log/slog"

	db "github.com/Yakov-Varnaev/ft/internal"
	"github.com/Yakov-Varnaev/ft/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	apiRouter := router.Group("/api")
	v1 := apiRouter.Group("/v1")
	{
		groupRouter := v1.Group("/groups")
		groupHandler := new(handlers.Groups)
		groupRouter.POST("/", groupHandler.Create)
		groupRouter.GET("/", groupHandler.List)
		groupRouter.PUT("/:id/", groupHandler.Update)
		groupRouter.DELETE("/:id/", groupHandler.Delete)
	}
	{
		categoryRouter := v1.Group("/categories")
		categoryHandler := new(handlers.Categories)
		categoryRouter.GET("/", categoryHandler.List)
		categoryRouter.POST("/", categoryHandler.Create)
	}

	if err := router.Run(); err != nil {
		slog.Error("Failed to run server: %v", err)
	}
}
