package repository

import "go-training/domain/model"

type FlashcardRepository interface {
	Create(flashcard *model.Flashcard) error
	GetByID(id string) (*model.Flashcard, error)
	GetByStudySetID(studySetID string) ([]*model.Flashcard, error)
	Update(flashcard *model.Flashcard) error
	Delete(id string) error
}
