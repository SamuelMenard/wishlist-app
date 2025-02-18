package controllers

import (
	"fmt"
	"smenard/wishlist-web-service-gin/app/pkg/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitControllers(router *gin.Engine, db *gorm.DB) {
	fmt.Println("registering controllers")

	initWishlistsController(router, db)
	initWishlistItemsController(router, db)

	fmt.Println("controllers have been registered")
}

func initWishlistsController(router *gin.Engine, db *gorm.DB) {
	var ctrl = WishlistsController {
		repo: repositories.NewWishlistRepository(db),
	}

	fmt.Println("registering wishlists controller")

	router.GET("/wishlist/:id", ctrl.getWishlistById)
	router.GET("/wishlist/items/:wishlistId", ctrl.getWishlistItemsById)
	router.GET("/wishlists", ctrl.getWishlists)

	fmt.Println("wishlists controller has been registered!")
}

func initWishlistItemsController(router *gin.Engine, db *gorm.DB) {
	var ctrl = WishlistItemsController {
		repo: repositories.NewWishlistItemsRepository(db),
	}

	fmt.Println("registering wishlist items controller")

	router.GET("/wishlistItem/:id", ctrl.getWishlistItemById)

	fmt.Println("wishlist items controller has been registered!")
}