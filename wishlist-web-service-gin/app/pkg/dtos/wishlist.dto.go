package dtos

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string    `gorm:"uniqueIndex"`
	LastName  string    `gorm:"uniqueIndex"`
	Email     string    `gorm:"not null"`
	Username  string    `gorm:"not null"`
	Country   string    `gorm:"not null"`
	Role      string    `gorm:"not null"`
	Age       int       `gorm:"not null;size:3"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt
}

type Wishlist struct {
	gorm.Model
	Name      string    `gorm:"uniqueIndex"`
	Owner     string    `gorm:"uniqueIndex"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt
}
