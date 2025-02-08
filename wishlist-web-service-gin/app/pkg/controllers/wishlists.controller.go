package controllers

import (
	"net/http"
	"smenard/wishlist-web-service-gin/app/pkg/repositories"

	"github.com/gin-gonic/gin"
)

/*
* Description: Find a wishlist by id
* Verb: GET
 */
func GetWishlistById(c *gin.Context) {
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
 func GetWishlistItemsById(c *gin.Context) {
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
 func GetWishlists(c *gin.Context) {
	wishlists,err := repositories.GetWishlists()

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return 
	}

	c.IndentedJSON(http.StatusOK, wishlists)
}