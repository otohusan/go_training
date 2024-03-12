package main

import (
	"go-training/application/service"

	"github.com/gin-gonic/gin"

	// userRepo "go-training/infrastructure/rest"
	http "go-training/handlers"
	repository "go-training/infrastructure/InMemory/User"
)

func main() {

	userReposit := repository.NewUserRepositoryImpl()

	userService := service.NewUserService(userReposit)

	router := gin.Default()

	// UserHandlerのインスタンスを作成
	userHandler := http.NewUserHandler(userService)

	// UserHandlerにルーティングを登録
	userHandler.RegisterRoutes(router)

	router.Run(":8080")
}
