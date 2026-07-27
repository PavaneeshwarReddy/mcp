package router

import (
	"rest-mcp/internal/app"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/*
Module: Router Intialization
Usage: Main addition of all routes and middlewares from different modules
*/

func New(db *gorm.DB, app *app.Dependencies) *gin.Engine {
	r := gin.New()

	registerAuthRouters(r)                 // auth route group, /auth/...
	registerUserRouter(r, app.UserHandler) // user route group, /user/...
	registerOrderRoutes(r)                 // order router group, /orders/...
	registerPaymentRoutes(r)               // payment router group, /payments/...

	return r
}
