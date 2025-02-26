package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	router "sports_api/router/nba"
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
	router.SetupNBARoutes(r)
	router.SetupWNBARoutes(r)

	return r
}
