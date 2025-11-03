// gin-rest-api/api/routes.go
package api

import (
	"gin-rest-api/handlers"
	"gin-rest-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Public routes
	router.GET("/", handlers.Home)
	router.POST("/api/register", handlers.Register)
	router.POST("/api/login", handlers.Login)

	// Protected routes (require JWT)
	protected := router.Group("/api")
	protected.Use(middleware.JWTAuth())
	{
		protected.GET("/profile", handlers.GetProfile)
		protected.PUT("/profile", handlers.UpdateProfile)
		protected.GET("/users", handlers.GetAllUsers)
	}

	// Admin routes
	admin := protected.Group("/admin")
	admin.Use(middleware.AdminOnly())
	{
		admin.GET("/stats", handlers.GetStats)
	}
}
