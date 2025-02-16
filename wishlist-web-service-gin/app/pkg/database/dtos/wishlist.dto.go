package dtos

import (
	"time"

	"gorm.io/gorm"
)

type Wishlist struct {
	gorm.Model
	Name      string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt
	Users []User `gorm:"many2many:users_wishlists;"`
}
