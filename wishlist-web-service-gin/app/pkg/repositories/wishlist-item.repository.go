package repositories

import (
	"errors"
	"slices"
	"smenard/wishlist-web-service-gin/app/pkg/models"
)

/*
* Description: Find a wishlist item by id
* Verb: GET
 */
func GetWishlistItemById(wishlistItemId string) (models.WishlistItem, error) {
	if wishlistItemId == "" {
		return models.WishlistItem{}, errors.New("invalid parameter id")
	}

	idx := slices.IndexFunc(wishlistItems, func(item models.WishlistItem) bool { return item.ID == wishlistItemId })

	if idx == -1 {
		// No item was found, return an error
		return models.WishlistItem{}, errors.New("unable to find wishlit item with the provided id")
	}

	return wishlistItems[idx], nil
}
