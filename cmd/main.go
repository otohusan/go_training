package main

import (
	"go-training/application/service"
	"go-training/domain/model"
	userRepo "go-training/infrastructure/rest"
)

func main() {
	user := model.User{
		ID:   213141,
		Name: "sasas",
	}

	userReposit := userRepo.NewUserRepositoryImpl()

	userService := service.NewUserService(userReposit)

	retunr, error := userService.FindByID(user, 2112)
	if error != nil {
		return
	}

	kore := userService.Test("koreyone")

	dom := userService.Dom("sasa")

	println(retunr, kore, dom)
}
