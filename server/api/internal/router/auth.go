package router

import "github.com/gin-gonic/gin"

/*
Module: Auth Router intialization
Usage: All the authentication routes are listed
*/

func registerAuthRouters(r *gin.Engine) {
	authRouterGrp := r.Group("/auth")
	authRouterGrp.POST("/login")    // user login, which returns jwt token
	authRouterGrp.POST("/register") // create new user
	authRouterGrp.POST("/logout")   // mark the token as invalid
	authRouterGrp.POST("/refresh")  // refresh new token, if access token is invalid
}
