package controllers

import (
	"fmt"
	"net/http"
	"smenard/wishlist-web-service-gin/app/pkg/repositories"

	"github.com/gin-gonic/gin"
)

func InitWishlistItemsController(router *gin.Engine) {
	fmt.Println("registering wishlist items controller")

	router.GET("/wishlistItem/:id", getWishlistItemById)

	fmt.Println("wishlist items controller has been registered!")
}

/*
* Description: Find a wishlist item by id
* Verb: GET
 */
func getWishlistItemById(c *gin.Context) {
	id := c.Param("id")

	wishlist,err := repositories.GetWishlistItemById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return 
	}

	c.IndentedJSON(http.StatusOK, wishlist)
}