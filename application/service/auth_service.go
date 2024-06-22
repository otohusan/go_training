package service

import (
	"database/sql"
	"errors"
	"go-training/domain/model"
	"go-training/domain/repository"
	"go-training/utils"
	"log"
	"net/mail"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo         repository.UserRepository
	verificationRepo repository.EmailVerificationRepository
	googleUserRepo   repository.GoogleUserRepository
}

func NewAuthService(userRepo repository.UserRepository, verificationRepo repository.EmailVerificationRepository, googleUserRepo repository.GoogleUserRepository) *AuthService {
	return &AuthService{userRepo: userRepo, verificationRepo: verificationRepo, googleUserRepo: googleUserRepo}
}

func (s *AuthService) RegisterWithEmail(username, email, password string) (string, error) {

	// 必要情報があるかチェック
	if username == "" {
		return "", errors.New("username cannot be empty")
	}
	if email == "" {
		return "", errors.New("email cannot be empty")
	}
	if password == "" {
		return "", errors.New("password cannot be empty")
	}

	// メールアドレスの形式を検証
	if _, err := mail.ParseAddress(email); err != nil {
		return "", errors.New("invalid email format")
	}

	// メールアドレスの重複チェック
	// ユーザを取得して空じゃなかったらエラー
	exists, err := s.userRepo.IsEmailExist(email)
	if err != nil {
		return "", err
	}
	if exists {
		return "", errors.New("email already in use")
	}

	// パスワードのハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	// 検証トークンの生成と保存
	token := uuid.New().String()
	verification := &model.EmailVerification{
		Email:    email,
		Token:    token,
		Username: username,
		Password: string(hashedPassword),
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
		Email:    sql.NullString{String: verification.Email, Valid: true},
		Password: verification.Password,
	}
	if _, err := s.userRepo.CreateWithEmail(user); err != nil {
		return "", err
	}

	// トークンを削除
	if err := s.verificationRepo.DeleteVerificationToken(token); err != nil {
		return "", err
	}

	return "email verified and user created successfully", nil
}

func (s *AuthService) CreateOrGetUser(AccessToken string) (string, error) {
	// Googleのユーザー情報を取得
	googleUserInfo, err := utils.FetchGoogleUserInfo(AccessToken)
	if err != nil {
		log.Printf("googleユーザーの取得に失敗: %v", err)
		return "", err
	}

	// googleIDが登録されてるか確認
	googleUser, err := s.googleUserRepo.GetByGoogleID(googleUserInfo.ID)
	if err != nil {
		log.Printf("GoogleID確認時にエラー発生: %v", err)
		return "", err
	}

	// ユーザが存在すれば、取得して、IDを返す
	if googleUser != nil {
		user, err := s.userRepo.GetByID(googleUser.UserID)
		if err != nil {
			log.Printf("ユーザー取得時にエラー発生: %v", err)
			return "", err
		}

		return user.ID, nil
	}

	user := &model.User{
		Name: googleUserInfo.Name,
	}

	userID, err := s.userRepo.CreateWithGoogle(user)
	if err != nil {
		log.Printf("Googleからのユーザー作成時にエラー発生: %v", err)
		return "", err
	}

	googleUser = &model.GoogleUser{
		UserID:   userID,
		GoogleID: googleUserInfo.ID,
	}

	err = s.googleUserRepo.Create(googleUser)
	if err != nil {
		log.Printf("Googleテーブルに追加時にエラー発生: %v", err)
		return "", err
	}

	return user.ID, nil
}
