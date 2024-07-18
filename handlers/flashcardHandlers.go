package handlers

import (
	"bytes"
	"encoding/json"
	"go-training/application/service"
	"go-training/domain/model"
	"io"
	"log"
	"net/http"
	"os"

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

// クイズを作成してそのIDを受け取る
func (h *FlashcardHandler) CreateFlashcard(c *gin.Context) {
	var flashcard model.Flashcard
	if err := c.ShouldBindJSON(&flashcard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// ユーザが適切か確認する手順
	studySetID := c.Param("studySetID")
	flashcard.StudySetID = studySetID
	AuthUserID, exists := c.Get("AuthUserID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "userID not found in context"})
		return
	}

	flashcardID, err := h.flashcardService.CreateFlashcard(AuthUserID.(string), &flashcard)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "flashcard created successfully", "id": flashcardID})
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
	flashcard.ID = flashcardID

	// 認証IDを取り出す
	AuthUserID, exists := c.Get("AuthUserID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "userID not found in context"})
		return
	}

	// サービス呼び出し
	if err := h.flashcardService.UpdateFlashcard(AuthUserID.(string), &flashcard); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "flashcard updated successfully"})
}

func (h *FlashcardHandler) DeleteFlashcard(c *gin.Context) {

	flashcardID := c.Param("flashcardID")

	// 認証IDを取り出す
	AuthUserID, exists := c.Get("AuthUserID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "userID not found in context"})
		return
	}

	err := h.flashcardService.DeleteFlashcard(AuthUserID.(string), flashcardID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "flashcard deleted successfully"})
}

// AIに使う部分
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type RequestBody struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type ResponseBody struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func (h *FlashcardHandler) GenerateAnswerWithAI(c *gin.Context) {

	// 認証IDを取り出す
	_, exists := c.Get("AuthUserID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "userID not found in context"})
		return
	}

	var requestBody struct {
		Question string `json:"question"`
	}

	// Parse JSON request body
	if err := c.BindJSON(&requestBody); err != nil {
		log.Println("リクエストボディの解析エラー:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Load API key from environment variables
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Println("環境変数にAPIキーが見つかりません")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "API key not found"})
		return
	}

	// Prepare request body for OpenAI API
	openaiRequestBody := RequestBody{
		Model: "gpt-4o-mini",
		Messages: []Message{
			{
				Role:    "system",
				Content: "クイズの問題が与えられるので、その回答を出力してください。簡潔に50文字以内の回答をしてください。回答は文字数が少ない程好ましいです。",
			},
			{
				Role:    "user",
				Content: requestBody.Question,
			},
		},
	}

	reqBody, err := json.Marshal(openaiRequestBody)
	if err != nil {
		log.Println("リクエストボディのマーシャリングエラー:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error marshaling request body"})
		return
	}

	// Make HTTP request to OpenAI API
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Println("リクエストの作成エラー:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating request"})
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("リクエストの送信エラー:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending request"})
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		log.Println("APIからのエラーレスポンス:", bodyString)
		c.JSON(resp.StatusCode, gin.H{"error": bodyString})
		return
	}

	// Parse response from OpenAI API
	var responseBody ResponseBody
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		log.Println("レスポンスのデコードエラー:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding response"})
		return
	}

	// Extract the answer text from the response
	if len(responseBody.Choices) > 0 {
		answer := responseBody.Choices[0].Message.Content
		c.JSON(http.StatusOK, gin.H{"answer": answer})
	} else {
		log.Println("APIレスポンスにコンテンツが見つかりません")
		c.JSON(http.StatusOK, gin.H{"answer": ""})
	}
}
