package main

import (
	"go-training/application/service"
	"go-training/handlers"

	"github.com/gin-gonic/gin"

	// userRepo "go-training/infrastructure/rest"
	studySet "go-training/infrastructure/InMemory/StudySet"
	user "go-training/infrastructure/InMemory/User"
)

func main() {

	// リポジトリの初期化
	userRepo := user.NewUserRepository()
	studySetRepo := studySet.NewStudySetRepository()

	// サービスの初期化
	userService := service.NewUserService(userRepo)
	studySetService := service.NewStudySetService(studySetRepo)

	// ハンドラーの初期化
	userHandler := handlers.NewUserHandler(userService)
	studySetHandler := handlers.NewStudySetHandler(studySetService)

	// Ginのルーターを設定
	router := gin.Default()

	// ルートの設定
	setupRoutes(router, userHandler, studySetHandler)

	// サーバーの起動
	router.Run(":8080")
}

func setupRoutes(router *gin.Engine, userHandler *handlers.UserHandler, studySetHandler *handlers.StudySetHandler) {
	// ユーザー関連のルートをグループ化
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/", userHandler.GetAllUsers)
		userRoutes.POST("/", userHandler.CreateUser)
		userRoutes.GET("/:id", userHandler.GetUserByID)
		userRoutes.GET("/username/:username", userHandler.GetUserByUsername)
		userRoutes.PUT("/:id", userHandler.UpdateUser)
		userRoutes.DELETE("/:id", userHandler.DeleteUser)
	}

	studySetRoutes := router.Group("/studysets")
	{
		studySetRoutes.POST("/", studySetHandler.CreateStudySet)
		studySetRoutes.GET("/:id", studySetHandler.GetStudySetByID)
		studySetRoutes.GET("/user/:userID", studySetHandler.GetStudySetsByUserID)
		studySetRoutes.PUT("/:id", studySetHandler.UpdateStudySet)
		studySetRoutes.DELETE("/:id", studySetHandler.DeleteStudySet)
	}
}
