package main

import (
	"go-training/application/service"
	"go-training/application/service/auth"

	// userRepo "go-training/infrastructure/rest"
	http "go-training/handlers"
	repository "go-training/infrastructure/InMemory/User"
)

func main() {

	userReposit := repository.NewUserRepository()

	userService := service.NewUserService(userReposit)
	authService := auth.NewAuthService(userReposit)

	router := http.SetupRoutes(userService, authService, authService)

	router.Run(":8080")
}
