package controllers

import (
	"net/http"
	wishlist_repository "smenard/wishlist-web-service-gin/app/pkg/repositories"

	"github.com/gin-gonic/gin"
)

/*
* Description: Find a wishlist by id
* Verb: GET
 */
func GetWishlistById(c *gin.Context) {
	id := c.Param("id")

	wishlist,err := wishlist_repository.GetWishlistById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return 
	}

	c.IndentedJSON(http.StatusOK, wishlist)
}