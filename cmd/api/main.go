package main

import (
	"github.com/Yakov-Varnaev/ft/internal/config"
	"github.com/Yakov-Varnaev/ft/internal/database"
	"github.com/Yakov-Varnaev/ft/internal/server"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err.Error())
	}

	db := database.New(cfg.DB)
	defer db.Close()

	app := server.NewServer(db)
	if err = app.Run(); err != nil {
		panic(err.Error())
	}
}
