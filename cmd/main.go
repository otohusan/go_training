package main

import (
	"fmt"
	"go-training/application/service"
	"go-training/domain/model"

	// userRepo "go-training/infrastructure/rest"
	repository "go-training/infrastructure/InMemory/User"
)

func main() {
	user := model.User{
		ID:   213141,
		Name: "sasas",
	}

	userReposit := repository.NewUserRepositoryImpl()

	userService := service.NewUserService(userReposit)

	retunr, error := userService.ReturnUser(user, 2112)
	if error != nil {
		return
	}

	fmt.Println(retunr)
}
