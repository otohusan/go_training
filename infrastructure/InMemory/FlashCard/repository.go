package flashcard

import (
	"errors"
	"go-training/domain/model"
	"go-training/domain/repository"
	inmemory "go-training/infrastructure/InMemory"
	"go-training/utils"
	"sync"

	"github.com/google/uuid"
)

type FlashcardRepository struct {
	mu         sync.Mutex
	flashcards map[string]*model.Flashcard
}

func NewFlashcardRepository() repository.FlashcardRepository {
	repo := &FlashcardRepository{
		flashcards: make(map[string]*model.Flashcard),
	}
	for _, flashcard := range inmemory.Flashcards {
		repo.flashcards[flashcard.ID] = flashcard
	}
	return repo
}

func (r *FlashcardRepository) Create(flashcard *model.Flashcard) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	// TODO:本人確認が必要

	// 外部キー制約
	isStudySetExists := false
	for _, studySet := range inmemory.StudySets {
		if studySet.ID == flashcard.StudySetID {
			isStudySetExists = true
		}
	}
	if !isStudySetExists {
		return errors.New("flashCard doesn't exists")
	}

	// uuid作成
	flashcard.ID = uuid.New().String()

	return nil
}

func (r *FlashcardRepository) GetByID(id string) (*model.Flashcard, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, flashcard := range inmemory.Flashcards {
		if flashcard.ID == id {
			return flashcard, nil
		}
	}

	return nil, errors.New("flashcard not found")

}

func (r *FlashcardRepository) GetByStudySetID(studySetID string) ([]*model.Flashcard, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var studySetFlashcards []*model.Flashcard
	for _, flashcard := range inmemory.Flashcards {
		if flashcard.StudySetID == studySetID {
			studySetFlashcards = append(studySetFlashcards, flashcard)
		}
	}
	return studySetFlashcards, nil
}

func (r *FlashcardRepository) Update(flashcard *model.Flashcard) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	// TODO:本人確認が必要

	for i, flashcardFromDB := range inmemory.Flashcards {
		if flashcardFromDB.ID == flashcard.ID {
			inmemory.Flashcards[i] = flashcard
			return nil
		}
	}

	return errors.New("flashcard not found")

}

func (r *FlashcardRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	// TODO:本人確認が必要

	for i, flashcardFromDB := range inmemory.Flashcards {
		if flashcardFromDB.ID == id {
			inmemory.Flashcards = utils.RemoveElementFromSlice(inmemory.Flashcards, i)
			return nil
		}
	}

	return errors.New("flashcard not found")
}
