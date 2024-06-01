package flashcard

import (
	"database/sql"
	"go-training/domain/model"
)

type FlashcardRepository struct {
	db *sql.DB
}

func NewFlashcardRepository(db *sql.DB) *FlashcardRepository {
	return &FlashcardRepository{db: db}
}

func (r *FlashcardRepository) Create(authUserID string, flashcard *model.Flashcard) error {
	return nil
}

func (r *FlashcardRepository) GetByID(id string) (*model.Flashcard, error) {
	return nil, nil
}

func (r *FlashcardRepository) GetByStudySetID(studySetID string) ([]*model.Flashcard, error) {
	return nil, nil
}

func (r *FlashcardRepository) Update(authUserID string, flashcard *model.Flashcard) error {
	return nil
}

func (r *FlashcardRepository) Delete(authUserID, studySetID, flashcardID string) error {
	return nil
}
