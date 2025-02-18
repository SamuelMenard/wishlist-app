package repositories

import (
	"errors"
	"fmt"
	"smenard/wishlist-web-service-gin/app/pkg/models"

	"gorm.io/gorm"
)

type WishlistRepository struct {
	db *gorm.DB
}

func NewWishlistRepository(db *gorm.DB) *WishlistRepository {
	return &WishlistRepository{
		db: db,
	}
}

/*
* Description: Find a wishlist by id
* Verb: GET
 */
func (r *WishlistRepository) GetWishlistById(id string) (models.Wishlist, error) {
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
func (r *WishlistRepository) GetWishlistItemsById(wishlistId string) ([]models.WishlistItem, error) {
	if wishlistId == "" {
		return []models.WishlistItem{}, errors.New("invalid parameter id")
	}

	return wishlistItems, nil
}

/*
* Description: Return user wishlists
* Verb: GET
 */
 func (r *WishlistRepository) GetWishlists() ([]models.Wishlist, error) {
	fmt.Println(r.db)
	return wishlists, nil
}
