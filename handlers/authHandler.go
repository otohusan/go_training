package handlers

import (
	"fmt"
	"go-training/application/service"
	"go-training/utils"
	"log"
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

// Google
func (h *AuthHandler) GoogleLogin(c *gin.Context) {
	var req struct {
		AccessToken string `json:"access_token"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("JSONのバインドに失敗: %v", err)

		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// ユーザーIDを取得または作成
	userID, err := h.authService.CreateOrGetUser(req.AccessToken)
	if err != nil {
		log.Printf("ユーザーの作成に失敗: %v", err)

		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create or get user"})
		return
	}

	// JWTトークンを生成
	token, err := utils.GenerateToken(userID)
	if err != nil {
		log.Printf("トークンの生成に失敗: %v", err)

		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	// トークンを返す
	c.JSON(http.StatusOK, gin.H{"token": token})

}
