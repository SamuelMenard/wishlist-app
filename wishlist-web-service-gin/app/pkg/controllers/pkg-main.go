package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitControllers(router *gin.Engine) {
	fmt.Println("registering controllers")

	InitWishlistsController(router)
	InitWishlistItemsController(router)

	fmt.Println("controllers have been registered")
}