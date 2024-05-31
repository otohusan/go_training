package service

import (
	"errors"
	"go-training/domain/model"
	"go-training/domain/repository"
)

// StudySetが適切に値を持ってるか確かめる関数
func validateStudySet(studySet *model.StudySet) error {
	if studySet.Title == "" {
		return errors.New("title cannot be empty")
	}
	if studySet.Description == "" {
		return errors.New("description cannot be empty")
	}
	// userIDはJWTから取得するようにするか悩み中
	if studySet.UserID == "" {
		return errors.New("UserID cannot be empty")
	}
	return nil
}

type StudySetService struct {
	repo repository.StudySetRepository
}

func NewStudySetService(repo repository.StudySetRepository) *StudySetService {
	return &StudySetService{
		repo: repo,
	}
}

func (s *StudySetService) CreateStudySet(authUserID string, studySet *model.StudySet) error {
	if err := validateStudySet(studySet); err != nil {
		return err
	}

	if authUserID != studySet.UserID {
		return errors.New("not authorized to create study set")
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

func (s *StudySetService) SearchStudySetsByTitle(title string) ([]*model.StudySet, error) {
	return s.repo.SearchByTitle(title)
}
