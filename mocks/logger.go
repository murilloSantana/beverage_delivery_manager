// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Logger is an autogenerated mock type for the Logger type
type Logger struct {
	mock.Mock
}

// Error provides a mock function with given fields: values, msg
func (_m *Logger) Error(values map[string]interface{}, msg interface{}) {
	_m.Called(values, msg)
}

// Info provides a mock function with given fields: values, msg
func (_m *Logger) Info(values map[string]interface{}, msg interface{}) {
	_m.Called(values, msg)
}
