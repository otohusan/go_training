package verification

import (
	"errors"
	"go-training/domain/model"
	"go-training/domain/repository"
	inmemory "go-training/infrastructure/InMemory"
	"go-training/utils"
	"sync"
)

type VerificationRepository struct {
	mu sync.Mutex
}

func NewVerificationRepository() repository.EmailVerificationRepository {
	return &VerificationRepository{}
}

func (r *VerificationRepository) SaveVerificationToken(verificationValue *model.EmailVerification) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// EmailVerificationを保存
	inmemory.EmailVerifications = append(inmemory.EmailVerifications, verificationValue)
	return nil
}

func (r *VerificationRepository) GetVerificationInfoByToken(token string) (*model.EmailVerification, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, v := range inmemory.EmailVerifications {
		if v.Token == token {
			return v, nil
		}
	}

	return nil, errors.New("invalid or expired token")
}

func (r *VerificationRepository) DeleteVerificationToken(token string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, v := range inmemory.EmailVerifications {
		if v.Token == token {
			utils.RemoveElementFromSlice(inmemory.EmailVerifications, i)
			return nil
		}
	}

	return errors.New("token not found")
}
