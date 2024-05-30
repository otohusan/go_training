package main

import (
	"go-training/application/service"
	"go-training/handlers"

	"github.com/gin-gonic/gin"

	// userRepo "go-training/infrastructure/rest"
	flashcard "go-training/infrastructure/InMemory/FlashCard"
	studySet "go-training/infrastructure/InMemory/StudySet"
	user "go-training/infrastructure/InMemory/User"
)

func main() {

	// リポジトリの初期化
	userRepo := user.NewUserRepository()
	studySetRepo := studySet.NewStudySetRepository()
	flashcardRepo := flashcard.NewFlashcardRepository()

	// サービスの初期化
	userService := service.NewUserService(userRepo)
	studySetService := service.NewStudySetService(studySetRepo)
	flashcardService := service.NewFlashcardService(flashcardRepo)

	// ハンドラーの初期化
	userHandler := handlers.NewUserHandler(userService)
	studySetHandler := handlers.NewStudySetHandler(studySetService)
	flashcardHandler := handlers.NewFlashcardHandler(flashcardService)

	// Ginのルーターを設定
	router := gin.Default()

	// ルートの設定
	setupRoutes(router, userHandler, studySetHandler, flashcardHandler)

	// サーバーの起動
	router.Run(":8080")
}

func setupRoutes(router *gin.Engine, userHandler *handlers.UserHandler, studySetHandler *handlers.StudySetHandler, flashcardHandler *handlers.FlashcardHandler) {
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

	// 学習セット関連のルートをグループ化
	studySetRoutes := router.Group("/studysets")
	{
		studySetRoutes.POST("/", studySetHandler.CreateStudySet)
		studySetRoutes.GET("/:id", studySetHandler.GetStudySetByID)
		studySetRoutes.GET("/user/:userID", studySetHandler.GetStudySetsByUserID)
		studySetRoutes.PUT("/:id", studySetHandler.UpdateStudySet)
		studySetRoutes.DELETE("/:id", studySetHandler.DeleteStudySet)
		studySetRoutes.GET("/search", studySetHandler.SearchStudySetsByTitle)
	}

	// フラッシュカード関連のルートをグループ化
	flashcardRoutes := router.Group("/flashcards")
	{
		flashcardRoutes.POST("/", flashcardHandler.CreateFlashcard)
		flashcardRoutes.GET("/:id", flashcardHandler.GetFlashcardByID)
		flashcardRoutes.GET("/studyset/:studySetID", flashcardHandler.GetFlashcardsByStudySetID)
		flashcardRoutes.PUT("/:id", flashcardHandler.UpdateFlashcard)
		flashcardRoutes.DELETE("/:id", flashcardHandler.DeleteFlashcard)
	}
}
