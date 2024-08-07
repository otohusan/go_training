package main

import (
	"database/sql"
	"fmt"
	"go-training/application/service"
	"go-training/handlers"
	"go-training/middleware"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"go-training/infrastructure/rest/favorite"
	"go-training/infrastructure/rest/flashcard"
	"go-training/infrastructure/rest/googleUser"
	"go-training/infrastructure/rest/studyset"
	"go-training/infrastructure/rest/user"
	"go-training/infrastructure/rest/verification"
	// user "go-training/infrastructure/InMemory/User"
	// studyset "go-training/infrastructure/InMemory/StudySet"
	// flashcard "go-training/infrastructure/InMemory/FlashCard"
	// favorite "go-training/infrastructure/InMemory/Favorite"
	// verification "go-training/infrastructure/InMemory/verification"
)

func main() {

	env := os.Getenv("GIN_MODE")
	if env == "" {
		env = "development"
	}

	// .envファイルの読み込み
	err := godotenv.Load(fmt.Sprintf(".env.%s", env))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 環境変数の取得
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// データベース接続の設定
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s host=%s port=%s",
		dbUser, dbPassword, dbName, dbSSLMode, dbHost, dbPort)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// リポジトリの初期化
	userRepo := user.NewUserRepository(db)
	studySetRepo := studyset.NewStudySetRepository(db)
	flashcardRepo := flashcard.NewFlashcardRepository(db)
	favoriteRepo := favorite.NewFavoriteRepository(db)
	verificationRepo := verification.NewVerificationRepository(db)
	googleUserRepo := googleUser.NewGoogleUserRepository(db)

	// サービスの初期化
	userService := service.NewUserService(userRepo)
	studySetService := service.NewStudySetService(studySetRepo, flashcardRepo)
	flashcardService := service.NewFlashcardService(flashcardRepo)
	favoriteService := service.NewFavoriteService(favoriteRepo, flashcardRepo)
	authService := service.NewAuthService(userRepo, verificationRepo, googleUserRepo)

	// ハンドラーの初期化
	userHandler := handlers.NewUserHandler(userService)
	studySetHandler := handlers.NewStudySetHandler(studySetService)
	flashcardHandler := handlers.NewFlashcardHandler(flashcardService, studySetService)
	favoriteHandler := handlers.NewFavoriteHandler(favoriteService)
	authHandler := handlers.NewAuthHandler(authService)

	// Ginのルーターを設定
	router := gin.Default()

	// ルートの設定
	setupRoutes(router, userHandler, studySetHandler, flashcardHandler, favoriteHandler, authHandler)

	// サーバーの起動
	router.Run(":8080")
}

func setupRoutes(router *gin.Engine, userHandler *handlers.UserHandler, studySetHandler *handlers.StudySetHandler,
	flashcardHandler *handlers.FlashcardHandler, favoriteHandler *handlers.FavoriteHandler, authHandler *handlers.AuthHandler) {
	// CORSの設定
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "https://konwalk.jp"}, // クライアントのオリジンを許可
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 信頼できるプロキシを空に設定
	// CDNなんかは今のとこしよしてないから、これで良いはず
	// デプロイ環境次第では変える必要があるかも
	if err := router.SetTrustedProxies([]string{}); err != nil {
		log.Fatal(err)
	}

	// ユーザー関連のルートをグループ化
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/", userHandler.GetAllUsers)
		userRoutes.POST("/", userHandler.CreateUserWithEmail)
		userRoutes.GET("/public/:userID", userHandler.GetPublicUserInfo)
		userRoutes.POST("/login/email", userHandler.LoginWithEmail)
		// favoriteHandlerをここで呼び出すのが気になるけど、エンドポイントがこっちの方が直感的
		userRoutes.GET("/:userID/favorite", favoriteHandler.GetFavoriteStudySetsByUserID)
		userRoutes.POST("/email-exists", userHandler.IsEmailExist)
		userRoutes.POST("/username-exists", userHandler.IsUsernameExist)
	}

	// 認証が必要なユーザー関連のルート
	authUserRoutes := router.Group("/users")
	authUserRoutes.Use(middleware.AuthMiddleware()) // 認証ミドルウェアの適用
	{
		authUserRoutes.GET("/:userID", userHandler.GetUserByID)
		authUserRoutes.GET("/username/:username", userHandler.GetUserByUsername)
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
		authStudySetRoutes.POST("copy/:userID", studySetHandler.CopyStudySetForMe)
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
		authFlashcardRoutes.POST("/generate", flashcardHandler.GenerateAnswerWithAI)
		authFlashcardRoutes.POST("/:studySetID", flashcardHandler.CreateFlashcard)
		authFlashcardRoutes.PUT("/:flashcardID", flashcardHandler.UpdateFlashcard)
		authFlashcardRoutes.DELETE("/:flashcardID", flashcardHandler.DeleteFlashcard)
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

	// user登録のルート
	router.POST("/register/email", authHandler.RegisterWithEmail)
	router.GET("/verify", authHandler.VerifyEmail)
	// Googleログインのルート
	router.POST("/auth/google", authHandler.GoogleLogin)
}
