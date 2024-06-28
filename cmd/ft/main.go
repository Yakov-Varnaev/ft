package main

import (
	"log/slog"

	db "github.com/Yakov-Varnaev/ft/internal"
	"github.com/Yakov-Varnaev/ft/internal/server"
)

func main() {
	db.Init()

	router := server.SetupRouter()

	if err := router.Run(); err != nil {
		slog.Error("Failed to run server: %v", err)
	}
}
