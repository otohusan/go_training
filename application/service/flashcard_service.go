package service

import (
	"errors"
	"go-training/domain/model"
	"go-training/domain/repository"
)

// FlashCardが適切に値を持ってるか確かめる関数
func validateFlashCard(flashcard *model.Flashcard) error {
	if flashcard.Question == "" {
		return errors.New("question cannot be empty")
	}
	if flashcard.Answer == "" {
		return errors.New("answer cannot be empty")
	}
	if flashcard.StudySetID == "" {
		return errors.New("studySetID cannot be empty")
	}
	return nil
}

type FlashcardService struct {
	repo repository.FlashcardRepository
}

func NewFlashcardService(repo repository.FlashcardRepository) *FlashcardService {
	return &FlashcardService{
		repo: repo,
	}
}

func (s *FlashcardService) CreateFlashcard(flashcard *model.Flashcard) error {
	if err := validateFlashCard(flashcard); err != nil {
		return err
	}

	return s.repo.Create(flashcard)
}

func (s *FlashcardService) GetFlashcardByID(id string) (*model.Flashcard, error) {
	return s.repo.GetByID(id)
}

func (s *FlashcardService) GetFlashcardsByStudySetID(studySetID string) ([]*model.Flashcard, error) {
	return s.repo.GetByStudySetID(studySetID)
}

func (s *FlashcardService) UpdateFlashcard(flashcard *model.Flashcard) error {
	if err := validateFlashCard(flashcard); err != nil {
		return err
	}

	return s.repo.Update(flashcard)
}

func (s *FlashcardService) DeleteFlashcard(id string) error {
	return s.repo.Delete(id)
}
