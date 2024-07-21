package service

import (
	"errors"
	"go-training/domain/model"
	"go-training/domain/repository"
)

type StudySetService struct {
	repo          repository.StudySetRepository
	flashcardRepo repository.FlashcardRepository
}

func NewStudySetService(repo repository.StudySetRepository, flashcardRepo repository.FlashcardRepository) *StudySetService {
	return &StudySetService{
		repo:          repo,
		flashcardRepo: flashcardRepo,
	}
}

func (s *StudySetService) CreateStudySet(authUserID string, studySet *model.StudySet) (string, error) {
	if studySet.Title == "" {
		return "", errors.New("title cannot be empty")
	}
	if studySet.Description == "" {
		return "", errors.New("description cannot be empty")
	}

	return s.repo.Create(studySet)
}

func (s *StudySetService) GetStudySetByID(id string) (*model.StudySet, error) {
	return s.repo.GetByID(id)
}

func (s *StudySetService) GetStudySetsByUserID(userID string) ([]*model.StudySet, error) {
	return s.repo.GetByUserID(userID)
}

func (s *StudySetService) UpdateStudySet(authUserID, studySetID string, studySet *model.StudySet) error {
	if studySet.Title == "" {
		return errors.New("title cannot be empty")
	}
	if studySet.Description == "" {
		return errors.New("description cannot be empty")
	}

	// パフォーマンスを考慮してリポジトリに認可を移譲
	return s.repo.Update(authUserID, studySetID, studySet)
}

func (s *StudySetService) DeleteStudySet(authUserID, studySetID string) error {
	// パフォーマンスを考慮してリポジトリに認可を移譲
	return s.repo.Delete(authUserID, studySetID)
}

func (s *StudySetService) SearchStudySetsByKeyword(keyword string) ([]*model.StudySet, error) {
	studySets, err := s.repo.SearchByKeyword(keyword)
	if err != nil {
		return nil, err
	}

	//NOTICE: クエリをユーザの学習セット分行うから効率悪い
	// もっと良い方法ありそう
	for _, studySet := range studySets {
		flashcards, err := s.flashcardRepo.GetByStudySetID(studySet.ID)
		if err != nil {
			return nil, err
		}
		studySet.Flashcards = flashcards
	}

	return studySets, nil
}

// flashCardも含めて学習セットを返す
func (s *StudySetService) GetStudySetsWithFlashcardsByUserID(userID string) ([]*model.StudySet, error) {
	studySets, err := s.repo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	//NOTICE: クエリをユーザの学習セット分行うから効率悪い
	// もっと良い方法ありそう
	for _, studySet := range studySets {
		flashcards, err := s.flashcardRepo.GetByStudySetID(studySet.ID)
		if err != nil {
			return nil, err
		}
		studySet.Flashcards = flashcards
	}

	return studySets, nil
}

// studySetをコピーして自分のにする
func (s *StudySetService) CopyStudySetForMe(studySet model.StudySet, userID string) error {
	flashcards, err := s.flashcardRepo.GetByStudySetID(studySet.ID)
	if err != nil {
		return err
	}

	var newStudySet model.StudySet

	newStudySet.Title = studySet.Title + "のコピー"
	newStudySet.Description = studySet.Description
	newStudySet.UserID = userID

	newStudySetID, err := s.repo.Create(&newStudySet)
	if err != nil {
		return err
	}

	for _, f := range flashcards {
		f.StudySetID = newStudySetID
		_, err := s.flashcardRepo.Create(userID, f)

		if err != nil {
			return err
		}
	}

	return nil
}
