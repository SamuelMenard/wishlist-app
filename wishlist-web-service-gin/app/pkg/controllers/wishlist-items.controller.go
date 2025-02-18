package controllers

import (
	"net/http"
	"smenard/wishlist-web-service-gin/app/pkg/repositories"

	"github.com/gin-gonic/gin"
)

type WishlistItemsController struct {
	repo *repositories.WishlistItemsRepository
}

/*
* Description: Find a wishlist item by id
* Verb: GET
 */
func (ctrl *WishlistItemsController) getWishlistItemById(c *gin.Context) {
	id := c.Param("id")

	wishlist,err := ctrl.repo.GetWishlistItemById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return 
	}

	c.IndentedJSON(http.StatusOK, wishlist)
}