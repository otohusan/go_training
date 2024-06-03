package verification

import (
	"database/sql"
	"go-training/domain/model"
	"go-training/domain/repository"
)

type VerificationRepository struct {
	db *sql.DB
}

func NewVerificationRepository(db *sql.DB) repository.EmailVerificationRepository {
	return &VerificationRepository{db: db}
}

func (r *VerificationRepository) SaveVerificationToken(verificationValue *model.EmailVerification) error {
	return nil
}

func (r *VerificationRepository) GetVerificationInfoByToken(token string) (*model.EmailVerification, error) {
	return nil, nil
}

func (r *VerificationRepository) DeleteVerificationToken(token string) error {
	return nil
}
