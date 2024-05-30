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
		return errors.New("password cannot be empty")
	}
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

func (s *StudySetService) CreateStudySet(studySet *model.StudySet) error {
	if err := validateStudySet(studySet); err != nil {
		return err
	}

	return s.repo.Create(studySet)
}

func (s *StudySetService) GetStudySetByID(id string) (*model.StudySet, error) {
	return s.repo.GetByID(id)
}

func (s *StudySetService) GetStudySetsByUserID(userID string) ([]*model.StudySet, error) {
	return s.repo.GetByUserID(userID)
}

func (s *StudySetService) UpdateStudySet(studySet *model.StudySet) error {
	if err := validateStudySet(studySet); err != nil {
		return err
	}

	return s.repo.Update(studySet)
}

func (s *StudySetService) DeleteStudySet(id string) error {
	return s.repo.Delete(id)
}

func (s *StudySetService) SearchStudySetsByTitle(title string) ([]*model.StudySet, error) {
	return s.repo.SearchByTitle(title)
}
