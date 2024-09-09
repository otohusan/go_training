package handlers

import (
	service_interface "go-training/application/service/interface"
	"go-training/domain/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StudySetHandler struct {
	studySetService service_interface.StudySetService
}

func NewStudySetHandler(studySetService service_interface.StudySetService) *StudySetHandler {
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

	studySetID, err := h.studySetService.CreateStudySet(AuthUserID.(string), &studySet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "study set created successfully", "id": studySetID})
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
	title := c.Query("keyword")
	if title == "" {
		log.Println("Error: タイトルが入力されていない")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "title is empty"})
		return
	}

	studySets, err := h.studySetService.SearchStudySetsByKeyword(title)
	if err != nil {
		log.Println("検索中にエラー:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, studySets)
}

func (h *StudySetHandler) CopyStudySetForMe(c *gin.Context) {
	userID := c.Param("userID")
	AuthUserID, exists := c.Get("AuthUserID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "userID not found in context"})
		return
	}

	if userID != AuthUserID.(string) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "you can't copy"})
		return
	}

	var studySet model.StudySet
	if err := c.ShouldBindJSON(&studySet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err := h.studySetService.CopyStudySetForMe(studySet, AuthUserID.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not copy studySet"})
	}

	c.JSON(http.StatusOK, "コピーが完了")

}
