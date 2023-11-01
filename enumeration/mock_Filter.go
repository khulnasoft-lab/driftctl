// Code generated by mockery v2.35.4. DO NOT EDIT.

package enumeration

import (
	resource "github.com/khulnasoft-lab/driftctl/enumeration/resource"
	mock "github.com/stretchr/testify/mock"
)

// MockFilter is an autogenerated mock type for the Filter type
type MockFilter struct {
	mock.Mock
}

// IsResourceIgnored provides a mock function with given fields: res
func (_m *MockFilter) IsResourceIgnored(res *resource.Resource) bool {
	ret := _m.Called(res)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*resource.Resource) bool); ok {
		r0 = rf(res)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsTypeIgnored provides a mock function with given fields: ty
func (_m *MockFilter) IsTypeIgnored(ty resource.ResourceType) bool {
	ret := _m.Called(ty)

	var r0 bool
	if rf, ok := ret.Get(0).(func(resource.ResourceType) bool); ok {
		r0 = rf(ty)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// NewMockFilter creates a new instance of MockFilter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockFilter(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockFilter {
	mock := &MockFilter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
