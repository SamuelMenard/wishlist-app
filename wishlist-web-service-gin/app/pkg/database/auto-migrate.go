package database

import (
	"fmt"
	"smenard/wishlist-web-service-gin/app/pkg/database/dtos"
)

func AutoMigrateDatabase() {
	var err error
	db, err := NewPostgresDatabaseConnection().OpenConnection()

	if err != nil {
		fmt.Println("unable to connect to database")
	}

	// Create types
	db.Exec("CREATE TYPE wishlist_role_type AS ENUM ('owner', 'contributor');")

	// Create tables
	db.Migrator().CreateTable(&dtos.User{}, &dtos.Wishlist{}, &dtos.UsersWishlists{})

	err = db.SetupJoinTable(&dtos.User{}, "Wishlists", &dtos.UsersWishlists{})

	if err != nil {
		fmt.Println("unable to create users_wishlists table")
	}

	err = db.SetupJoinTable(&dtos.Wishlist{}, "Users", &dtos.UsersWishlists{})

	if err != nil {
		fmt.Println("unable to create users_wishlists table")
	}
}