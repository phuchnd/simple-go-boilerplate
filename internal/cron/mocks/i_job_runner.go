// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	cron "github.com/phuchnd/simple-go-boilerplate/internal/cron"
	mock "github.com/stretchr/testify/mock"
)

// IJobRunner is an autogenerated mock type for the IJobRunner type
type IJobRunner struct {
	mock.Mock
}

// RegisterJob provides a mock function with given fields: job
func (_m *IJobRunner) RegisterJob(job cron.IJob) error {
	ret := _m.Called(job)

	var r0 error
	if rf, ok := ret.Get(0).(func(cron.IJob) error); ok {
		r0 = rf(job)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIJobRunner creates a new instance of IJobRunner. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIJobRunner(t interface {
	mock.TestingT
	Cleanup(func())
}) *IJobRunner {
	mock := &IJobRunner{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
