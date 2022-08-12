// Code generated by MockGen. DO NOT EDIT.
// Source: application\product\service.go

// Package product_mock is a generated GoMock package.
package product_mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	errorhelper "github.com/samuelemwangi/jumia-mds-test/services/products/application/errorhelper"
	product "github.com/samuelemwangi/jumia-mds-test/services/products/application/product"
)

// MockProductService is a mock of ProductService interface.
type MockProductService struct {
	ctrl     *gomock.Controller
	recorder *MockProductServiceMockRecorder
}

// MockProductServiceMockRecorder is the mock recorder for MockProductService.
type MockProductServiceMockRecorder struct {
	mock *MockProductService
}

// NewMockProductService creates a new mock instance.
func NewMockProductService(ctrl *gomock.Controller) *MockProductService {
	mock := &MockProductService{ctrl: ctrl}
	mock.recorder = &MockProductServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductService) EXPECT() *MockProductServiceMockRecorder {
	return m.recorder
}

// GetProductBySKU mocks base method.
func (m *MockProductService) GetProductBySKU(arg0 *product.ProductRequestDTO) (*product.ProductResponseDTO, *errorhelper.ErrorResponseDTO) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductBySKU", arg0)
	ret0, _ := ret[0].(*product.ProductResponseDTO)
	ret1, _ := ret[1].(*errorhelper.ErrorResponseDTO)
	return ret0, ret1
}

// GetProductBySKU indicates an expected call of GetProductBySKU.
func (mr *MockProductServiceMockRecorder) GetProductBySKU(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductBySKU", reflect.TypeOf((*MockProductService)(nil).GetProductBySKU), arg0)
}

// GetProducts mocks base method.
func (m *MockProductService) GetProducts() (*product.ProductsResponseDTO, *errorhelper.ErrorResponseDTO) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProducts")
	ret0, _ := ret[0].(*product.ProductsResponseDTO)
	ret1, _ := ret[1].(*errorhelper.ErrorResponseDTO)
	return ret0, ret1
}

// GetProducts indicates an expected call of GetProducts.
func (mr *MockProductServiceMockRecorder) GetProducts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProducts", reflect.TypeOf((*MockProductService)(nil).GetProducts))
}
