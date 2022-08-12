// Code generated by MockGen. DO NOT EDIT.
// Source: application\uploadprocess\service.go

// Package uploadprocess_mock is a generated GoMock package.
package uploadprocess_mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	errorhelper "github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/application/errorhelper"
	uploadprocess "github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/application/uploadprocess"
)

// MockUploadProcessorService is a mock of UploadProcessorService interface.
type MockUploadProcessorService struct {
	ctrl     *gomock.Controller
	recorder *MockUploadProcessorServiceMockRecorder
}

// MockUploadProcessorServiceMockRecorder is the mock recorder for MockUploadProcessorService.
type MockUploadProcessorServiceMockRecorder struct {
	mock *MockUploadProcessorService
}

// NewMockUploadProcessorService creates a new mock instance.
func NewMockUploadProcessorService(ctrl *gomock.Controller) *MockUploadProcessorService {
	mock := &MockUploadProcessorService{ctrl: ctrl}
	mock.recorder = &MockUploadProcessorServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUploadProcessorService) EXPECT() *MockUploadProcessorServiceMockRecorder {
	return m.recorder
}

// GetProcessingStatus mocks base method.
func (m *MockUploadProcessorService) GetProcessingStatus(arg0 *uploadprocess.UploadProcessRequestDTO) (*uploadprocess.UploadProcessResponseDTO, *errorhelper.ErrorResponseDTO) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProcessingStatus", arg0)
	ret0, _ := ret[0].(*uploadprocess.UploadProcessResponseDTO)
	ret1, _ := ret[1].(*errorhelper.ErrorResponseDTO)
	return ret0, ret1
}

// GetProcessingStatus indicates an expected call of GetProcessingStatus.
func (mr *MockUploadProcessorServiceMockRecorder) GetProcessingStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProcessingStatus", reflect.TypeOf((*MockUploadProcessorService)(nil).GetProcessingStatus), arg0)
}

// ProcessUpload mocks base method.
func (m *MockUploadProcessorService) ProcessUpload(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessUpload", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessUpload indicates an expected call of ProcessUpload.
func (mr *MockUploadProcessorServiceMockRecorder) ProcessUpload(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessUpload", reflect.TypeOf((*MockUploadProcessorService)(nil).ProcessUpload), arg0, arg1)
}
