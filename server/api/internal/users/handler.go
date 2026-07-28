package users

import (
	"net/http"
	"rest-mcp/internal/shared"

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
			ctx.JSON(http.StatusBadRequest, shared.ErrorResponse("invalid request", err.Error()))
			return
		}
		if err := h.svc.CreateUser(req); err != nil {
			ctx.JSON(http.StatusInternalServerError, shared.ErrorResponse("internal server error", err.Error()))
			return
		}

		ctx.JSON(http.StatusCreated, shared.SuccessResponse("user created succssfully", RegisterUserResponse{}))
	}
}
