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