package controllers

import (
	"net/http"
	"smenard/wishlist-web-service-gin/app/pkg/repositories"

	"github.com/gin-gonic/gin"
)

type WishlistsController struct {
	repo *repositories.WishlistRepository
}

/*
* Description: Find a wishlist by id
* Verb: GET
*/
 func (ctrl *WishlistsController) getWishlistById(c *gin.Context) {
	id := c.Param("id")

	wishlist,err := ctrl.repo.GetWishlistById(id)

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
func (ctrl *WishlistsController) getWishlistItemsById(c *gin.Context) {
	id := c.Param("wishlistId")

	wishlistItems,err := ctrl.repo.GetWishlistItemsById(id)

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
func (ctrl *WishlistsController) getWishlists(c *gin.Context) {
	wishlists,err := ctrl.repo.GetWishlists()

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return 
	}

	c.IndentedJSON(http.StatusOK, wishlists)
}