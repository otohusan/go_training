package handlers

import (
	"fmt"
	"go-training/application/service"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) RegisterWithEmail(c *gin.Context) {
	// 受け取るデータ構造を定義、受け取り
	var registrationData struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&registrationData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// ユーザを仮登録
	message, err := h.authService.RegisterWithEmail(
		registrationData.Username,
		registrationData.Email,
		registrationData.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": message})
}

func (h *AuthHandler) VerifyEmail(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	message, err := h.authService.VerifyEmail(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: リダイレクト先が作成できれば、そこに行ってもらう
	FRONTEND_URL := os.Getenv("FRONTEND_URL")
	// リダイレクト先を設定
	redirectURL := fmt.Sprintf("%s/Login", FRONTEND_URL)

	// リダイレクト
	c.Redirect(http.StatusFound, redirectURL)

	c.JSON(http.StatusOK, gin.H{"message": message})
}
