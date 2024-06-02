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
	mu sync.Mutex
}

func NewFlashcardRepository() repository.FlashcardRepository {
	return &FlashcardRepository{}
}

func (r *FlashcardRepository) Create(authUserID string, flashcard *model.Flashcard) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	// パフォーマンスを考慮して
	// 本番のクエリを1回にするためにリポジトリで認可行う

	var studySet *model.StudySet

	// 外部キー制約
	for _, s := range inmemory.StudySets {
		if s.ID == flashcard.StudySetID {
			studySet = s
		}
	}
	if studySet == nil {
		return errors.New("studySet doesn't exists")
	}

	// 認可できるか
	if studySet.UserID != authUserID {
		return errors.New("not authorized to update flashcard")
	}

	// uuid作成
	flashcard.ID = uuid.New().String()

	inmemory.Flashcards = append(inmemory.Flashcards, flashcard)
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

func (r *FlashcardRepository) Update(authUserID string, flashcard *model.Flashcard) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	// パフォーマンスを考慮して
	// 本番のクエリを1回にするためにリポジトリで認可行う

	var studySet *model.StudySet

	// 学習セットのオーナーを調べる
	for _, s := range inmemory.StudySets {
		if s.ID == flashcard.StudySetID {
			studySet = s
		}
	}
	if studySet == nil {
		return errors.New("studySet doesn't exists")
	}

	// 認可できるか
	if studySet.UserID != authUserID {
		return errors.New("not authorized to update flashcard")
	}

	for i, f := range inmemory.Flashcards {
		if f.ID == flashcard.ID {
			inmemory.Flashcards[i].Question = flashcard.Question
			inmemory.Flashcards[i].Answer = flashcard.Answer
			return nil
		}
	}

	return errors.New("flashcard not found")

}

func (r *FlashcardRepository) Delete(authUserID, flashcardID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	// パフォーマンスを考慮して
	// 本番のクエリを1回にするためにリポジトリで認可行う

	var studySet *model.StudySet

	// フラッシュカード調べてstudysetのid取得
	var studySetID string
	for _, f := range inmemory.Favorites {
		if f.ID == flashcardID {
			studySetID = f.StudySetID
		}
	}

	// 学習セットのオーナーを調べる
	for _, s := range inmemory.StudySets {
		if s.ID == studySetID {
			studySet = s
		}
	}
	if studySet == nil {
		return errors.New("studySet doesn't exists")
	}

	// 認可できるか
	if studySet.UserID != authUserID {
		return errors.New("not authorized to update flashcard")
	}

	// 削除
	for i, flashcardFromDB := range inmemory.Flashcards {
		if flashcardFromDB.ID == flashcardID {
			inmemory.Flashcards = utils.RemoveElementFromSlice(inmemory.Flashcards, i)
			return nil
		}
	}

	return errors.New("flashcard not found")
}
