// Code generated by MockGen. DO NOT EDIT.
// Source: internal/data/pgx/interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/jasonkwh/wex-test/internal/data/model"
)

// MockPurchaseRepository is a mock of PurchaseRepository interface.
type MockPurchaseRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPurchaseRepositoryMockRecorder
}

// MockPurchaseRepositoryMockRecorder is the mock recorder for MockPurchaseRepository.
type MockPurchaseRepositoryMockRecorder struct {
	mock *MockPurchaseRepository
}

// NewMockPurchaseRepository creates a new mock instance.
func NewMockPurchaseRepository(ctrl *gomock.Controller) *MockPurchaseRepository {
	mock := &MockPurchaseRepository{ctrl: ctrl}
	mock.recorder = &MockPurchaseRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPurchaseRepository) EXPECT() *MockPurchaseRepositoryMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockPurchaseRepository) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockPurchaseRepositoryMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockPurchaseRepository)(nil).Close))
}

// GetPurchase mocks base method.
func (m *MockPurchaseRepository) GetPurchase(ctx context.Context, id string) (*model.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPurchase", ctx, id)
	ret0, _ := ret[0].(*model.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPurchase indicates an expected call of GetPurchase.
func (mr *MockPurchaseRepositoryMockRecorder) GetPurchase(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPurchase", reflect.TypeOf((*MockPurchaseRepository)(nil).GetPurchase), ctx, id)
}

// SavePurchase mocks base method.
func (m *MockPurchaseRepository) SavePurchase(ctx context.Context, purchase *model.Transaction) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SavePurchase", ctx, purchase)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SavePurchase indicates an expected call of SavePurchase.
func (mr *MockPurchaseRepositoryMockRecorder) SavePurchase(ctx, purchase interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SavePurchase", reflect.TypeOf((*MockPurchaseRepository)(nil).SavePurchase), ctx, purchase)
}
