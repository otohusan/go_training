package service

import "go-training/domain/model"

type StudySetService interface {
	CreateStudySet(authUserID string, studySet *model.StudySet) (string, error)
	GetStudySetByID(id string) (*model.StudySet, error)
	GetStudySetsByUserID(userID string) ([]*model.StudySet, error)
	UpdateStudySet(authUserID, studySetID string, studySet *model.StudySet) error
	DeleteStudySet(authUserID, studySetID string) error
	SearchStudySetsByKeyword(keyword string) ([]*model.StudySet, error)
	GetStudySetsWithFlashcardsByUserID(userID string) ([]*model.StudySet, error)
	CopyStudySetForMe(studySet model.StudySet, userID string) error
}
