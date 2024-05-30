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

	err := h.studySetService.CreateStudySet(&studySet)
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

	studySets, err := h.studySetService.GetStudySetsByUserID(userID)
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

	id := c.Param("id")
	studySet.ID = id

	err := h.studySetService.UpdateStudySet(&studySet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "study set updated successfully"})
}

func (h *StudySetHandler) DeleteStudySet(c *gin.Context) {
	id := c.Param("id")

	err := h.studySetService.DeleteStudySet(id)
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
