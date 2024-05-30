package repository

import "go-training/domain/model"

type StudySetRepository interface {
	Create(studySet *model.StudySet) error
	GetByID(id string) (*model.StudySet, error)
	GetByUserID(userID string) ([]*model.StudySet, error)
	Update(studySet *model.StudySet) error
	Delete(id string) error
	SearchByTitle(title string) ([]*model.StudySet, error)
}
