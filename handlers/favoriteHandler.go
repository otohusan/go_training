package handlers

import (
	"go-training/application/service"
	"go-training/domain/model"
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
	var favorite model.Favorite

	userID := c.Param("userID")
	studySetID := c.Param("studySetID")

	if err := c.ShouldBindJSON(&favorite); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err := h.favoriteService.AddFavorite(userID, studySetID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "favorite added successfully"})
}

func (h *FavoriteHandler) RemoveFavorite(c *gin.Context) {
	userID := c.Param("userID")
	studySetID := c.Param("studySetID")

	err := h.favoriteService.RemoveFavorite(userID, studySetID)
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
