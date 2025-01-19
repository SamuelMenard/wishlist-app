package wishlist_repository

import (
	"errors"
	"smenard/wishlist-web-service-gin/app/pkg/models"
	"time"
)

// Dummy data
var wishlists = []models.Wishlist {
	{
		ID: "1",
		Name:      "Hello world",
		Owner:     "1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}

/*
* Description: Find a wishlist by id
* Verb: GET
 */
func GetWishlistById(id string) (models.Wishlist, error){
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