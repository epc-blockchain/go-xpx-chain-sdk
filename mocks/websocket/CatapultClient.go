// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import sdk "github.com/epc-blockchain/go-xpx-chain-sdk/sdk"
import subscribers "github.com/epc-blockchain/go-xpx-chain-sdk/sdk/websocket/subscribers"

// CatapultClient is an autogenerated mock type for the CatapultClient type
type CatapultClient struct {
	mock.Mock
}

// AddBlockHandlers provides a mock function with given fields: handlers
func (_m *CatapultClient) AddBlockHandlers(handlers ...subscribers.BlockHandler) error {
	_va := make([]interface{}, len(handlers))
	for _i := range handlers {
		_va[_i] = handlers[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...subscribers.BlockHandler) error); ok {
		r0 = rf(handlers...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddConfirmedAddedHandlers provides a mock function with given fields: address, handlers
func (_m *CatapultClient) AddConfirmedAddedHandlers(address *sdk.Address, handlers ...subscribers.ConfirmedAddedHandler) error {
	_va := make([]interface{}, len(handlers))
	for _i := range handlers {
		_va[_i] = handlers[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, address)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(*sdk.Address, ...subscribers.ConfirmedAddedHandler) error); ok {
		r0 = rf(address, handlers...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddCosignatureHandlers provides a mock function with given fields: address, handlers
func (_m *CatapultClient) AddCosignatureHandlers(address *sdk.Address, handlers ...subscribers.CosignatureHandler) error {
	_va := make([]interface{}, len(handlers))
	for _i := range handlers {
		_va[_i] = handlers[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, address)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(*sdk.Address, ...subscribers.CosignatureHandler) error); ok {
		r0 = rf(address, handlers...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddDriveStateHandlers provides a mock function with given fields: address, handlers
func (_m *CatapultClient) AddDriveStateHandlers(address *sdk.Address, handlers ...subscribers.DriveStateHandler) error {
	_va := make([]interface{}, len(handlers))
	for _i := range handlers {
		_va[_i] = handlers[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, address)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(*sdk.Address, ...subscribers.DriveStateHandler) error); ok {
		r0 = rf(address, handlers...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddPartialAddedHandlers provides a mock function with given fields: address, handlers
func (_m *CatapultClient) AddPartialAddedHandlers(address *sdk.Address, handlers ...subscribers.PartialAddedHandler) error {
	_va := make([]interface{}, len(handlers))
	for _i := range handlers {
		_va[_i] = handlers[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, address)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(*sdk.Address, ...subscribers.PartialAddedHandler) error); ok {
		r0 = rf(address, handlers...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddPartialRemovedHandlers provides a mock function with given fields: address, handlers
func (_m *CatapultClient) AddPartialRemovedHandlers(address *sdk.Address, handlers ...subscribers.PartialRemovedHandler) error {
	_va := make([]interface{}, len(handlers))
	for _i := range handlers {
		_va[_i] = handlers[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, address)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(*sdk.Address, ...subscribers.PartialRemovedHandler) error); ok {
		r0 = rf(address, handlers...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddStatusHandlers provides a mock function with given fields: address, handlers
func (_m *CatapultClient) AddStatusHandlers(address *sdk.Address, handlers ...subscribers.StatusHandler) error {
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

// AddUnconfirmedAddedHandlers provides a mock function with given fields: address, handlers
func (_m *CatapultClient) AddUnconfirmedAddedHandlers(address *sdk.Address, handlers ...subscribers.UnconfirmedAddedHandler) error {
	_va := make([]interface{}, len(handlers))
	for _i := range handlers {
		_va[_i] = handlers[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, address)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(*sdk.Address, ...subscribers.UnconfirmedAddedHandler) error); ok {
		r0 = rf(address, handlers...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddUnconfirmedRemovedHandlers provides a mock function with given fields: address, handlers
func (_m *CatapultClient) AddUnconfirmedRemovedHandlers(address *sdk.Address, handlers ...subscribers.UnconfirmedRemovedHandler) error {
	_va := make([]interface{}, len(handlers))
	for _i := range handlers {
		_va[_i] = handlers[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, address)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(*sdk.Address, ...subscribers.UnconfirmedRemovedHandler) error); ok {
		r0 = rf(address, handlers...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *CatapultClient) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Listen provides a mock function with given fields:
func (_m *CatapultClient) Listen() {
	_m.Called()
}
