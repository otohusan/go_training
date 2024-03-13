package http

import (
	"go-training/application/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

// NewUserHandler はUserHandlerの新しいインスタンスを作成します。
func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// RegisterRoutes はルーターにエンドポイントを登録します。
func (h *UserHandler) RegisterRoutes(router *gin.Engine) {
	router.GET("/users/:id", h.ReturnUser)
	router.GET("/users/", h.GetUserList)
	router.POST("/users/", h.CreateUser)
	router.DELETE("/users/:id", h.DeleteUser)
}
