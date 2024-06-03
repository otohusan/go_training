package verification

import (
	"database/sql"
	"errors"
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
	query := `SELECT email, token, username, password FROM email_verifications WHERE token = $1`
	row := r.db.QueryRow(query, token)

	verification := &model.EmailVerification{}
	err := row.Scan(&verification.Email, &verification.Token, &verification.Username, &verification.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("invalid or expired token")
		}
		return nil, err
	}

	return verification, nil
}

func (r *VerificationRepository) DeleteVerificationToken(token string) error {
	query := `DELETE FROM email_verifications WHERE token = $1`
	_, err := r.db.Exec(query, token)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("token not found")
		}
		return err
	}
	return nil
}
