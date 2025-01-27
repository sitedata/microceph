// Code generated by mockery v2.30.10. DO NOT EDIT.

package mocks

import (
	common "github.com/canonical/microceph/microceph/common"
	mock "github.com/stretchr/testify/mock"
)

// PlacementIntf is an autogenerated mock type for the PlacementIntf type
type PlacementIntf struct {
	mock.Mock
}

// DbUpdate provides a mock function with given fields: _a0
func (_m *PlacementIntf) DbUpdate(_a0 common.StateInterface) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(common.StateInterface) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HospitalityCheck provides a mock function with given fields: _a0
func (_m *PlacementIntf) HospitalityCheck(_a0 common.StateInterface) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(common.StateInterface) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PopulateParams provides a mock function with given fields: _a0, _a1
func (_m *PlacementIntf) PopulateParams(_a0 common.StateInterface, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(common.StateInterface, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PostPlacementCheck provides a mock function with given fields: _a0
func (_m *PlacementIntf) PostPlacementCheck(_a0 common.StateInterface) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(common.StateInterface) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ServiceInit provides a mock function with given fields: _a0
func (_m *PlacementIntf) ServiceInit(_a0 common.StateInterface) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(common.StateInterface) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewPlacementIntf creates a new instance of PlacementIntf. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPlacementIntf(t interface {
	mock.TestingT
	Cleanup(func())
}) *PlacementIntf {
	mock := &PlacementIntf{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
