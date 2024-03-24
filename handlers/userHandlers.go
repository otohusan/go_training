package http

import (
	"go-training/application/service"
	"go-training/application/service/auth"
	products "go-training/handlers/auth"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
	authService *auth.AuthService
}

// NewUserHandler はUserHandlerの新しいインスタンスを作成します。
func createUserHandler(userService *service.UserService, authService *auth.AuthService) *UserHandler {
	return &UserHandler{
		userService: userService,
		authService: authService,
	}
}

// RegisterRoutes はルーターにエンドポイントを登録します。
func SetupRoutes(userService *service.UserService, authService *auth.AuthService, productService *auth.AuthService) *gin.Engine {
	router := gin.Default()

	productsHandler := products.NewProductHandler(*productService)

	// ここでハンドラーをインスタンス化し、ルーティングに登録
	userHandler := createUserHandler(userService, authService)

	router.GET("/users/:id", userHandler.ReturnUser)
	router.GET("/users/", userHandler.GetUserList)
	router.POST("/users/", userHandler.CreateUser)
	router.DELETE("/users/:id", userHandler.DeleteUser)
	router.POST("/auth/", userHandler.Login)
	router.POST("/auth/parse", userHandler.ParseToken)
	router.POST("/auth/register", userHandler.Register)
	router.POST("/post/", userHandler.GetPost)
	router.POST("/post/create", userHandler.CreatePost)
	router.GET("/post/tame", productsHandler.CreateProduct)

	return router
}
