package http

import (
	"go-training/application/service"
	"go-training/application/service/auth"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
	authService *auth.AuthService
}

// NewUserHandler はUserHandlerの新しいインスタンスを作成します。
func NewUserHandler(userService *service.UserService, authService *auth.AuthService) *UserHandler {
	return &UserHandler{
		userService: userService,
		authService: authService,
	}
}

// RegisterRoutes はルーターにエンドポイントを登録します。
func (h *UserHandler) RegisterRoutes(router *gin.Engine) {
	router.GET("/users/:id", h.ReturnUser)
	router.GET("/users/", h.GetUserList)
	router.POST("/users/", h.CreateUser)
	router.DELETE("/users/:id", h.DeleteUser)
	router.POST("/auth/", h.Login)
	router.POST("/auth/parse", h.ParseToken)
	router.POST("/auth/register", h.Register)
	router.POST("/post/", h.GetPost)
}
