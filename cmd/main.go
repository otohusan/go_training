package main

import (
	"fmt"
	"go-training/application/service"

	// userRepo "go-training/infrastructure/rest"
	repository "go-training/infrastructure/InMemory/User"
)

func main() {

	userReposit := repository.NewUserRepositoryImpl()

	userService := service.NewUserService(userReposit)

	retunr, error := userService.ReturnUser(31)
	if error != nil {
		fmt.Print(error)
		return
	}

	fmt.Println(retunr)
}
