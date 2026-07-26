package router

import (
	"rest-mcp/internal/auth"

	"github.com/gin-gonic/gin"
)

/*
Module: Orders Router intialization
Usage: All the order routes are listed
*/

func registerOrderRoutes(r *gin.Engine) {
	orderRouterGrp := r.Group("/orders")

	orderRouterGrp.Use(auth.AuthMiddleware()) // verify authentication

	orderRouterGrp.POST("")    // create a new order
	orderRouterGrp.GET("/:id") // get order details
	orderRouterGrp.PUT("/:id") // update an order
}
