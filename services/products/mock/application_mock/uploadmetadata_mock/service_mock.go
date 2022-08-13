// Code generated by MockGen. DO NOT EDIT.
// Source: application\uploadmetadata\service.go

// Package uploadmetadata_mock is a generated GoMock package.
package uploadmetadata_mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	errorhelper "github.com/samuelemwangi/jumia-mds-test/services/products/application/errorhelper"
	uploadmetadata "github.com/samuelemwangi/jumia-mds-test/services/products/application/uploadmetadata"
)

// MockUploadMetadataService is a mock of UploadMetadataService interface.
type MockUploadMetadataService struct {
	ctrl     *gomock.Controller
	recorder *MockUploadMetadataServiceMockRecorder
}

// MockUploadMetadataServiceMockRecorder is the mock recorder for MockUploadMetadataService.
type MockUploadMetadataServiceMockRecorder struct {
	mock *MockUploadMetadataService
}

// NewMockUploadMetadataService creates a new mock instance.
func NewMockUploadMetadataService(ctrl *gomock.Controller) *MockUploadMetadataService {
	mock := &MockUploadMetadataService{ctrl: ctrl}
	mock.recorder = &MockUploadMetadataServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUploadMetadataService) EXPECT() *MockUploadMetadataServiceMockRecorder {
	return m.recorder
}

// SaveUploadMetadaData mocks base method.
func (m *MockUploadMetadataService) SaveUploadMetadaData(arg0 *uploadmetadata.UploadMetadataDTO) (*uploadmetadata.UploadResponseDTO, *errorhelper.ErrorResponseDTO) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveUploadMetadaData", arg0)
	ret0, _ := ret[0].(*uploadmetadata.UploadResponseDTO)
	ret1, _ := ret[1].(*errorhelper.ErrorResponseDTO)
	return ret0, ret1
}

// SaveUploadMetadaData indicates an expected call of SaveUploadMetadaData.
func (mr *MockUploadMetadataServiceMockRecorder) SaveUploadMetadaData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveUploadMetadaData", reflect.TypeOf((*MockUploadMetadataService)(nil).SaveUploadMetadaData), arg0)
}