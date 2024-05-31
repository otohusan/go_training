package flashcard

import (
	"errors"
	"go-training/domain/model"
	"go-training/domain/repository"
	inmemory "go-training/infrastructure/InMemory"
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

	if _, exists := r.flashcards[flashcard.ID]; exists {
		return errors.New("flashcard already exists")
	}

	r.flashcards[flashcard.ID] = flashcard
	return nil
}

func (r *FlashcardRepository) GetByID(id string) (*model.Flashcard, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	flashcard, exists := r.flashcards[id]
	if !exists {
		return nil, errors.New("flashcard not found")
	}

	return flashcard, nil
}

func (r *FlashcardRepository) GetByStudySetID(studySetID string) ([]*model.Flashcard, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var studySetFlashcards []*model.Flashcard
	for _, flashcard := range r.flashcards {
		if flashcard.StudySetID == studySetID {
			studySetFlashcards = append(studySetFlashcards, flashcard)
		}
	}
	return studySetFlashcards, nil
}

func (r *FlashcardRepository) Update(flashcard *model.Flashcard) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.flashcards[flashcard.ID]; !exists {
		return errors.New("flashcard not found")
	}

	r.flashcards[flashcard.ID] = flashcard
	return nil
}

func (r *FlashcardRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.flashcards[id]; !exists {
		return errors.New("flashcard not found")
	}

	delete(r.flashcards, id)
	return nil
}
