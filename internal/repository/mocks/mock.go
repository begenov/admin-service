// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	domain "github.com/begenov/admin-service/internal/domain"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAdmin is a mock of Admin interface.
type MockAdmin struct {
	ctrl     *gomock.Controller
	recorder *MockAdminMockRecorder
}

// MockAdminMockRecorder is the mock recorder for MockAdmin.
type MockAdminMockRecorder struct {
	mock *MockAdmin
}

// NewMockAdmin creates a new mock instance.
func NewMockAdmin(ctrl *gomock.Controller) *MockAdmin {
	mock := &MockAdmin{ctrl: ctrl}
	mock.recorder = &MockAdminMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdmin) EXPECT() *MockAdminMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAdmin) Create(ctx context.Context, admin domain.Admin) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, admin)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockAdminMockRecorder) Create(ctx, admin interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAdmin)(nil).Create), ctx, admin)
}

// GetByEmail mocks base method.
func (m *MockAdmin) GetByEmail(ctx context.Context, email string) (domain.Admin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByEmail", ctx, email)
	ret0, _ := ret[0].(domain.Admin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByEmail indicates an expected call of GetByEmail.
func (mr *MockAdminMockRecorder) GetByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByEmail", reflect.TypeOf((*MockAdmin)(nil).GetByEmail), ctx, email)
}

// GetByRefresh mocks base method.
func (m *MockAdmin) GetByRefresh(ctx context.Context, refreshToken string) (domain.Admin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByRefresh", ctx, refreshToken)
	ret0, _ := ret[0].(domain.Admin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByRefresh indicates an expected call of GetByRefresh.
func (mr *MockAdminMockRecorder) GetByRefresh(ctx, refreshToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByRefresh", reflect.TypeOf((*MockAdmin)(nil).GetByRefresh), ctx, refreshToken)
}

// SetSession mocks base method.
func (m *MockAdmin) SetSession(ctx context.Context, session domain.Session, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetSession", ctx, session, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetSession indicates an expected call of SetSession.
func (mr *MockAdminMockRecorder) SetSession(ctx, session, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSession", reflect.TypeOf((*MockAdmin)(nil).SetSession), ctx, session, id)
}
