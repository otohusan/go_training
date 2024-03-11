package http

import (
	"go-training/application/service"
	"net/http"

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
	router.GET("/users", h.GetUserList)
}

// GetUserList はユーザーリストを取得するためのハンドラーです。
func (h *UserHandler) GetUserList(c *gin.Context) {
	userList, err := h.userService.GetUserList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, userList)
}
