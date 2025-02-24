package nba

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sports_api/stats/endpoints/nba"
	static "sports_api/stats/static/nba"
)

// SetupNBARoutes registers NBA-related routes in the Gin engine
func SetupWNBARoutes(router *gin.Engine) {
	wnbaGroup := router.Group("/wnba")
	{
		wnbaGroup.GET("/teams", func(c *gin.Context) {

			c.JSON(http.StatusOK, static.GetWNBATeams())
		})

		wnbaGroup.GET("/players/current", func(c *gin.Context) {
			players := nba.GetAllWNBAPlayers()
			if players == nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "no players"})

			}
			c.JSON(http.StatusOK, players)
		})

		// Register the PlayerGameLog route
		wnbaGroup.GET("/player/gamelog", func(c *gin.Context) {
			playerID := c.Query("playerID")
			season := c.Query("season")
			seasonType := c.Query("seasonType")

			leagueID := "10"

			// Call PlayerGameLog function
			result, err := nba.PlayerGameLog(playerID, season, seasonType, &leagueID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			dict, err := result.GetNormalizedDict2()
			if err != nil {
				return
			}

			c.JSON(http.StatusOK, dict)
		})

	}
}
