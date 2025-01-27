package wishlist_repository

import (
	"errors"
	"smenard/wishlist-web-service-gin/app/pkg/models"
	"time"
)

// Dummy data
var wishlists = []models.Wishlist{
	{
		ID:            "1",
		Name:          "Hello world",
		OwnerFullName: "John Doe",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	},
}

var wishlistItems = []models.WishlistItem{
	{
		ID:        "1",
		Name:      "Hello world, I want this",
		Status:    "Free",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		ID:        "2",
		Name:      "A big gift for me",
		Status:    "Taken",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}

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
