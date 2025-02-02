package repositories

import (
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