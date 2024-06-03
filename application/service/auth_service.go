package service

import (
	"errors"
	"go-training/domain/model"
	"go-training/domain/repository"
	"go-training/utils"
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

func (s *AuthService) RegisterWithEmail(username, email, password string) (string, error) {
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
	// NOTICE: endVerificationEmailは形だけの状態
	if err := utils.SendVerificationEmail(email, token); err != nil {
		return "", err
	}

	return "verification email sent, please check your email for verification", nil
}

// emailをユーザが開くとトークンを確認して、本登録
func (s *AuthService) VerifyEmail(token string) (string, error) {
	// トークンの確認と仮登録情報の取得
	verification, err := s.verificationRepo.GetVerificationInfoByToken(token)
	if err != nil {
		return "", errors.New("invalid or expired token")
	}

	// ユーザーの作成
	user := &model.User{
		Name:     verification.Username,
		Email:    verification.Email,
		Password: verification.Password,
	}
	if err := s.userRepo.CreateWithEmail(user); err != nil {
		return "", err
	}

	// トークンを削除
	if err := s.verificationRepo.DeleteVerificationToken(token); err != nil {
		return "", err
	}

	return "email verified and user created successfully", nil
}
