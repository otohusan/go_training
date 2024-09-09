package service_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"

	"go-training/application/service"
	"go-training/domain/model"
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

	// テストケース1: 正常な学習セット取得
	t.Run("Get valid study set by ID", func(t *testing.T) {
		studySetID := "123"
		userID := "user_456"

		createdAt := time.Now().Truncate(time.Second) // ミリ秒以下を削除
		updatedAt := createdAt

		// モックリポジトリの期待される返答
		expectedStudySet := &model.StudySet{
			ID:          studySetID,
			UserID:      userID,
			Title:       "Test Study Set",
			Description: "This is a test description",
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
			Flashcards:  nil, // フラッシュカードなし
		}

		// モックの期待される振る舞いを定義
		mockStudySetRepo.EXPECT().GetByID(studySetID).Return(expectedStudySet, nil)
		mockFlashcardRepo.EXPECT().GetByStudySetID(studySetID).Return(nil, nil)

		// サービスの呼び出し
		studySet, err := studySetService.GetStudySetByID(studySetID)

		// 結果の検証
		assert.NoError(t, err)
		assert.NotNil(t, studySet)
		assert.Equal(t, expectedStudySet, studySet)
	})

}
