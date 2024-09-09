package service_test

import (
	"testing"

	gomock "go.uber.org/mock/gomock"

	"go-training/application/service"
	"go-training/domain/repository/mocks"
)

func TestGetStudySetByID(t *testing.T) {
	// GoMockコントローラーの初期化
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// モックサービスの作成
	mockStudySetRepo := mocks.NewMockStudySetRepository(ctrl)
	mockFlashcardRepo := mocks.NewMockFlashcardRepository(ctrl)

	// サービスを初期化
	studySetService := service.NewStudySetService(mockStudySetRepo, mockFlashcardRepo)

}
