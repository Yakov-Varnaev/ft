package main

import (
	"fmt"
	"github.com/Yakov-Varnaev/ft/internal/server"
)

func main() {

	server := server.New()

	err := server.Run()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
