package controllers

import (
	"net/http"
	"smenard/wishlist-web-service-gin/app/pkg/repositories"

	"github.com/gin-gonic/gin"
)

/*
* Description: Find a wishlist item by id
* Verb: GET
 */
func GetWishlistItemById(c *gin.Context) {
	id := c.Param("id")

	wishlist,err := repositories.GetWishlistItemById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return 
	}

	c.IndentedJSON(http.StatusOK, wishlist)
}