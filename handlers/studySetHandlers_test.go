package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"

	"go-training/application/service/mocks"
	"go-training/domain/model"
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

	// テストケース1: 正常な学習セット取得
	t.Run("Get valid study set", func(t *testing.T) {
		// モックの期待される動作を定義
		createdAt := time.Now().Truncate(time.Second) // ミリ秒以下を削除
		updatedAt := createdAt

		expectedStudySet := &model.StudySet{
			ID:          "123",
			UserID:      "user_456",
			Title:       "Test Study Set",
			Description: "This is a test study set",
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
			Flashcards:  nil, // この場合フラッシュカードなしでテスト
		}

		// モックサービスがGetStudySetByIDを呼ばれた際の振る舞いを定義
		mockStudySetService.EXPECT().GetStudySetByID("123").Return(expectedStudySet, nil)

		// HTTPリクエストを作成
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/studysets/123", nil)

		// リクエストを処理
		router.ServeHTTP(w, req)

		// レスポンスの検証
		assert.Equal(t, http.StatusOK, w.Code)
		expectedBody := `{
			"id":"123",
			"user_id":"user_456",
			"title":"Test Study Set",
			"description":"This is a test study set",
			"created_at":"` + expectedStudySet.CreatedAt.Format(time.RFC3339) + `",
			"updated_at":"` + expectedStudySet.UpdatedAt.Format(time.RFC3339) + `",
			"flashcards":null
		}`
		assert.JSONEq(t, expectedBody, w.Body.String())
	})

}
