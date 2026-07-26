package router

import (
	"rest-mcp/internal/auth"

	"github.com/gin-gonic/gin"
)

/*
Module: Payment Router intialization
Usage: All the payment routes are listed
*/

func registerPaymentRoutes(r *gin.Engine) {
	paymentRouterGrp := r.Group("/payments")

	paymentRouterGrp.Use(auth.AuthMiddleware()) // user authentication

	paymentRouterGrp.POST("")       // create new payments
	paymentRouterGrp.GET("/:id")    // get payment details
	paymentRouterGrp.DELETE("/:id") // delete payment
}
