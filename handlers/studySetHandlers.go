package handlers

import (
	"go-training/application/service"
	"go-training/domain/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StudySetHandler struct {
	studySetService *service.StudySetService
}

func NewStudySetHandler(studySetService *service.StudySetService) *StudySetHandler {
	return &StudySetHandler{studySetService: studySetService}
}

func (h *StudySetHandler) CreateStudySet(c *gin.Context) {
	var studySet model.StudySet
	if err := c.ShouldBindJSON(&studySet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	AuthUserID, exists := c.Get("AuthUserID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "userID not found in context"})
		return
	}

	studySet.UserID = AuthUserID.(string)

	err := h.studySetService.CreateStudySet(AuthUserID.(string), &studySet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "study set created successfully"})
}

func (h *StudySetHandler) GetStudySetByID(c *gin.Context) {
	id := c.Param("id")

	studySet, err := h.studySetService.GetStudySetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "study set not found"})
		return
	}

	c.JSON(http.StatusOK, studySet)
}

func (h *StudySetHandler) GetStudySetsByUserID(c *gin.Context) {
	userID := c.Param("userID")

	studySets, err := h.studySetService.GetStudySetsWithFlashcardsByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, studySets)
}

func (h *StudySetHandler) UpdateStudySet(c *gin.Context) {
	var studySet model.StudySet
	if err := c.ShouldBindJSON(&studySet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	studySetID := c.Param("studySetID")
	studySet.ID = studySetID

	// middlewareで設定されたAuthUserIDの取得
	AuthUserID, exists := c.Get("AuthUserID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "userID not found in context"})
		return
	}

	err := h.studySetService.UpdateStudySet(AuthUserID.(string), studySetID, &studySet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "study set updated successfully"})
}

func (h *StudySetHandler) DeleteStudySet(c *gin.Context) {
	studySetID := c.Param("studySetID")

	// middlewareで設定されたAuthUserIDの取得
	AuthUserID, exists := c.Get("AuthUserID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "userID not found in context"})
		return
	}

	err := h.studySetService.DeleteStudySet(AuthUserID.(string), studySetID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "study set deleted successfully"})
}

func (h *StudySetHandler) SearchStudySetsByTitle(c *gin.Context) {
	title := c.Query("title")
	if title == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "title is empty"})
		return
	}

	studySets, err := h.studySetService.SearchStudySetsByTitle(title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, studySets)
}
