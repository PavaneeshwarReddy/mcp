package router

import "github.com/gin-gonic/gin"

/*
Module: Router Intialization
Usage: Main addition of all routes and middlewares from different modules
*/

func New() *gin.Engine {
	r := gin.New()

	registerAuthRouters(r)   // auth route group, /auth/...
	registerUserRouter(r)    // user route group, /user/...
	registerOrderRoutes(r)   // order router group, /orders/...
	registerPaymentRoutes(r) // payment router group, /payments/...

	return r
}
