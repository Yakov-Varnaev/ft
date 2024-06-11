package main

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	if err := r.Run(); err != nil {
		slog.Error("Failed to run server: %v", err)
	}
}
