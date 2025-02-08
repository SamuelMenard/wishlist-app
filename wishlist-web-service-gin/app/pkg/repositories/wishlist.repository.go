package repositories

import (
	"errors"
	"smenard/wishlist-web-service-gin/app/pkg/models"
)

/*
* Description: Find a wishlist by id
* Verb: GET
 */
func GetWishlistById(id string) (models.Wishlist, error) {
	if id == "" {
		return models.Wishlist{}, errors.New("invalid parameter id")
	}

	for _, wishlist := range wishlists {
		if wishlist.ID == id {
			return wishlist, nil
		}
	}

	return models.Wishlist{}, errors.New("unable to find wishlist")
}

/*
* Description: Find a wishlist items by id
* Verb: GET
 */
func GetWishlistItemsById(wishlistId string) ([]models.WishlistItem, error) {
	if wishlistId == "" {
		return []models.WishlistItem{}, errors.New("invalid parameter id")
	}

	return wishlistItems, nil
}

/*
* Description: Return user wishlists
* Verb: GET
 */
 func GetWishlists() ([]models.Wishlist, error) {
	return wishlists, nil
}
