package http

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"go-training/domain/model"
)

var mySigningKey = []byte("secret") // 実際の環境では安全にキーを管理してください。

// CreateToken はJWTトークンを生成します。
func CreateToken(id string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	return tokenString, err
}

// 新規登録
func (h *UserHandler) Register(c *gin.Context) {
	var creds model.CreatedUserData
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "認証情報のバインドに失敗しました"})
		return
	}
	h.authService.RegisterUser(creds)
}

func (h *UserHandler) Login(c *gin.Context) {
	var creds model.UserCredentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "認証情報のバインドに失敗しました"})
		return
	}

	if creds.ID == "" || creds.Password == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "情報が足りません"})
		return
	}

	//　ユーザ情報の取得
	userinfo, err := h.userService.ReturnUser(creds.ID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "ユーザーが存在しません"})
		return
	}

	// パスワードの照合
	if userinfo.Password != creds.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "パスワードが適切ではありません"})
		return
	}

	tokenString, err := CreateToken(creds.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "トークンの生成に失敗しました"})
		return
	}

	c.JSON(http.StatusOK, tokenString)
}

type UserToken struct {
	UserToken string `json:"usertoken"`
}

func (h *UserHandler) ParseToken(c *gin.Context) {
	var utoken UserToken
	if err := c.ShouldBindJSON(&utoken); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "認証情報のバインドに失敗しました"})
		return
	}

	token, err := jwt.Parse(utoken.UserToken, func(token *jwt.Token) (interface{}, error) {
		// 署名アルゴリズムの検証
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "認証に失敗しました"})
		}
		return mySigningKey, nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "認証に失敗しました"})
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id := claims["id"].(string)
		// ユーザーを返す
		userinfo, err := h.userService.ReturnUser(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		// パスワードを含まない形に変換
		res := model.UserResponse{
			Name: userinfo.Name,
			ID:   userinfo.ID,
		}

		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "認証に失敗しました"})
	}

}
