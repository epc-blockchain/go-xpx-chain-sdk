// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import sdk "github.com/proximax-storage/go-xpx-chain-sdk/sdk"

// DriveStateMapper is an autogenerated mock type for the DriveStateMapper type
type DriveStateMapper struct {
	mock.Mock
}

// MapDriveState provides a mock function with given fields: m
func (_m *DriveStateMapper) MapDriveState(m []byte) (*sdk.DriveStateInfo, error) {
	ret := _m.Called(m)

	var r0 *sdk.DriveStateInfo
	if rf, ok := ret.Get(0).(func([]byte) *sdk.DriveStateInfo); ok {
		r0 = rf(m)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sdk.DriveStateInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(m)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
