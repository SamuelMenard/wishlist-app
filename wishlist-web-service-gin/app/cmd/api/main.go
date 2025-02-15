// main entrypoint of the REST API

package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"smenard/wishlist-web-service-gin/app/pkg/controllers"
	"smenard/wishlist-web-service-gin/app/pkg/database"
	enums "smenard/wishlist-web-service-gin/app/pkg/database/emums"
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

	// Connect to database
	connectToDatabase()

	router.Run("localhost:8080")
}

func connectToDatabase() {
	fmt.Println("connecting to database...")

	config := database.DbConfig{
		Host:     "localhost",
		User:     "admin",
		Password: "root",
		DbName:   "wishlist_db",
		Port:     "5432",
		SslMode:  "disable",
		TimeZone: "GMT",
	}

	db,err := database.ConnectToDb(enums.Postgre, config)

	if (err != nil) {
		panic("failed to connect to db")
	} else {
		fmt.Println("successfully connected to database")
		fmt.Println(db)
	}
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
