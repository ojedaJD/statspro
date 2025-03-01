package nba

import (
	"github.com/gin-gonic/gin"
	"net/http"
	endpoints "sports_api/stats/endpoints/nba"
	static "sports_api/stats/static/nba"
	"strconv"
)

// SetupNBARoutes registers NBA-related routes in the Gin engine
func SetupNBARoutes(router *gin.Engine) {
	nbaGroup := router.Group("/nba")
	{
		nbaGroup.GET("/teams", func(c *gin.Context) {

			c.JSON(http.StatusOK, static.GetNBATeamsWithPlayers())
		})
		nbaGroup.GET("v1/matchups", func(c *gin.Context) {
			c.JSON(http.StatusOK, static.GetNBAMatchups())
		})
		nbaGroup.GET("v2/matchups", func(c *gin.Context) {
			c.JSON(http.StatusOK, static.GetNBAMatchupsWithOdds())
		})
		nbaGroup.GET("v2/player/gamelogs", func(c *gin.Context) {
			playerIDStr := c.Query("playerID")
			period := c.Query("period") // Expected values: "Season", "1H", "2H", "q3", "q4", "first_half", "second_half"

			// Validate period parameter
			if period == "" {
				period = "Season" // Default to full season stats if not provided
			}
			playerID, err := strconv.Atoi(playerIDStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid playerID, must be an integer"})
				return
			}

			gameLogs := endpoints.GetCurrentSeasonStats(period).GetPlayerGameLog(playerID)

			if gameLogs == nil || len(gameLogs) == 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "No Game Logs found for the given period"})
				return
			}

			c.JSON(http.StatusOK, gameLogs)
		})

		nbaGroup.GET("/matchups/players", func(c *gin.Context) {

			c.JSON(http.StatusOK, static.GetActivePlayerForToday())
		})

		nbaGroup.GET("/players/current", func(c *gin.Context) {
			players := endpoints.GetAllNBAPlayers()
			if players == nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "no players"})

			}
			c.JSON(http.StatusOK, players)
		})

		// Register the PlayerGameLog route
		nbaGroup.GET("v1/player/gamelog", func(c *gin.Context) {
			playerID := c.Query("playerID")
			season := c.Query("season")
			seasonType := c.Query("seasonType")

			league := "00"

			gamelogs := endpoints.GetPlayerGameLog(playerID, season, seasonType, &league)

			if gamelogs == nil || len(gamelogs) == 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "No Game Logs"})
				return
			}

			c.JSON(http.StatusOK, gamelogs)
		})

		nbaGroup.GET("v2/player/hit-rate", func(c *gin.Context) {
			playerIDStr := c.Query("playerID")
			stat := c.Query("stat")              // e.g., "PTS", "REB", "AST", "FG3M"
			numGamesStr := c.Query("num")        // Number of games to calculate hit rate
			thresholdStr := c.Query("threshold") // Minimum stat value to count as a "hit"

			playerID, err := strconv.Atoi(playerIDStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid playerID, must be an integer"})
				return
			}
			numGames, err := strconv.Atoi(numGamesStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid numGames, must be an integer"})
				return
			}
			threshold, err := strconv.ParseFloat(thresholdStr, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid threshold, must be an integer"})
				return
			}

			gameLogs := endpoints.GetCurrentSeasonStats("full").GetPlayerGameLog(playerID)
			hitRate := CalculateHitRate(gameLogs, stat, numGames, threshold)

			c.JSON(http.StatusOK, gin.H{
				"playerID":  playerID,
				"stat":      stat,
				"numGames":  numGames,
				"threshold": threshold,
				"hitRate":   hitRate,
			})
		})

		nbaGroup.GET("v2/player/streak", func(c *gin.Context) {
			playerIDStr := c.Query("playerID")
			stat := c.Query("stat")              // e.g., "PTS", "REB", "AST", "FG3M"
			thresholdStr := c.Query("threshold") // Minimum stat value to count as a "hit"

			playerID, err := strconv.Atoi(playerIDStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid playerID, must be an integer"})
				return
			}
			threshold, err := strconv.ParseFloat(thresholdStr, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid threshold, must be an integer"})
				return
			}

			gameLogs := endpoints.GetCurrentSeasonStats("full").GetPlayerGameLog(playerID)
			streak := CalculateStreak(gameLogs, stat, threshold)

			c.JSON(http.StatusOK, gin.H{
				"playerID":  playerID,
				"stat":      stat,
				"threshold": threshold,
				"streak":    streak,
			})
		})

	}
}
