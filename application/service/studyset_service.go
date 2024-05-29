package service

import (
	"go-training/domain/model"
	"go-training/domain/repository"
)

type StudySetService struct {
	repo repository.StudySetRepository
}

func NewStudySetService(repo repository.StudySetRepository) *StudySetService {
	return &StudySetService{
		repo: repo,
	}
}

func (s *StudySetService) CreateStudySet(studySet *model.StudySet) error {
	return s.repo.Create(studySet)
}

func (s *StudySetService) GetStudySetByID(id string) (*model.StudySet, error) {
	return s.repo.GetByID(id)
}

func (s *StudySetService) GetStudySetsByUserID(userID string) ([]*model.StudySet, error) {
	return s.repo.GetByUserID(userID)
}

func (s *StudySetService) UpdateStudySet(studySet *model.StudySet) error {
	return s.repo.Update(studySet)
}

func (s *StudySetService) DeleteStudySet(id string) error {
	return s.repo.Delete(id)
}
