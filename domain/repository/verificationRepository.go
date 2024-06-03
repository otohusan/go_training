package repository

import "go-training/domain/model"

type EmailVerificationRepository interface {
	SaveVerificationToken(*model.EmailVerification) error
	GetVerificationInfoByToken(token string) (*model.EmailVerification, error)
	DeleteVerificationToken(token string) error
}
