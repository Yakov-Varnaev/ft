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
		groupHandler := new(handlers.Groups)
		groupRouter := v1.Group("/groups")
		groupRouter.POST("/", groupHandler.Create)
		groupRouter.GET("/", groupHandler.List)
		groupRouter.PUT("/:id/", groupHandler.Update)
		groupRouter.DELETE("/:id/", groupHandler.Delete)
	}
	{
		categoryHandler := new(handlers.Categories)
		categoryRouter := v1.Group("/categories")
		categoryRouter.GET("/", categoryHandler.List)
		categoryRouter.POST("/", categoryHandler.Create)
		categoryRouter.PUT("/:id/", categoryHandler.Update)
	}
	{
		spendingsHandler := new(handlers.Spendings)
		spendingsRouter := v1.Group("/spendings")
		spendingsRouter.POST("/", spendingsHandler.Create)
		spendingsRouter.GET("/", spendingsHandler.List)
		spendingsRouter.DELETE("/:id/", spendingsHandler.Delete)
	}

	if err := router.Run(); err != nil {
		slog.Error("Failed to run server: %v", err)
	}
}
