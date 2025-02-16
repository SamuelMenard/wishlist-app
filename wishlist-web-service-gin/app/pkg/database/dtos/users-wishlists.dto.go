package dtos

import (
	"smenard/wishlist-web-service-gin/app/pkg/database/types"
	"time"

	"gorm.io/gorm"
)

type UsersWishlists struct {
	UserID  int `gorm:"primaryKey"`
  	WishlistID int `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt
	UserRole types.UserRole `gorm:"type:wishlist_role_type"`
}