// main entrypoint of the REST API

package main

import (
	"github.com/gin-gonic/gin"

	"smenard/wishlist-web-service-gin/app/pkg/controllers"
	"smenard/wishlist-web-service-gin/app/pkg/database"
)

type User struct {
	Id 		  uint8
	Username string
}

func main() {
	router := gin.Default()

	router.Use(corsMiddleware())

	// Initialize controllers and register routes
	controllers.InitControllers(router)

	// Connect db
	database.AutoMigrateDatabase()

	router.Run("localhost:8080")
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
