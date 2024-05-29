// Code generated by MockGen. DO NOT EDIT.
// Source: log/slog (interfaces: Handler)
//
// Generated by this command:
//
//	mockgen -typed -package=mocks -destination=./mocks/mocks.go log/slog Handler
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	slog "log/slog"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockHandler is a mock of Handler interface.
type MockHandler struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerMockRecorder
}

// MockHandlerMockRecorder is the mock recorder for MockHandler.
type MockHandlerMockRecorder struct {
	mock *MockHandler
}

// NewMockHandler creates a new mock instance.
func NewMockHandler(ctrl *gomock.Controller) *MockHandler {
	mock := &MockHandler{ctrl: ctrl}
	mock.recorder = &MockHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHandler) EXPECT() *MockHandlerMockRecorder {
	return m.recorder
}

// Enabled mocks base method.
func (m *MockHandler) Enabled(arg0 context.Context, arg1 slog.Level) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enabled", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Enabled indicates an expected call of Enabled.
func (mr *MockHandlerMockRecorder) Enabled(arg0, arg1 any) *MockHandlerEnabledCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enabled", reflect.TypeOf((*MockHandler)(nil).Enabled), arg0, arg1)
	return &MockHandlerEnabledCall{Call: call}
}

// MockHandlerEnabledCall wrap *gomock.Call
type MockHandlerEnabledCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockHandlerEnabledCall) Return(arg0 bool) *MockHandlerEnabledCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockHandlerEnabledCall) Do(f func(context.Context, slog.Level) bool) *MockHandlerEnabledCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockHandlerEnabledCall) DoAndReturn(f func(context.Context, slog.Level) bool) *MockHandlerEnabledCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Handle mocks base method.
func (m *MockHandler) Handle(arg0 context.Context, arg1 slog.Record) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handle", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Handle indicates an expected call of Handle.
func (mr *MockHandlerMockRecorder) Handle(arg0, arg1 any) *MockHandlerHandleCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockHandler)(nil).Handle), arg0, arg1)
	return &MockHandlerHandleCall{Call: call}
}

// MockHandlerHandleCall wrap *gomock.Call
type MockHandlerHandleCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockHandlerHandleCall) Return(arg0 error) *MockHandlerHandleCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockHandlerHandleCall) Do(f func(context.Context, slog.Record) error) *MockHandlerHandleCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockHandlerHandleCall) DoAndReturn(f func(context.Context, slog.Record) error) *MockHandlerHandleCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// WithAttrs mocks base method.
func (m *MockHandler) WithAttrs(arg0 []slog.Attr) slog.Handler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithAttrs", arg0)
	ret0, _ := ret[0].(slog.Handler)
	return ret0
}

// WithAttrs indicates an expected call of WithAttrs.
func (mr *MockHandlerMockRecorder) WithAttrs(arg0 any) *MockHandlerWithAttrsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithAttrs", reflect.TypeOf((*MockHandler)(nil).WithAttrs), arg0)
	return &MockHandlerWithAttrsCall{Call: call}
}

// MockHandlerWithAttrsCall wrap *gomock.Call
type MockHandlerWithAttrsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockHandlerWithAttrsCall) Return(arg0 slog.Handler) *MockHandlerWithAttrsCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockHandlerWithAttrsCall) Do(f func([]slog.Attr) slog.Handler) *MockHandlerWithAttrsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockHandlerWithAttrsCall) DoAndReturn(f func([]slog.Attr) slog.Handler) *MockHandlerWithAttrsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// WithGroup mocks base method.
func (m *MockHandler) WithGroup(arg0 string) slog.Handler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithGroup", arg0)
	ret0, _ := ret[0].(slog.Handler)
	return ret0
}

// WithGroup indicates an expected call of WithGroup.
func (mr *MockHandlerMockRecorder) WithGroup(arg0 any) *MockHandlerWithGroupCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithGroup", reflect.TypeOf((*MockHandler)(nil).WithGroup), arg0)
	return &MockHandlerWithGroupCall{Call: call}
}

// MockHandlerWithGroupCall wrap *gomock.Call
type MockHandlerWithGroupCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockHandlerWithGroupCall) Return(arg0 slog.Handler) *MockHandlerWithGroupCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockHandlerWithGroupCall) Do(f func(string) slog.Handler) *MockHandlerWithGroupCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockHandlerWithGroupCall) DoAndReturn(f func(string) slog.Handler) *MockHandlerWithGroupCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
