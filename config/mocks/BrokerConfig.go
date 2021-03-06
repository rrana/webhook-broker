// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// BrokerConfig is an autogenerated mock type for the BrokerConfig type
type BrokerConfig struct {
	mock.Mock
}

// GetMaxMessageQueueSize provides a mock function with given fields:
func (_m *BrokerConfig) GetMaxMessageQueueSize() uint {
	ret := _m.Called()

	var r0 uint
	if rf, ok := ret.Get(0).(func() uint); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint)
	}

	return r0
}

// GetMaxRetry provides a mock function with given fields:
func (_m *BrokerConfig) GetMaxRetry() uint8 {
	ret := _m.Called()

	var r0 uint8
	if rf, ok := ret.Get(0).(func() uint8); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint8)
	}

	return r0
}

// GetMaxWorkers provides a mock function with given fields:
func (_m *BrokerConfig) GetMaxWorkers() uint {
	ret := _m.Called()

	var r0 uint
	if rf, ok := ret.Get(0).(func() uint); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint)
	}

	return r0
}

// GetRationalDelay provides a mock function with given fields:
func (_m *BrokerConfig) GetRationalDelay() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// GetRetriggerBaseEndpoint provides a mock function with given fields:
func (_m *BrokerConfig) GetRetriggerBaseEndpoint() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetRetryBackoffDelays provides a mock function with given fields:
func (_m *BrokerConfig) GetRetryBackoffDelays() []time.Duration {
	ret := _m.Called()

	var r0 []time.Duration
	if rf, ok := ret.Get(0).(func() []time.Duration); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]time.Duration)
		}
	}

	return r0
}

// IsPriorityDispatcherEnabled provides a mock function with given fields:
func (_m *BrokerConfig) IsPriorityDispatcherEnabled() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsRecoveryWorkersEnabled provides a mock function with given fields:
func (_m *BrokerConfig) IsRecoveryWorkersEnabled() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
