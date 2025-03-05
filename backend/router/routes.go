package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"sports_api/router/mlb"
	"sports_api/router/nba"
	"sports_api/router/nhl"
	"time"
)

// SetupRouter initializes the main router and includes league-specific routes with CORS enabled
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Configure CORS settings
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Adjust to specific domains if needed
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour, // Cache the preflight request for 12 hours
	}))

	// Setup routes
	nba.SetupNBARoutes(r)
	nba.SetupWNBARoutes(r)
	mlb.SetupMLBRoutes(r)
	nhl.SetupNHLRoutes(r)
	return r
}
