// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go
//
// Generated by this command:
//
//	mockgen -source=repository.go -destination=mocks/repository.go
//

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"
	models "server/models"

	gomock "go.uber.org/mock/gomock"
)

// MockApi is a mock of Api interface.
type MockApi struct {
	ctrl     *gomock.Controller
	recorder *MockApiMockRecorder
}

// MockApiMockRecorder is the mock recorder for MockApi.
type MockApiMockRecorder struct {
	mock *MockApi
}

// NewMockApi creates a new mock instance.
func NewMockApi(ctrl *gomock.Controller) *MockApi {
	mock := &MockApi{ctrl: ctrl}
	mock.recorder = &MockApiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApi) EXPECT() *MockApiMockRecorder {
	return m.recorder
}

// GetHistory mocks base method.
func (m *MockApi) GetHistory(id string) ([]models.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHistory", id)
	ret0, _ := ret[0].([]models.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHistory indicates an expected call of GetHistory.
func (mr *MockApiMockRecorder) GetHistory(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHistory", reflect.TypeOf((*MockApi)(nil).GetHistory), id)
}

// GetWallet mocks base method.
func (m *MockApi) GetWallet(id string) (models.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWallet", id)
	ret0, _ := ret[0].(models.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWallet indicates an expected call of GetWallet.
func (mr *MockApiMockRecorder) GetWallet(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWallet", reflect.TypeOf((*MockApi)(nil).GetWallet), id)
}

// NewWallet mocks base method.
func (m *MockApi) NewWallet() (models.Wallet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewWallet")
	ret0, _ := ret[0].(models.Wallet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewWallet indicates an expected call of NewWallet.
func (mr *MockApiMockRecorder) NewWallet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewWallet", reflect.TypeOf((*MockApi)(nil).NewWallet))
}

// SendMoney mocks base method.
func (m *MockApi) SendMoney(fromId string, input models.SendMoneyInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMoney", fromId, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMoney indicates an expected call of SendMoney.
func (mr *MockApiMockRecorder) SendMoney(fromId, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMoney", reflect.TypeOf((*MockApi)(nil).SendMoney), fromId, input)
}
