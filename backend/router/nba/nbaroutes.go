package nba

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sports_api/stats/endpoints/nba"
	static "sports_api/stats/static/nba"
)

// SetupNBARoutes registers NBA-related routes in the Gin engine
func SetupNBARoutes(router *gin.Engine) {
	nbaGroup := router.Group("/nba")
	{
		nbaGroup.GET("/teams", func(c *gin.Context) {

			c.JSON(http.StatusOK, static.GetNBATeamsWithPlayers())
		})
		nbaGroup.GET("/matchups", func(c *gin.Context) {

			c.JSON(http.StatusOK, static.GetNBAMatchups())
		})

		nbaGroup.GET("/players/current", func(c *gin.Context) {
			players := nba.GetAllNBAPlayers()
			if players == nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "no players"})

			}
			c.JSON(http.StatusOK, players)
		})

		// Register the PlayerGameLog route
		nbaGroup.GET("/player/gamelog", func(c *gin.Context) {
			playerID := c.Query("playerID")
			season := c.Query("season")
			seasonType := c.Query("seasonType")

			var leaguePtr *string
			*leaguePtr = "00"

			// Call PlayerGameLog function
			result, err := nba.PlayerGameLog(playerID, season, seasonType, leaguePtr)
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
