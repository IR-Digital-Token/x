// Code generated by mockery v2.16.0. DO NOT EDIT.

package repomocks

import (
	common "github.com/ethereum/go-ethereum/common"
	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"
)

// Handler is an autogenerated mock type for the Handler type
type Handler struct {
	mock.Mock
}

// HandleTransaction provides a mock function with given fields:
func (_m *Handler) HandleTransaction() func(types.Header, *types.Receipt) error {
	ret := _m.Called()

	var r0 func(types.Header, *types.Receipt) error
	if rf, ok := ret.Get(0).(func() func(types.Header, *types.Receipt) error); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(func(types.Header, *types.Receipt) error)
		}
	}

	return r0
}

// ID provides a mock function with given fields:
func (_m *Handler) ID() common.Hash {
	ret := _m.Called()

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func() common.Hash); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	return r0
}

type mockConstructorTestingTNewHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewHandler creates a new instance of Handler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHandler(t mockConstructorTestingTNewHandler) *Handler {
	mock := &Handler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
