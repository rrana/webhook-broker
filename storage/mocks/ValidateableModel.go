// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ValidateableModel is an autogenerated mock type for the ValidateableModel type
type ValidateableModel struct {
	mock.Mock
}

// IsInValidState provides a mock function with given fields:
func (_m *ValidateableModel) IsInValidState() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// QuickFix provides a mock function with given fields:
func (_m *ValidateableModel) QuickFix() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
