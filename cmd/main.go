package main

import (
	"go-training/application/service"
	"go-training/handlers"

	"github.com/gin-gonic/gin"

	// userRepo "go-training/infrastructure/rest"
	repository "go-training/infrastructure/InMemory/User"
)

func main() {

	// リポジトリの初期化
	userRepo := repository.NewUserRepository()

	// サービスの初期化
	userService := service.NewUserService(userRepo)

	// ハンドラーの初期化
	userHandler := handlers.NewUserHandler(userService)

	// Ginのルーターを設定
	router := gin.Default()

	// ルートの設定
	setupUserRoutes(router, userHandler)

	router.Run(":8080")
}

func setupUserRoutes(router *gin.Engine, userHandler *handlers.UserHandler) {
	// ユーザー関連のルートをグループ化
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/", userHandler.CreateUser)
		userRoutes.GET("/:id", userHandler.GetUserByID)
		userRoutes.GET("/username/:username", userHandler.GetUserByUsername)
		userRoutes.PUT("/:id", userHandler.UpdateUser)
		userRoutes.DELETE("/:id", userHandler.DeleteUser)
	}
}
