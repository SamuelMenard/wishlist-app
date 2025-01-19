// main entrypoint of the REST API

package main

import (
	"github.com/gin-gonic/gin"

	"smenard/wishlist-web-service-gin/app/pkg/controllers"
)

func main() {
    router := gin.Default()

    // Register routes
    registerControllers(router)

    router.Run("localhost:8080")
}

func registerControllers(router *gin.Engine) {
    router.GET("/wishlist/:id", controllers.GetWishlistById)
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