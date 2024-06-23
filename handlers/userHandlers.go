package handlers

import (
	"database/sql"
	"go-training/application/service"
	"go-training/domain/model"
	"go-training/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) CreateUserWithEmail(c *gin.Context) {
	// emailのsql.NullStringを直接Bindできないから
	// 受け取るようのやつを用意
	type CreateUserRequest struct {
		Username string
		Password string
		Email    string
	}

	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// sql.NullStringに変換
	email := sql.NullString{
		String: req.Email,
		Valid:  req.Email != "",
	}

	// NOTICE: service層ですべき行動が行われてしまってる
	// パスワードのハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	// ユーザの作成
	user := model.User{
		ID:       "",
		Name:     req.Username,
		Password: string(hashedPassword),
		Email:    email,
	}

	// ユーザの作成
	_, err = h.userService.CreateUserWithEmail(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// JWTトークンの生成
	tokenString, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("userID")
	authUserID, exists := c.Get("AuthUserID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "userID not found in context"})
		return
	}

	if authUserID.(string) != id {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
		return
	}

	user, err := h.userService.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	// parseNullString を使って user.Email を *string に変換
	email := utils.ParseNullString(user.Email)

	response := gin.H{
		"ID":        user.ID,
		"Name":      user.Name,
		"Email":     email,
		"CreatedAt": user.CreatedAt,
	}

	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) GetUserByUsername(c *gin.Context) {
	username := c.Param("username")

	user, err := h.userService.GetUserByUsername(username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	// parseNullString を使って user.Email を *string に変換
	email := utils.ParseNullString(user.Email)

	response := gin.H{
		"ID":        user.ID,
		"Name":      user.Name,
		"Email":     email,
		"CreatedAt": user.CreatedAt,
	}

	c.JSON(http.StatusOK, response)
}

// 一般的な情報のみを返す
// 情報のフィルタリングはサービスが行ってる
func (h *UserHandler) GetPublicUserInfo(c *gin.Context) {
	userID := c.Param("userID")
	user, err := h.userService.GetPublicUserInfo(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	id := c.Param("userID")

	// middlewareで設定されたAuthUserIDの取得
	AuthUserID, exists := c.Get("AuthUserID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "userID not found in context"})
		return
	}

	user.ID = id

	err := h.userService.UpdateUser(AuthUserID.(string), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user updated successfully"})
}

// NOTICE:
// 今のところあまり使いたくない
// 本当にDBから削除するのか、仮想的に削除するかを悩んでる
// study_setテーブルにON DELETE CASCADEを割り当ててないから、基本消せない
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("userID")

	// middlewareで設定されたAuthUserIDの取得
	AuthUserID, exists := c.Get("AuthUserID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "userID not found in context"})
		return
	}

	err := h.userService.DeleteUser(AuthUserID.(string), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}

func (h *UserHandler) LoginWithEmail(c *gin.Context) {
	// emailのsql.NullStringを直接Bindできないから
	// 受け取るようのやつを用意
	type LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// sql.NullStringに変換
	email := sql.NullString{
		String: req.Email,
		Valid:  req.Email != "",
	}

	// NOTICE: service層ですべき行動が行われてしまってる
	// 必要な情報があるか
	if !email.Valid || email.String == "" || req.Password == "" {
		log.Printf("Email or password is empty")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "情報が足りません"})
		return
	}

	// emailからユーザ情報を取得
	user, err := h.userService.GetUserByEmail(email.String)
	if err != nil {
		log.Printf("Error getting user by email: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "there are not the email"})
		return
	}

	// パスワードを比較
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		log.Printf("Password mismatch: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "password is not valid"})
		return
	}

	// JWTトークンの生成
	tokenString, err := utils.GenerateToken(user.ID)
	if err != nil {
		log.Printf("Error generating token: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	// トークンを返す
	log.Printf("User %s logged in successfully", user.Email.String)
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) IsEmailExist(c *gin.Context) {
	var request struct {
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	result, err := h.userService.IsEmailExist(request.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}

func (h *UserHandler) IsUsernameExist(c *gin.Context) {
	var request struct {
		Username string `json:"username" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	result, err := h.userService.IsUsernameExist(request.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}
