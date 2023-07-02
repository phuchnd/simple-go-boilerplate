// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	pb "github.com/phuchnd/simple-go-boilerplate/server/grpc/pb"
)

// IGRPCService is an autogenerated mock type for the IGRPCService type
type IGRPCService struct {
	mock.Mock
}

// ListBooks provides a mock function with given fields: _a0, _a1
func (_m *IGRPCService) ListBooks(_a0 context.Context, _a1 *pb.ListBookRequest) (*pb.ListBookResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *pb.ListBookResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pb.ListBookRequest) (*pb.ListBookResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pb.ListBookRequest) *pb.ListBookResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.ListBookResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pb.ListBookRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIGRPCService creates a new instance of IGRPCService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIGRPCService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IGRPCService {
	mock := &IGRPCService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
