// Code generated by mockery v2.28.1. DO NOT EDIT.

package repository

import (
	context "context"

	armpostgresql "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresql"

	mock "github.com/stretchr/testify/mock"
)

// mockPostgresqlServersClient is an autogenerated mock type for the postgresqlServersClient type
type mockPostgresqlServersClient struct {
	mock.Mock
}

// List provides a mock function with given fields: _a0, _a1
func (_m *mockPostgresqlServersClient) List(_a0 context.Context, _a1 *armpostgresql.ServersListOptions) (armpostgresql.ServersListResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 armpostgresql.ServersListResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *armpostgresql.ServersListOptions) (armpostgresql.ServersListResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *armpostgresql.ServersListOptions) armpostgresql.ServersListResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(armpostgresql.ServersListResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *armpostgresql.ServersListOptions) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTnewMockPostgresqlServersClient interface {
	mock.TestingT
	Cleanup(func())
}

// newMockPostgresqlServersClient creates a new instance of mockPostgresqlServersClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMockPostgresqlServersClient(t mockConstructorTestingTnewMockPostgresqlServersClient) *mockPostgresqlServersClient {
	mock := &mockPostgresqlServersClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
