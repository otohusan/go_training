package middleware

import (
	"go-training/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddlewareはJWTトークンを解析し、コンテキストにuserIDを設定するミドルウェアです
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		userID, err := utils.ParseToken(authHeader)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("AuthUserID", userID)
		c.Next()
	}
}
