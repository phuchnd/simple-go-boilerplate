// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	gorm "gorm.io/gorm"
)

// IMySqlDB is an autogenerated mock type for the IMySqlDB type
type IMySqlDB struct {
	mock.Mock
}

// DB provides a mock function with given fields:
func (_m *IMySqlDB) DB() *gorm.DB {
	ret := _m.Called()

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func() *gorm.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// Ping provides a mock function with given fields:
func (_m *IMySqlDB) Ping() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIMySqlDB creates a new instance of IMySqlDB. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIMySqlDB(t interface {
	mock.TestingT
	Cleanup(func())
}) *IMySqlDB {
	mock := &IMySqlDB{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
