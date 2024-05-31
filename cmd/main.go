package main

import (
	"go-training/application/service"
	"go-training/handlers"
	"go-training/middleware"

	"github.com/gin-gonic/gin"

	// userRepo "go-training/infrastructure/rest"
	favorite "go-training/infrastructure/InMemory/Favorite"
	flashcard "go-training/infrastructure/InMemory/FlashCard"
	studySet "go-training/infrastructure/InMemory/StudySet"
	user "go-training/infrastructure/InMemory/User"
)

func main() {

	// リポジトリの初期化
	userRepo := user.NewUserRepository()
	studySetRepo := studySet.NewStudySetRepository()
	flashcardRepo := flashcard.NewFlashcardRepository()
	favoriteRepo := favorite.NewFavoriteRepository()

	// サービスの初期化
	userService := service.NewUserService(userRepo)
	studySetService := service.NewStudySetService(studySetRepo)
	flashcardService := service.NewFlashcardService(flashcardRepo)
	favoriteService := service.NewFavoriteService(favoriteRepo)

	// ハンドラーの初期化
	userHandler := handlers.NewUserHandler(userService)
	studySetHandler := handlers.NewStudySetHandler(studySetService)
	flashcardHandler := handlers.NewFlashcardHandler(flashcardService, studySetService)
	favoriteHandler := handlers.NewFavoriteHandler(favoriteService)

	// Ginのルーターを設定
	router := gin.Default()

	// ルートの設定
	setupRoutes(router, userHandler, studySetHandler, flashcardHandler, favoriteHandler)

	// サーバーの起動
	router.Run(":8080")
}

func setupRoutes(router *gin.Engine, userHandler *handlers.UserHandler, studySetHandler *handlers.StudySetHandler,
	flashcardHandler *handlers.FlashcardHandler, favoriteHandler *handlers.FavoriteHandler) {
	// ユーザー関連のルートをグループ化
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/", userHandler.GetAllUsers)
		userRoutes.POST("/", userHandler.CreateUserWithEmail)
		userRoutes.GET("/:userID", userHandler.GetUserByID)
		userRoutes.POST("/login/email", userHandler.LoginWithEmail)
		userRoutes.GET("/username/:username", userHandler.GetUserByUsername)
		// favoriteHandlerをここで呼び出すのが気になるけど、エンドポイントがこっちの方が直感的
		userRoutes.GET("/:userID/favorite", favoriteHandler.GetFavoriteStudySetsByUserID)
	}

	// 認証が必要なユーザー関連のルート
	authUserRoutes := router.Group("/users")
	authUserRoutes.Use(middleware.AuthMiddleware()) // 認証ミドルウェアの適用
	{
		// nameしか変えられない
		authUserRoutes.PUT("/:userID", userHandler.UpdateUser)
		authUserRoutes.DELETE("/:userID", userHandler.DeleteUser)
	}

	// 学習セット関連のルートをグループ化
	studySetRoutes := router.Group("/studysets")
	{
		studySetRoutes.GET("/:id", studySetHandler.GetStudySetByID)
		studySetRoutes.GET("/user/:userID", studySetHandler.GetStudySetsByUserID)
		studySetRoutes.GET("/search", studySetHandler.SearchStudySetsByTitle)
	}

	// 認証が必要な学習セット関連のルート
	authStudySetRoutes := router.Group("/studysets")
	authStudySetRoutes.Use(middleware.AuthMiddleware()) // 認証ミドルウェアの適用
	{
		authStudySetRoutes.POST("/", studySetHandler.CreateStudySet)
		authStudySetRoutes.PUT("/:studySetID", studySetHandler.UpdateStudySet)
		authStudySetRoutes.DELETE("/:studySetID", studySetHandler.DeleteStudySet)
	}

	// フラッシュカード関連のルートをグループ化
	flashcardRoutes := router.Group("/flashcards")
	{
		flashcardRoutes.GET("/:id", flashcardHandler.GetFlashcardByID)
		flashcardRoutes.GET("/studyset/:studySetID", flashcardHandler.GetFlashcardsByStudySetID)
	}

	// 認証が必要なフラッシュカード関連のルート
	authFlashcardRoutes := router.Group("/flashcards")
	authFlashcardRoutes.Use(middleware.AuthMiddleware()) // 認証ミドルウェアの適用
	{
		authFlashcardRoutes.POST("/:studySetID", flashcardHandler.CreateFlashcard)
		authFlashcardRoutes.PUT("/:flashcardID/:studySetID", flashcardHandler.UpdateFlashcard)
		authFlashcardRoutes.DELETE("/:flashcardID/:studySetID", flashcardHandler.DeleteFlashcard)
	}

	// お気に入り関連のルート
	favoriteRoutes := router.Group("favorites")
	{
		favoriteRoutes.GET("/is_favorite", favoriteHandler.IsFavorite)
	}

	// 認証が必要なお気に入り関連のルート
	authFavoriteRoutes := router.Group("/favorites")
	authFavoriteRoutes.Use(middleware.AuthMiddleware()) // 認証ミドルウェアの適用
	{
		authFavoriteRoutes.POST("/user/:userID/studyset/:studySetID", favoriteHandler.AddFavorite)
		authFavoriteRoutes.DELETE("/user/:userID/studyset/:studySetID", favoriteHandler.RemoveFavorite)
		authFavoriteRoutes.GET("/user/:userID", favoriteHandler.GetFavoritesByUserID)
	}
}
