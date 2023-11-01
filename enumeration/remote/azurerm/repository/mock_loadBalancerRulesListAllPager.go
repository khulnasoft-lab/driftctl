// Code generated by mockery v2.28.1. DO NOT EDIT.

package repository

import (
	context "context"

	armnetwork "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"

	mock "github.com/stretchr/testify/mock"
)

// mockLoadBalancerRulesListAllPager is an autogenerated mock type for the loadBalancerRulesListAllPager type
type mockLoadBalancerRulesListAllPager struct {
	mock.Mock
}

// Err provides a mock function with given fields:
func (_m *mockLoadBalancerRulesListAllPager) Err() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NextPage provides a mock function with given fields: ctx
func (_m *mockLoadBalancerRulesListAllPager) NextPage(ctx context.Context) bool {
	ret := _m.Called(ctx)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context) bool); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// PageResponse provides a mock function with given fields:
func (_m *mockLoadBalancerRulesListAllPager) PageResponse() armnetwork.LoadBalancerLoadBalancingRulesListResponse {
	ret := _m.Called()

	var r0 armnetwork.LoadBalancerLoadBalancingRulesListResponse
	if rf, ok := ret.Get(0).(func() armnetwork.LoadBalancerLoadBalancingRulesListResponse); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(armnetwork.LoadBalancerLoadBalancingRulesListResponse)
	}

	return r0
}

type mockConstructorTestingTnewMockLoadBalancerRulesListAllPager interface {
	mock.TestingT
	Cleanup(func())
}

// newMockLoadBalancerRulesListAllPager creates a new instance of mockLoadBalancerRulesListAllPager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMockLoadBalancerRulesListAllPager(t mockConstructorTestingTnewMockLoadBalancerRulesListAllPager) *mockLoadBalancerRulesListAllPager {
	mock := &mockLoadBalancerRulesListAllPager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
