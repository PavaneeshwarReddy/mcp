package router

import (
	"rest-mcp/internal/users"

	"github.com/gin-gonic/gin"
)

/*
Module: Users Router intialization
Usage: All the user routes are listed
*/

func registerUserRouter(r *gin.Engine, hdlr *users.Handler) {
	userRouterGrp := r.Group("/user")

	userRouterGrp.POST("/register", hdlr.Register()) // create new user

	userRouterGrp.GET("/:id") // get user details
	userRouterGrp.PUT("/:id") // edit user details
}
