package controllers

import (
	"fmt"
	"net/http"
	"smenard/wishlist-web-service-gin/app/pkg/repositories"

	"github.com/gin-gonic/gin"
)

func InitWishlistsController(router *gin.Engine) {
	fmt.Println("registering wishlists controller")

	router.GET("/wishlist/:id", getWishlistById)
	router.GET("/wishlist/items/:wishlistId", getWishlistItemsById)
	router.GET("/wishlists", getWishlists)

	fmt.Println("wishlists controller has been registered!")
}

/*
* Description: Find a wishlist by id
* Verb: GET
 */
func getWishlistById(c *gin.Context) {
	id := c.Param("id")

	wishlist,err := repositories.GetWishlistById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return 
	}

	c.IndentedJSON(http.StatusOK, wishlist)
}

/*
* Description: Find a wishlist by id
* Verb: GET
 */
 func getWishlistItemsById(c *gin.Context) {
	id := c.Param("wishlistId")

	wishlistItems,err := repositories.GetWishlistItemsById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return 
	}

	c.IndentedJSON(http.StatusOK, wishlistItems)
}

/*
* Description: Return user wishlists
* Verb: GET
 */
 func getWishlists(c *gin.Context) {
	wishlists,err := repositories.GetWishlists()

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return 
	}

	c.IndentedJSON(http.StatusOK, wishlists)
}