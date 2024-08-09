// Code generated by mockery v2.35.4. DO NOT EDIT.

package mocks

import (
	eventstreamapi "github.com/argoproj-labs/argocd-agent/pkg/api/grpc/eventstreamapi"
	eventstreammock "github.com/argoproj-labs/argocd-agent/principal/apis/eventstream/mock"

	mock "github.com/stretchr/testify/mock"
)

// SendHook is an autogenerated mock type for the SendHook type
type SendHook struct {
	mock.Mock
}

// Execute provides a mock function with given fields: s, sub
func (_m *SendHook) Execute(s *eventstreammock.MockEventServer, sub *eventstreamapi.Event) error {
	ret := _m.Called(s, sub)

	var r0 error
	if rf, ok := ret.Get(0).(func(*eventstreammock.MockEventServer, *eventstreamapi.Event) error); ok {
		r0 = rf(s, sub)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewSendHook creates a new instance of SendHook. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSendHook(t interface {
	mock.TestingT
	Cleanup(func())
}) *SendHook {
	mock := &SendHook{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
