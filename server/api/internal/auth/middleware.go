package auth

import "github.com/gin-gonic/gin"

/*
Module: Authentication Middleware
Usage: Used for verification of jwt token
*/

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()

	}
}
