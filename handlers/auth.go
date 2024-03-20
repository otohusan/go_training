package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var mySigningKey = []byte("secret") // 実際の環境では安全にキーを管理してください。

// UserCredentials はリクエストからユーザー情報を受け取るための構造体です。
type UserCredentials struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// CreateToken はJWTトークンを生成します。
func CreateToken(id string, username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["id"] = id
	claims["user"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	return tokenString, err
}

func (h *UserHandler) Login(c *gin.Context) {
	var creds UserCredentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "認証情報のバインドに失敗しました"})
		return
	}

	if creds.Username == "" || creds.ID == "" || creds.Password == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "情報が足りません"})
		return
	}

	// ユーザー名とパスワードの検証（ここでは仮の検証を行っています）
	if creds.Username != "user" || creds.Password != "password" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "認証情報が無効です"})
		return
	}

	tokenString, err := CreateToken(creds.ID, creds.Username)
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
		fmt.Println(claims["user"], claims["exp"])
		id := claims["id"].(string)
		// ユーザーを返す
		userinfo, err := h.userService.ReturnUser(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		c.JSON(http.StatusOK, userinfo)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "認証に失敗しました"})
	}

}
