package main

import (
	"fmt"
	"go-training/application/service"
	"go-training/domain/model"

	// userRepo "go-training/infrastructure/rest"
	repository "go-training/infrastructure/InMemory/User"
)

func main() {

	userReposit := repository.NewUserRepositoryImpl()

	userService := service.NewUserService(userReposit)

	// retunr, error := userService.ReturnUser(31)
	// if error != nil {
	// 	fmt.Print(error)
	// 	return
	// }

	user := model.User{Name: "saii", ID: 2}

	userService.Create(&user)

	userList, eerror := userService.GetUserList()
	if eerror != nil {
		fmt.Print(eerror)
		return
	}

	fmt.Println(userList)
}
