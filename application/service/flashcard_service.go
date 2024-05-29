package service

import (
	"go-training/domain/model"
	"go-training/domain/repository"
)

type FlashcardService struct {
	repo repository.FlashcardRepository
}

func NewFlashcardService(repo repository.FlashcardRepository) *FlashcardService {
	return &FlashcardService{
		repo: repo,
	}
}

func (s *FlashcardService) CreateFlashcard(flashcard *model.Flashcard) error {
	return s.repo.Create(flashcard)
}

func (s *FlashcardService) GetFlashcardByID(id string) (*model.Flashcard, error) {
	return s.repo.GetByID(id)
}

func (s *FlashcardService) GetFlashcardsByStudySetID(studySetID string) ([]*model.Flashcard, error) {
	return s.repo.GetByStudySetID(studySetID)
}

func (s *FlashcardService) UpdateFlashcard(flashcard *model.Flashcard) error {
	return s.repo.Update(flashcard)
}

func (s *FlashcardService) DeleteFlashcard(id string) error {
	return s.repo.Delete(id)
}
