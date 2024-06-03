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
	query := `INSERT INTO email_verifications (email, token, username, password) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(query, verificationValue.Email, verificationValue.Token, verificationValue.Username, verificationValue.Password)
	if err != nil {
		return err
	}
	return nil
}

func (r *VerificationRepository) GetVerificationInfoByToken(token string) (*model.EmailVerification, error) {
	return nil, nil
}

func (r *VerificationRepository) DeleteVerificationToken(token string) error {
	return nil
}
