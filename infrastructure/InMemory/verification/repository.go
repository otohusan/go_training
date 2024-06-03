package verification

import (
	"go-training/domain/model"
	"go-training/domain/repository"
	"sync"
)

type VerificationRepository struct {
	mu sync.Mutex
}

func NewVerificationRepository() repository.EmailVerificationRepository {
	return &VerificationRepository{}
}

func (r *VerificationRepository) SaveVerificationToken(*model.EmailVerification) error {
	return nil
}

func (r *VerificationRepository) GetVerificationInfoByToken(token string) (*model.EmailVerification, error) {
	return nil, nil
}

func (r *VerificationRepository) DeleteVerificationToken(token string) error {
	return nil
}
