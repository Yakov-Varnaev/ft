package main

import (
	"fmt"

	"github.com/Yakov-Varnaev/ft/internal/config"
	"github.com/Yakov-Varnaev/ft/internal/server"
)

func main() {

	config, err := config.New()
	if err != nil {
		panic(err.Error())
	}

	server := server.New(config)
	defer server.Close()

	if err = server.Run(); err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
