package main

import (
	"github.com/metsawyr/abios-api/internal/config"
	"github.com/metsawyr/abios-api/internal/server"
)

func main() {
	config := config.NewConfig()
	server := server.NewRestServer(&config)

	server.Start()
}
