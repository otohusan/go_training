package http

import (
	"go-training/domain/model"
	"net/http"

	// "strings"

	"github.com/gin-gonic/gin"
	// "github.com/google/uuid"
)

type CreateUserRequest struct {
	Name string `json:"name"`
}

func (h *UserHandler) ReturnUser(c *gin.Context) {
	id := c.Param("id")
	user, err := h.userService.ReturnUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// func (h *UserHandler) GetUserList(c *gin.Context) {
// 	userList, err := h.userService.GetUserList()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, userList)
// }

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req CreateUserRequest

	// リクエストボディから JSON データをバインド
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser := model.User{Name: req.Name, ID: "1"}
	err := h.userService.Create(&createdUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "ユーザー作成完了")
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	err := h.userService.DeleteUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "ユーザは削除されました")
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	c.JSON(http.StatusOK, "ユーザは削除されました")
}

type GetPostRequest struct {
	ID string `json:"id"`
}

// func (h *UserHandler) GetPost(c *gin.Context) {
// 	var req GetPostRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "認証が足りない"})
// 		return
// 	}

// 	res, err := h.userService.GetPosts(req.ID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "IDが存在しない"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, res)
// }

// func (h *UserHandler) CreatePost(c *gin.Context) {
// 	var req model.RequestedPost
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "認証が足りない"})
// 	}

// 	authHeader := c.GetHeader("Authorization")
// 	token := strings.TrimPrefix(authHeader, "Bearer ")

// 	if token == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "トークンが空です"})
// 	}

// 	id, err := h.authService.ParseToken(token)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 	}

// 	newUUID := uuid.New()

// 	newPost := model.Post{
// 		Title:  req.Title,
// 		Detail: req.Detail,
// 		ID:     newUUID.String(),
// 		Author: id,
// 	}

// 	h.userService.CreatePost(newPost)
// }
