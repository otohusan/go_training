package auth

import (
	"errors"
	"go-training/domain/model"
	"go-training/domain/repository"

	"github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("secret") // 実際の環境では安全にキーを管理してください。

type AuthService struct {
	repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (h *AuthService) RegisterUser(user model.CreatedUserData) error {
	userInfo := model.User{
		Name:     user.Username,
		Password: user.Password,
	}
	return h.repo.CreateWithEmail(&userInfo)
}

func (h *AuthService) ParseToken(utoken string) (string, error) {

	token, err := jwt.Parse(utoken, func(token *jwt.Token) (interface{}, error) {
		// 署名アルゴリズムの検証
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("検証できない")
		}
		return mySigningKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id := claims["id"].(string)

		return id, nil
	} else {
		return "", errors.New("idが取得できない")
	}

}
