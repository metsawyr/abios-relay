package server

import (
	"github.com/gin-gonic/gin"
	"github.com/metsawyr/abios-api/internal/config"
	"github.com/metsawyr/abios-api/internal/server/abios"
	"github.com/metsawyr/abios-api/internal/server/handlers"
)

type RestServer struct {
	config *config.Config
}

func NewRestServer(config *config.Config) RestServer {
	return RestServer{
		config,
	}
}

func (s *RestServer) Start() {
	router := gin.Default()
	abiosClient := abios.NewAbiosClient(s.config)
	liveSeriesHandler := handlers.NewLiveSeriesHandler(&abiosClient)
	livePlayersHandler := handlers.NewLivePlayersHandler(&abiosClient)
	liveTeamsHandler := handlers.NewLiveTeamsHandler(&abiosClient)

	router.GET("/series/live", liveSeriesHandler)
	router.GET("/players/live", livePlayersHandler)
	router.GET("/teams/live", liveTeamsHandler)

	router.Run(":3000")
}
