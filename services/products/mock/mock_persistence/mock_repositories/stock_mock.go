// Code generated by MockGen. DO NOT EDIT.
// Source: persistence\repositories\stock.go

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/samuelemwangi/jumia-mds-test/services/products/domain"
)

// MockStockRepository is a mock of StockRepository interface.
type MockStockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockStockRepositoryMockRecorder
}

// MockStockRepositoryMockRecorder is the mock recorder for MockStockRepository.
type MockStockRepositoryMockRecorder struct {
	mock *MockStockRepository
}

// NewMockStockRepository creates a new mock instance.
func NewMockStockRepository(ctrl *gomock.Controller) *MockStockRepository {
	mock := &MockStockRepository{ctrl: ctrl}
	mock.recorder = &MockStockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStockRepository) EXPECT() *MockStockRepositoryMockRecorder {
	return m.recorder
}

// GetStockByProductAndCountry mocks base method.
func (m *MockStockRepository) GetStockByProductAndCountry(arg0 *domain.Stock) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStockByProductAndCountry", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetStockByProductAndCountry indicates an expected call of GetStockByProductAndCountry.
func (mr *MockStockRepositoryMockRecorder) GetStockByProductAndCountry(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStockByProductAndCountry", reflect.TypeOf((*MockStockRepository)(nil).GetStockByProductAndCountry), arg0)
}

// SaveStock mocks base method.
func (m *MockStockRepository) SaveStock(arg0 *domain.Stock) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveStock", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveStock indicates an expected call of SaveStock.
func (mr *MockStockRepositoryMockRecorder) SaveStock(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveStock", reflect.TypeOf((*MockStockRepository)(nil).SaveStock), arg0)
}

// UpdateStock mocks base method.
func (m *MockStockRepository) UpdateStock(arg0 *domain.Stock) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStock", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStock indicates an expected call of UpdateStock.
func (mr *MockStockRepositoryMockRecorder) UpdateStock(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStock", reflect.TypeOf((*MockStockRepository)(nil).UpdateStock), arg0)
}