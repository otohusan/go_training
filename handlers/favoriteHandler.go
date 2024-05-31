package handlers

import (
	"go-training/application/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FavoriteHandler struct {
	favoriteService *service.FavoriteService
}

func NewFavoriteHandler(favoriteService *service.FavoriteService) *FavoriteHandler {
	return &FavoriteHandler{favoriteService: favoriteService}
}

func (h *FavoriteHandler) AddFavorite(c *gin.Context) {
	userID := c.Param("userID")
	studySetID := c.Param("studySetID")

	// middlewareで設定されたAuthUserIDの取得
	AuthUserID, exists := c.Get("AuthUserID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "userID not found in context"})
		return
	}

	// サービスを呼び出す
	err := h.favoriteService.AddFavorite(AuthUserID.(string), userID, studySetID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "favorite added successfully"})
}

func (h *FavoriteHandler) RemoveFavorite(c *gin.Context) {
	userID := c.Param("userID")
	studySetID := c.Param("studySetID")

	// middlewareで設定されたAuthUserIDの取得
	AuthUserID, exists := c.Get("AuthUserID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "userID not found in context"})
		return
	}

	err := h.favoriteService.RemoveFavorite(AuthUserID.(string), userID, studySetID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "favorite removed successfully"})
}

func (h *FavoriteHandler) GetFavoritesByUserID(c *gin.Context) {
	userID := c.Param("userID")

	favorites, err := h.favoriteService.GetFavoritesByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, favorites)
}

func (h *FavoriteHandler) IsFavorite(c *gin.Context) {
	userID := c.Param("userID")
	studySetID := c.Param("studySetID")

	isFavorite, err := h.favoriteService.IsFavorite(userID, studySetID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"is_favorite": isFavorite})
}

func (h *FavoriteHandler) GetFavoriteStudySetsByUserID(c *gin.Context) {
	userID := c.Param("userID")

	studySets, err := h.favoriteService.GetFavoriteStudySetsByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, studySets)
}
