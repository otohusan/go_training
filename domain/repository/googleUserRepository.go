package repository

import "go-training/domain/model"

type GoogleUserRepository interface {
	Create(googleUser *model.GoogleUser) error
	GetByGoogleID(googleID string) (*model.GoogleUser, error)
}
