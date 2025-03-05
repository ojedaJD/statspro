package mlb

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sports_api/stats/endpoints/mlb"
)

// SetupNBARoutes registers NBA-related routes in the Gin engine
func SetupMLBRoutes(router *gin.Engine) {
	nbaGroup := router.Group("/mlb")
	{
		nbaGroup.GET("/teams", func(c *gin.Context) {

			c.JSON(http.StatusOK, mlb.GetAndParseMLBTeams())
		})
	}
}
