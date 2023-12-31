// Code generated by mockery v2.28.1. DO NOT EDIT.

package repository

import (
	autoscaling "github.com/aws/aws-sdk-go/service/autoscaling"
	mock "github.com/stretchr/testify/mock"
)

// MockAutoScalingRepository is an autogenerated mock type for the AutoScalingRepository type
type MockAutoScalingRepository struct {
	mock.Mock
}

// DescribeLaunchConfigurations provides a mock function with given fields:
func (_m *MockAutoScalingRepository) DescribeLaunchConfigurations() ([]*autoscaling.LaunchConfiguration, error) {
	ret := _m.Called()

	var r0 []*autoscaling.LaunchConfiguration
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*autoscaling.LaunchConfiguration, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*autoscaling.LaunchConfiguration); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*autoscaling.LaunchConfiguration)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockAutoScalingRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockAutoScalingRepository creates a new instance of MockAutoScalingRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockAutoScalingRepository(t mockConstructorTestingTNewMockAutoScalingRepository) *MockAutoScalingRepository {
	mock := &MockAutoScalingRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
