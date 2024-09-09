// Code generated by MockGen. DO NOT EDIT.
// Source: application/service/interface/studyset_service.go
//
// Generated by this command:
//
//	mockgen -source=application/service/interface/studyset_service.go -destination=application/service/mocks/studyset_service.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	model "go-training/domain/model"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockStudySetService is a mock of StudySetService interface.
type MockStudySetService struct {
	ctrl     *gomock.Controller
	recorder *MockStudySetServiceMockRecorder
}

// MockStudySetServiceMockRecorder is the mock recorder for MockStudySetService.
type MockStudySetServiceMockRecorder struct {
	mock *MockStudySetService
}

// NewMockStudySetService creates a new mock instance.
func NewMockStudySetService(ctrl *gomock.Controller) *MockStudySetService {
	mock := &MockStudySetService{ctrl: ctrl}
	mock.recorder = &MockStudySetServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStudySetService) EXPECT() *MockStudySetServiceMockRecorder {
	return m.recorder
}

// CopyStudySetForMe mocks base method.
func (m *MockStudySetService) CopyStudySetForMe(studySet model.StudySet, userID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CopyStudySetForMe", studySet, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// CopyStudySetForMe indicates an expected call of CopyStudySetForMe.
func (mr *MockStudySetServiceMockRecorder) CopyStudySetForMe(studySet, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CopyStudySetForMe", reflect.TypeOf((*MockStudySetService)(nil).CopyStudySetForMe), studySet, userID)
}

// CreateStudySet mocks base method.
func (m *MockStudySetService) CreateStudySet(authUserID string, studySet *model.StudySet) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStudySet", authUserID, studySet)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateStudySet indicates an expected call of CreateStudySet.
func (mr *MockStudySetServiceMockRecorder) CreateStudySet(authUserID, studySet any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStudySet", reflect.TypeOf((*MockStudySetService)(nil).CreateStudySet), authUserID, studySet)
}

// DeleteStudySet mocks base method.
func (m *MockStudySetService) DeleteStudySet(authUserID, studySetID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStudySet", authUserID, studySetID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStudySet indicates an expected call of DeleteStudySet.
func (mr *MockStudySetServiceMockRecorder) DeleteStudySet(authUserID, studySetID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStudySet", reflect.TypeOf((*MockStudySetService)(nil).DeleteStudySet), authUserID, studySetID)
}

// GetStudySetByID mocks base method.
func (m *MockStudySetService) GetStudySetByID(id string) (*model.StudySet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStudySetByID", id)
	ret0, _ := ret[0].(*model.StudySet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStudySetByID indicates an expected call of GetStudySetByID.
func (mr *MockStudySetServiceMockRecorder) GetStudySetByID(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStudySetByID", reflect.TypeOf((*MockStudySetService)(nil).GetStudySetByID), id)
}

// GetStudySetsByUserID mocks base method.
func (m *MockStudySetService) GetStudySetsByUserID(userID string) ([]*model.StudySet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStudySetsByUserID", userID)
	ret0, _ := ret[0].([]*model.StudySet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStudySetsByUserID indicates an expected call of GetStudySetsByUserID.
func (mr *MockStudySetServiceMockRecorder) GetStudySetsByUserID(userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStudySetsByUserID", reflect.TypeOf((*MockStudySetService)(nil).GetStudySetsByUserID), userID)
}

// GetStudySetsWithFlashcardsByUserID mocks base method.
func (m *MockStudySetService) GetStudySetsWithFlashcardsByUserID(userID string) ([]*model.StudySet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStudySetsWithFlashcardsByUserID", userID)
	ret0, _ := ret[0].([]*model.StudySet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStudySetsWithFlashcardsByUserID indicates an expected call of GetStudySetsWithFlashcardsByUserID.
func (mr *MockStudySetServiceMockRecorder) GetStudySetsWithFlashcardsByUserID(userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStudySetsWithFlashcardsByUserID", reflect.TypeOf((*MockStudySetService)(nil).GetStudySetsWithFlashcardsByUserID), userID)
}

// SearchStudySetsByKeyword mocks base method.
func (m *MockStudySetService) SearchStudySetsByKeyword(keyword string) ([]*model.StudySet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchStudySetsByKeyword", keyword)
	ret0, _ := ret[0].([]*model.StudySet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchStudySetsByKeyword indicates an expected call of SearchStudySetsByKeyword.
func (mr *MockStudySetServiceMockRecorder) SearchStudySetsByKeyword(keyword any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchStudySetsByKeyword", reflect.TypeOf((*MockStudySetService)(nil).SearchStudySetsByKeyword), keyword)
}

// UpdateStudySet mocks base method.
func (m *MockStudySetService) UpdateStudySet(authUserID, studySetID string, studySet *model.StudySet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStudySet", authUserID, studySetID, studySet)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStudySet indicates an expected call of UpdateStudySet.
func (mr *MockStudySetServiceMockRecorder) UpdateStudySet(authUserID, studySetID, studySet any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStudySet", reflect.TypeOf((*MockStudySetService)(nil).UpdateStudySet), authUserID, studySetID, studySet)
}
