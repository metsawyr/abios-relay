package server

import (
	"github.com/gin-gonic/gin"
	"github.com/metsawyr/abios-api/internal/config"
	"github.com/metsawyr/abios-api/internal/server/abios"
	"github.com/metsawyr/abios-api/internal/server/handlers"
	"github.com/redis/go-redis/v9"
)

type RestServer struct {
	config      *config.Config
	redisClient *redis.Client
}

func NewRestServer(config *config.Config, redisClient *redis.Client) RestServer {
	return RestServer{
		config,
		redisClient,
	}
}

func (s *RestServer) Start() {
	router := gin.Default()
	abiosClient := abios.NewAbiosClient(s.config)
	liveSeriesHandler := handlers.NewLiveSeriesHandler(&abiosClient)
	livePlayersHandler := handlers.NewLivePlayersHandler(&abiosClient)
	liveTeamsHandler := handlers.NewLiveTeamsHandler(&abiosClient)

	rateLimiter := newRateLimiter(s.config, s.redisClient)

	router.GET("/series/live", rateLimiter, liveSeriesHandler)
	router.GET("/players/live", rateLimiter, livePlayersHandler)
	router.GET("/teams/live", rateLimiter, liveTeamsHandler)

	router.Run(":8000")
}
