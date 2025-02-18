// main entrypoint for the CLI tool

package main

import "smenard/wishlist-web-service-gin/app/pkg/database"

func main() {
	database.AutoMigrateDatabase()
}