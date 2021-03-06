// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dbsteward/dbsteward/lib/output (interfaces: OutputFileSegmenter)

// Package output is a generated GoMock package.
package output

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockOutputFileSegmenter is a mock of OutputFileSegmenter interface
type MockOutputFileSegmenter struct {
	ctrl     *gomock.Controller
	recorder *MockOutputFileSegmenterMockRecorder
}

// MockOutputFileSegmenterMockRecorder is the mock recorder for MockOutputFileSegmenter
type MockOutputFileSegmenterMockRecorder struct {
	mock *MockOutputFileSegmenter
}

// NewMockOutputFileSegmenter creates a new mock instance
func NewMockOutputFileSegmenter(ctrl *gomock.Controller) *MockOutputFileSegmenter {
	mock := &MockOutputFileSegmenter{ctrl: ctrl}
	mock.recorder = &MockOutputFileSegmenterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOutputFileSegmenter) EXPECT() *MockOutputFileSegmenterMockRecorder {
	return m.recorder
}

// AppendFooter mocks base method
func (m *MockOutputFileSegmenter) AppendFooter(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AppendFooter", varargs...)
}

// AppendFooter indicates an expected call of AppendFooter
func (mr *MockOutputFileSegmenterMockRecorder) AppendFooter(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendFooter", reflect.TypeOf((*MockOutputFileSegmenter)(nil).AppendFooter), varargs...)
}

// AppendHeader mocks base method
func (m *MockOutputFileSegmenter) AppendHeader(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AppendHeader", varargs...)
}

// AppendHeader indicates an expected call of AppendHeader
func (mr *MockOutputFileSegmenterMockRecorder) AppendHeader(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendHeader", reflect.TypeOf((*MockOutputFileSegmenter)(nil).AppendHeader), varargs...)
}

// Close mocks base method
func (m *MockOutputFileSegmenter) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockOutputFileSegmenterMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockOutputFileSegmenter)(nil).Close))
}

// SetHeader mocks base method
func (m *MockOutputFileSegmenter) SetHeader(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "SetHeader", varargs...)
}

// SetHeader indicates an expected call of SetHeader
func (mr *MockOutputFileSegmenterMockRecorder) SetHeader(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockOutputFileSegmenter)(nil).SetHeader), varargs...)
}

// Write mocks base method
func (m *MockOutputFileSegmenter) Write(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Write", varargs...)
}

// Write indicates an expected call of Write
func (mr *MockOutputFileSegmenterMockRecorder) Write(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockOutputFileSegmenter)(nil).Write), varargs...)
}

// WriteSql mocks base method
func (m *MockOutputFileSegmenter) WriteSql(arg0 ...ToSql) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "WriteSql", varargs...)
}

// WriteSql indicates an expected call of WriteSql
func (mr *MockOutputFileSegmenterMockRecorder) WriteSql(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteSql", reflect.TypeOf((*MockOutputFileSegmenter)(nil).WriteSql), arg0...)
}
