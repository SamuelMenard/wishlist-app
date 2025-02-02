// main entrypoint of the REST API

package main

import (
	"github.com/gin-gonic/gin"

	"smenard/wishlist-web-service-gin/app/pkg/controllers"
)

func main() {
	router := gin.Default()

	router.Use(corsMiddleware())

	// Register routes
	registerControllers(router)

	router.Run("localhost:8080")
}

func registerControllers(router *gin.Engine) {
	router.GET("/wishlist/:id", controllers.GetWishlistById)
	router.GET("/wishlist/items/:wishlistId", controllers.GetWishlistItemsById)
	router.GET("/wishlistItem/:id", controllers.GetWishlistItemById)
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

// func connectDatabase() {
//     //Create a new Postgresql database connection
//     dsn := "host=localhost user=<your_user> password=<your_password> dbname=<your_dbname> port=<your_port>"

//     // Open a connection to the database
//     db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
//     if err != nil {
//         panic("failed to connect to database: " + err.Error())
//     }
// }
