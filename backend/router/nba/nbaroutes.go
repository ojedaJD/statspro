package nba

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sports_api/stats/endpoints/nba"
	static "sports_api/stats/static/nba"
	"strconv"
)

// SetupNBARoutes registers NBA-related routes in the Gin engine
func SetupNBARoutes(router *gin.Engine) {
	nbaGroup := router.Group("/nba")
	{
		nbaGroup.GET("/teams", func(c *gin.Context) {

			c.JSON(http.StatusOK, static.GetNBATeams())
		})

		nbaGroup.GET("/players", func(c *gin.Context) {
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
			leagueID := c.Query("leagueID")

			var leaguePtr *string
			if leagueID != "" {
				leaguePtr = &leagueID

			}

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

		nbaGroup.GET("/players/all", func(c *gin.Context) {
			isOnlyCurrentSeasonStr := c.DefaultQuery("isOnlyCurrentSeason", "1") // Default to current season players
			// Default to NBA
			season := c.Query("season")

			// Convert isOnlyCurrentSeason to int
			isOnlyCurrentSeason, err := strconv.Atoi(isOnlyCurrentSeasonStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value for isOnlyCurrentSeason, must be 0 or 1"})
				return
			}

			// Call CommonAllPlayers function
			result, err := nba.CommonAllPlayers(isOnlyCurrentSeason, "00", season)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			nba.GetAllNBAPlayers()
			// Convert response to a normalized structure
			dict, err := result.GetNormalizedDict()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, dict)
		})
	}
}
