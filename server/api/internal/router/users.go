package router

import (
	"rest-mcp/internal/auth"

	"github.com/gin-gonic/gin"
)

/*
Module: Users Router intialization
Usage: All the user routes are listed
*/

func registerUserRouter(r *gin.Engine) {
	userRouterGrp := r.Group("/user")

	userRouterGrp.Use(auth.AuthMiddleware()) // user authentication

	userRouterGrp.GET("/:id") // get user details
	userRouterGrp.PUT("/:id") // edit user details
}
