package repository

import "go-training/domain/model"

type StudySetRepository interface {
	Create(studySet *model.StudySet) (string, error)
	GetByID(id string) (*model.StudySet, error)
	GetByUserID(userID string) ([]*model.StudySet, error)
	Update(authUserID, studySetID string, studySet *model.StudySet) error
	Delete(authUserID, studySetID string) error
	SearchByTitle(title string) ([]*model.StudySet, error)
}
