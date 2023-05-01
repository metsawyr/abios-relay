package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/metsawyr/abios-api/internal/server/abios"
)

func NewLiveTeamsHandler(client *abios.AbiosClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response, err := client.LiveTeams(ctx)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, response)
	}
}
