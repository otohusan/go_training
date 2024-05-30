package utils

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// JWTトークンを署名するために使用する秘密鍵
var jwtSecret = []byte("secret")

// ParseTokenはJWTトークンを解析し、userIDを返します
func ParseToken(authHeader string) (string, error) {
	if authHeader == "" {
		return "", errors.New("authHeaderが空です")
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 署名方法が期待するものかを検証
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("予期しない署名方法です")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["userID"].(string)
		if !ok {
			return "", errors.New("トークンにuserIDが含まれていません")
		}
		return userID, nil
	}

	return "", errors.New("無効なトークンです")
}
