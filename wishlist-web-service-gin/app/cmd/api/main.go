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

	// Register routes
	registerControllers(router)

	// Connect to database
	connectToDatabase()

	router.Run("localhost:8080")
}

func registerControllers(router *gin.Engine) {
	fmt.Println("registering controllers")

	router.GET("/wishlist/:id", controllers.GetWishlistById)
	router.GET("/wishlist/items/:wishlistId", controllers.GetWishlistItemsById)
	router.GET("/wishlistItem/:id", controllers.GetWishlistItemById)
	router.GET("/wishlists", controllers.GetWishlists)

	fmt.Println("controllers have been registered")
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
		fmt.Println("failed to connect to database")
		panic("failed to connect to db")
	} else {
		fmt.Println("successfully connected to database")
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
