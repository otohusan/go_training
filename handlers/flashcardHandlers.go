package handlers

import (
	"go-training/application/service"
	"go-training/domain/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FlashcardHandler struct {
	flashcardService *service.FlashcardService
	// studySetのサービスをここで読んでいいのかは疑問だけど
	studySetService *service.StudySetService
}

func NewFlashcardHandler(flashcardService *service.FlashcardService, studySetService *service.StudySetService) *FlashcardHandler {
	return &FlashcardHandler{flashcardService: flashcardService, studySetService: studySetService}
}

func (h *FlashcardHandler) CreateFlashcard(c *gin.Context) {
	var flashcard model.Flashcard
	if err := c.ShouldBindJSON(&flashcard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// ユーザが適切か確認する手順
	studySetID := c.Param("studySetID")
	AuthUserID, exists := c.Get("AuthUserID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "userID not found in context"})
		return
	}

	// studySetのサービスをここで読んでいいのかは疑問だけど使ってる,良い方法他にあればいいな
	studySet, err := h.studySetService.GetStudySetByID(studySetID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "studySet doesn't exist"})
		return
	}

	// 認可できるか
	if studySet.UserID != AuthUserID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
		return
	}

	if err := h.flashcardService.CreateFlashcard(&flashcard); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "flashcard created successfully"})
}

func (h *FlashcardHandler) GetFlashcardByID(c *gin.Context) {
	id := c.Param("id")

	flashcard, err := h.flashcardService.GetFlashcardByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "flashcard not found"})
		return
	}

	c.JSON(http.StatusOK, flashcard)
}

func (h *FlashcardHandler) GetFlashcardsByStudySetID(c *gin.Context) {
	studySetID := c.Param("studySetID")

	flashcards, err := h.flashcardService.GetFlashcardsByStudySetID(studySetID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, flashcards)
}

func (h *FlashcardHandler) UpdateFlashcard(c *gin.Context) {
	var flashcard model.Flashcard
	if err := c.ShouldBindJSON(&flashcard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// flashcardの作成者を確かめるために色々取り出す
	flashcardID := c.Param("flashcardID")
	studySetID := c.Param("studySetID")
	AuthUserID, exists := c.Get("AuthUserID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "userID not found in context"})
		return
	}

	// studySetのサービスをここで読んでいいのかは疑問だけど使ってる,良い方法他にあればいいな
	studySet, err := h.studySetService.GetStudySetByID(studySetID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "studySet doesn't exist"})
		return
	}

	// 認可できるか
	if studySet.UserID != AuthUserID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
		return
	}

	flashcard.ID = flashcardID

	if err := h.flashcardService.UpdateFlashcard(&flashcard); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "flashcard updated successfully"})
}

func (h *FlashcardHandler) DeleteFlashcard(c *gin.Context) {
	id := c.Param("id")

	err := h.flashcardService.DeleteFlashcard(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "flashcard deleted successfully"})
}
