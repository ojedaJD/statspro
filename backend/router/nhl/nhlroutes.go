package nhl

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sports_api/stats/endpoints/nhl"
	static "sports_api/stats/static/nhl"
)

// SetupNBARoutes registers NBA-related routes in the Gin engine
func SetupNHLRoutes(router *gin.Engine) {
	nbaGroup := router.Group("/nhl")
	{
		nbaGroup.GET("/teams", func(c *gin.Context) {
			teams, err := nhl.GetAndParseNHLTeams()
			if err != nil || len(teams) == 0 {
				c.JSON(http.StatusInternalServerError, gin.H{"error getting nhl teams": err})
			}
			c.JSON(http.StatusOK, teams)
		})
		nbaGroup.GET("/matchups", func(c *gin.Context) {
			matchups := static.GetNHLMatchupsWithOdds()
			if len(matchups) == 0 || matchups == nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error getting nhl teams": "No Matchups found"})
			}
			c.JSON(http.StatusOK, matchups)
		})
	}
}
