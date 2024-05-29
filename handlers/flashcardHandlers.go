package handlers

import (
	"go-training/application/service"
	"go-training/domain/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FlashcardHandler struct {
	flashcardService *service.FlashcardService
}

func NewFlashcardHandler(flashcardService *service.FlashcardService) *FlashcardHandler {
	return &FlashcardHandler{flashcardService: flashcardService}
}

func (h *FlashcardHandler) CreateFlashcard(c *gin.Context) {
	var flashcard model.Flashcard
	if err := c.ShouldBindJSON(&flashcard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err := h.flashcardService.CreateFlashcard(&flashcard)
	if err != nil {
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

	id := c.Param("id")
	flashcard.ID = id

	err := h.flashcardService.UpdateFlashcard(&flashcard)
	if err != nil {
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
