package main

import (
	"fmt"
	"go-training/application/service"

	"github.com/gin-gonic/gin"

	// userRepo "go-training/infrastructure/rest"
	http "go-training/handlers"
	repository "go-training/infrastructure/InMemory/User"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	userReposit := repository.NewUserRepositoryImpl()

	userService := service.NewUserService(userReposit)

	router := gin.Default()

	moji := "shin"
	hashed, _ := bcrypt.GenerateFromPassword([]byte(moji), 10)
	fmt.Println(string(hashed))
	err := bcrypt.CompareHashAndPassword(hashed, []byte("hehe"))
	if err != nil {
		fmt.Println("間違ってんで")
	}
	errr := bcrypt.CompareHashAndPassword(hashed, []byte("shin"))
	fmt.Println(errr)

	// UserHandlerのインスタンスを作成
	userHandler := http.NewUserHandler(userService)

	// UserHandlerにルーティングを登録
	userHandler.RegisterRoutes(router)

	router.Run(":8080")
}
