// gin-rest-api/main.go
package main

import (
	"fmt"
	"gin-rest-api/api"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set Gin mode
	gin.SetMode(gin.ReleaseMode)
	if os.Getenv("DEBUG") == "true" {
		gin.SetMode(gin.DebugMode)
	}

	// Create Gin router
	router := gin.Default()

	// Setup routes
	api.SetupRoutes(router)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := fmt.Sprintf(":%s", port)
	log.Printf("Server starting on http://localhost%s", addr)
	log.Printf("Available endpoints:")
	log.Printf("  POST /api/register    - Register new user")
	log.Printf("  POST /api/login       - Login and get JWT")
	log.Printf("  GET  /api/profile     - Get user profile (requires JWT)")
	log.Printf("  PUT  /api/profile     - Update profile (requires JWT)")
	log.Printf("  GET  /api/users       - List all users (requires JWT)")

	if err := router.Run(addr); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
