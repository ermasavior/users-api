// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepositoryInterface is a mock of RepositoryInterface interface.
type MockRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryInterfaceMockRecorder
}

// MockRepositoryInterfaceMockRecorder is the mock recorder for MockRepositoryInterface.
type MockRepositoryInterfaceMockRecorder struct {
	mock *MockRepositoryInterface
}

// NewMockRepositoryInterface creates a new mock instance.
func NewMockRepositoryInterface(ctrl *gomock.Controller) *MockRepositoryInterface {
	mock := &MockRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryInterface) EXPECT() *MockRepositoryInterfaceMockRecorder {
	return m.recorder
}

// ComparePasswords mocks base method.
func (m *MockRepositoryInterface) ComparePasswords(hashedPwd, plainPwd string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComparePasswords", hashedPwd, plainPwd)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ComparePasswords indicates an expected call of ComparePasswords.
func (mr *MockRepositoryInterfaceMockRecorder) ComparePasswords(hashedPwd, plainPwd interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComparePasswords", reflect.TypeOf((*MockRepositoryInterface)(nil).ComparePasswords), hashedPwd, plainPwd)
}

// GenerateHashedAndSaltedPassword mocks base method.
func (m *MockRepositoryInterface) GenerateHashedAndSaltedPassword(password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateHashedAndSaltedPassword", password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateHashedAndSaltedPassword indicates an expected call of GenerateHashedAndSaltedPassword.
func (mr *MockRepositoryInterfaceMockRecorder) GenerateHashedAndSaltedPassword(password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateHashedAndSaltedPassword", reflect.TypeOf((*MockRepositoryInterface)(nil).GenerateHashedAndSaltedPassword), password)
}

// GenerateToken mocks base method.
func (m *MockRepositoryInterface) GenerateToken(user User) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", user)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockRepositoryInterfaceMockRecorder) GenerateToken(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockRepositoryInterface)(nil).GenerateToken), user)
}

// GetUserByPhoneNumber mocks base method.
func (m *MockRepositoryInterface) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByPhoneNumber", ctx, phoneNumber)
	ret0, _ := ret[0].(User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByPhoneNumber indicates an expected call of GetUserByPhoneNumber.
func (mr *MockRepositoryInterfaceMockRecorder) GetUserByPhoneNumber(ctx, phoneNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByPhoneNumber", reflect.TypeOf((*MockRepositoryInterface)(nil).GetUserByPhoneNumber), ctx, phoneNumber)
}

// IncrementSuccessLoginCount mocks base method.
func (m *MockRepositoryInterface) IncrementSuccessLoginCount(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncrementSuccessLoginCount", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// IncrementSuccessLoginCount indicates an expected call of IncrementSuccessLoginCount.
func (mr *MockRepositoryInterfaceMockRecorder) IncrementSuccessLoginCount(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrementSuccessLoginCount", reflect.TypeOf((*MockRepositoryInterface)(nil).IncrementSuccessLoginCount), ctx, id)
}

// InsertNewUser mocks base method.
func (m *MockRepositoryInterface) InsertNewUser(ctx context.Context, input User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertNewUser", ctx, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertNewUser indicates an expected call of InsertNewUser.
func (mr *MockRepositoryInterfaceMockRecorder) InsertNewUser(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertNewUser", reflect.TypeOf((*MockRepositoryInterface)(nil).InsertNewUser), ctx, input)
}
