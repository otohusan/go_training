package studySet

import (
	"errors"
	"go-training/domain/model"
	"go-training/domain/repository"
	inmemory "go-training/infrastructure/InMemory"
	"go-training/utils"
	"strings"
	"sync"

	"github.com/google/uuid"
)

type StudySetRepository struct {
	mu sync.Mutex
}

func NewStudySetRepository() repository.StudySetRepository {
	return &StudySetRepository{}
}

func (r *StudySetRepository) Create(studySet *model.StudySet) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// 外部キーのチェック: UserIDが存在するか
	isUserExists := false
	for _, user := range inmemory.Users {
		if user.ID == studySet.UserID {
			isUserExists = true
		}
	}
	if !isUserExists {
		return errors.New("user doesn't exist")
	}

	// uuid作成
	studySet.ID = uuid.New().String()

	inmemory.StudySets = append(inmemory.StudySets, studySet)
	return nil
}

func (r *StudySetRepository) GetByID(id string) (*model.StudySet, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, studySet := range inmemory.StudySets {
		if studySet.ID == id {
			return studySet, nil
		}
	}

	return nil, errors.New("study set not found")
}

func (r *StudySetRepository) GetByUserID(userID string) ([]*model.StudySet, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var userStudySets []*model.StudySet

	for _, studySet := range inmemory.StudySets {
		if studySet.UserID == userID {
			userStudySets = append(userStudySets, studySet)
		}
	}

	return userStudySets, nil
}

func (r *StudySetRepository) Update(authUserID, studySetID string, studySet *model.StudySet) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	// パフォーマンスを考慮して
	// 本番のクエリを1回にするためにリポジトリで認可行う

	var studySetFromDB *model.StudySet
	var targetIndex int

	for i, studySet := range inmemory.StudySets {
		if studySet.ID == studySetID {
			studySetFromDB = studySet
			targetIndex = i
		}
	}

	if studySetFromDB == nil {
		return errors.New("study set not found")
	}

	if studySetFromDB.UserID != authUserID {
		return errors.New("not authorized to update study set")
	}

	// 変更可能な場所のみを変更する
	inmemory.StudySets[targetIndex].Title = studySet.Title
	inmemory.StudySets[targetIndex].Description = studySet.Description
	return nil

}

func (r *StudySetRepository) Delete(authUserID, studySetID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	// パフォーマンスを考慮して
	// 本番のクエリを1回にするためにリポジトリで認可行う

	var studySet *model.StudySet
	var targetIndex int

	for i, studySetFromDB := range inmemory.StudySets {
		if studySetFromDB.ID == studySetID {
			studySet = studySetFromDB
			targetIndex = i
		}
	}

	if studySet == nil {
		return errors.New("study set not found")
	}

	if studySet.UserID != authUserID {
		return errors.New("not authorized to delete study set")
	}

	inmemory.StudySets = utils.RemoveElementFromSlice(inmemory.StudySets, targetIndex)

	return nil

}

func (r *StudySetRepository) SearchByTitle(title string) ([]*model.StudySet, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var results []*model.StudySet
	for _, studySet := range inmemory.StudySets {
		if strings.Contains(strings.ToLower(studySet.Title), strings.ToLower(title)) {
			results = append(results, studySet)
		}
	}
	return results, nil
}
