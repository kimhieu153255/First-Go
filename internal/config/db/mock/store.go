// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kimhieu153255/first-go/internal/config/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	db "github.com/kimhieu153255/first-go/internal/config/db/sqlc"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// DeleteUserByEmail mocks base method.
func (m *MockStore) DeleteUserByEmail(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserByEmail", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUserByEmail indicates an expected call of DeleteUserByEmail.
func (mr *MockStoreMockRecorder) DeleteUserByEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserByEmail", reflect.TypeOf((*MockStore)(nil).DeleteUserByEmail), arg0, arg1)
}

// DeleteUserByID mocks base method.
func (m *MockStore) DeleteUserByID(arg0 context.Context, arg1 int64) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserByID", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUserByID indicates an expected call of DeleteUserByID.
func (mr *MockStoreMockRecorder) DeleteUserByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserByID", reflect.TypeOf((*MockStore)(nil).DeleteUserByID), arg0, arg1)
}

// GetListUsers mocks base method.
func (m *MockStore) GetListUsers(arg0 context.Context) ([]db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListUsers", arg0)
	ret0, _ := ret[0].([]db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListUsers indicates an expected call of GetListUsers.
func (mr *MockStoreMockRecorder) GetListUsers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListUsers", reflect.TypeOf((*MockStore)(nil).GetListUsers), arg0)
}

// GetUserByEmail mocks base method.
func (m *MockStore) GetUserByEmail(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockStoreMockRecorder) GetUserByEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockStore)(nil).GetUserByEmail), arg0, arg1)
}

// GetUserById mocks base method.
func (m *MockStore) GetUserById(arg0 context.Context, arg1 int64) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockStoreMockRecorder) GetUserById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockStore)(nil).GetUserById), arg0, arg1)
}

// SelectUserForUpdate mocks base method.
func (m *MockStore) SelectUserForUpdate(arg0 context.Context, arg1 int64) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectUserForUpdate", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectUserForUpdate indicates an expected call of SelectUserForUpdate.
func (mr *MockStoreMockRecorder) SelectUserForUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectUserForUpdate", reflect.TypeOf((*MockStore)(nil).SelectUserForUpdate), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockStore) UpdateUser(arg0 context.Context, arg1 db.UpdateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockStoreMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockStore)(nil).UpdateUser), arg0, arg1)
}

// UpdateUserUseStore mocks base method.
func (m *MockStore) UpdateUserUseStore(arg0 context.Context, arg1 db.UpdateUserTxParams) (db.UpdateUserResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserUseStore", arg0, arg1)
	ret0, _ := ret[0].(db.UpdateUserResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserUseStore indicates an expected call of UpdateUserUseStore.
func (mr *MockStoreMockRecorder) UpdateUserUseStore(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserUseStore", reflect.TypeOf((*MockStore)(nil).UpdateUserUseStore), arg0, arg1)
}
