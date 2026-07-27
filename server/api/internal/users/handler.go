package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc *Service
}

func NewHandler(service *Service) Handler {
	return Handler{
		svc: service,
	}
}

func (h *Handler) Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req RegisterUserRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
}
