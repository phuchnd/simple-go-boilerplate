// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	entities "github.com/phuchnd/simple-go-boilerplate/internal/service/http/entities"

	mock "github.com/stretchr/testify/mock"
)

// IHTTPService is an autogenerated mock type for the IHTTPService type
type IHTTPService struct {
	mock.Mock
}

// ListBooks provides a mock function with given fields: _a0, _a1
func (_m *IHTTPService) ListBooks(_a0 context.Context, _a1 *entities.ListBookRequest) (*entities.ListBookResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *entities.ListBookResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.ListBookRequest) (*entities.ListBookResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entities.ListBookRequest) *entities.ListBookResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.ListBookResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entities.ListBookRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIHTTPService creates a new instance of IHTTPService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIHTTPService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IHTTPService {
	mock := &IHTTPService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
