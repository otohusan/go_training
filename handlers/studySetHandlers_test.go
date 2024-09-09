package handlers_test

import (
	"testing"

	"github.com/gin-gonic/gin"
	gomock "go.uber.org/mock/gomock"

	"go-training/application/service/mocks"
	"go-training/handlers"
)

func TestGetStudySetByID(t *testing.T) {

	// GoMockコントローラーの初期化
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// モックサービスの作成
	mockStudySetService := mocks.NewMockStudySetService(ctrl)
	// Ginのルーターを作成
	router := gin.Default()
	studySetHandler := handlers.NewStudySetHandler(mockStudySetService)
	router.GET("/studysets/:id", studySetHandler.GetStudySetByID)

}
