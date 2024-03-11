package main

import (
	"fmt"
	"go-training/application/service"

	"github.com/gin-gonic/gin"

	// userRepo "go-training/infrastructure/rest"
	http "go-training/handlers"
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

	// userService.DeleteUserByID(39)

	userList, eerror := userService.GetUserList()
	if eerror != nil {
		fmt.Print(eerror)
		return
	}

	fmt.Println(userList)

	router := gin.Default()

	// UserHandlerのインスタンスを作成
	userHandler := http.NewUserHandler(userService)

	// UserHandlerにルーティングを登録
	userHandler.RegisterRoutes(router)

	router.Run(":8080")
}
