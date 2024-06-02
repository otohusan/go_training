package repository

import "go-training/domain/model"

type FlashcardRepository interface {
	Create(authUSerID string, flashcard *model.Flashcard) error
	GetByID(id string) (*model.Flashcard, error)
	GetByStudySetID(studySetID string) ([]*model.Flashcard, error)
	Update(authUSerID string, flashcard *model.Flashcard) error
	Delete(authUSerID, flashcardID string) error
}
