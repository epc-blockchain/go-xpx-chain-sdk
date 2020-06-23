// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import sdk "github.com/epc-blockchain/go-xpx-chain-sdk/sdk"
import subscribers "github.com/epc-blockchain/go-xpx-chain-sdk/sdk/websocket/subscribers"

// Status is an autogenerated mock type for the Status type
type Status struct {
	mock.Mock
}

// AddHandlers provides a mock function with given fields: address, handlers
func (_m *Status) AddHandlers(address *sdk.Address, handlers ...subscribers.StatusHandler) error {
	_va := make([]interface{}, len(handlers))
	for _i := range handlers {
		_va[_i] = handlers[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, address)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(*sdk.Address, ...subscribers.StatusHandler) error); ok {
		r0 = rf(address, handlers...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAddresses provides a mock function with given fields:
func (_m *Status) GetAddresses() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// GetHandlers provides a mock function with given fields: address
func (_m *Status) GetHandlers(address *sdk.Address) []*subscribers.StatusHandler {
	ret := _m.Called(address)

	var r0 []*subscribers.StatusHandler
	if rf, ok := ret.Get(0).(func(*sdk.Address) []*subscribers.StatusHandler); ok {
		r0 = rf(address)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*subscribers.StatusHandler)
		}
	}

	return r0
}

// HasHandlers provides a mock function with given fields: address
func (_m *Status) HasHandlers(address *sdk.Address) bool {
	ret := _m.Called(address)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*sdk.Address) bool); ok {
		r0 = rf(address)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// RemoveHandlers provides a mock function with given fields: address, handlers
func (_m *Status) RemoveHandlers(address *sdk.Address, handlers ...*subscribers.StatusHandler) bool {
	_va := make([]interface{}, len(handlers))
	for _i := range handlers {
		_va[_i] = handlers[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, address)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*sdk.Address, ...*subscribers.StatusHandler) bool); ok {
		r0 = rf(address, handlers...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
