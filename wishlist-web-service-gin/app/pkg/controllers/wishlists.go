package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
* Description: Find a wishlist by id
* Verb: GET
 */
func GetWishlistById(c *gin.Context) {
	id := c.Param("id")
	c.IndentedJSON(http.StatusOK, id)
}