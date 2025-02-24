package router

import (
	"github.com/gin-gonic/gin"
	router "sports_api/router/nba"
)

// SetupRouter initializes the main router and includes league-specific routes
func SetupRouter() *gin.Engine {
	r := gin.Default()

	router.SetupNBARoutes(r)
	router.SetupWNBARoutes(r)

	return r
}
