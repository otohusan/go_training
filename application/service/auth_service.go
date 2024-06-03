package service

import (
	"errors"
	"go-training/domain/model"
	"go-training/domain/repository"
	"net/mail"

	"github.com/google/uuid"
)

type AuthService struct {
	userRepo         repository.UserRepository
	verificationRepo repository.EmailVerificationRepository
}

func NewAuthService(userRepo repository.UserRepository, verificationRepo repository.EmailVerificationRepository) *AuthService {
	return &AuthService{userRepo: userRepo, verificationRepo: verificationRepo}
}

func (s *AuthService) Register(username, email, password string) (string, error) {
	// メールアドレスの形式を検証
	if _, err := mail.ParseAddress(email); err != nil {
		return "", errors.New("invalid email format")
	}

	// メールアドレスの重複チェック
	// ユーザを取得して空じゃなかったらエラー
	exists, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return "", err
	}
	if exists != nil {
		return "", errors.New("email already in use")
	}

	// 検証トークンの生成と保存
	token := uuid.New().String()
	verification := &model.EmailVerification{
		Email:    email,
		Token:    token,
		Username: username,
		Password: password,
	}

	if err := s.verificationRepo.SaveVerificationToken(verification); err != nil {
		return "", err
	}

	// 検証メール送信

	return "verification email sent, please check your email for verification", nil

}
