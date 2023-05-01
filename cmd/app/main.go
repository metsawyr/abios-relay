package main

import (
	"github.com/metsawyr/abios-api/internal/config"
	"github.com/metsawyr/abios-api/internal/resources"
	"github.com/metsawyr/abios-api/internal/server"
)

func main() {
	config := config.NewConfig()
	redisClient := resources.NewRedisClient(config)
	server := server.NewRestServer(config, redisClient)

	server.Start()
}
