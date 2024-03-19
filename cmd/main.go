package main

import (
	"go-training/application/service"
	"go-training/application/service/auth"

	"github.com/gin-gonic/gin"

	// userRepo "go-training/infrastructure/rest"
	http "go-training/handlers"
	repository "go-training/infrastructure/InMemory/User"
)

func main() {

	userReposit := repository.NewUserRepositoryImpl()

	userService := service.NewUserService(userReposit)
	authService := auth.NewAuthService(userReposit)

	router := gin.Default()

	// UserHandlerのインスタンスを作成
	userHandler := http.NewUserHandler(userService, authService)

	// UserHandlerにルーティングを登録
	userHandler.RegisterRoutes(router)

	router.Run(":8080")
}
