// Code generated by mockery v2.42.0. DO NOT EDIT.

package domain

import mock "github.com/stretchr/testify/mock"

// MockAuthentication is an autogenerated mock type for the Authentication type
type MockAuthentication[T interface{}, R interface{}] struct {
	mock.Mock
}

type MockAuthentication_Expecter[T interface{}, R interface{}] struct {
	mock *mock.Mock
}

func (_m *MockAuthentication[T, R]) EXPECT() *MockAuthentication_Expecter[T, R] {
	return &MockAuthentication_Expecter[T, R]{mock: &_m.Mock}
}

// Auth provides a mock function with given fields: input
func (_m *MockAuthentication[T, R]) Auth(input T) (R, error) {
	ret := _m.Called(input)

	if len(ret) == 0 {
		panic("no return value specified for Auth")
	}

	var r0 R
	var r1 error
	if rf, ok := ret.Get(0).(func(T) (R, error)); ok {
		return rf(input)
	}
	if rf, ok := ret.Get(0).(func(T) R); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(R)
	}

	if rf, ok := ret.Get(1).(func(T) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAuthentication_Auth_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Auth'
type MockAuthentication_Auth_Call[T interface{}, R interface{}] struct {
	*mock.Call
}

// Auth is a helper method to define mock.On call
//   - input T
func (_e *MockAuthentication_Expecter[T, R]) Auth(input interface{}) *MockAuthentication_Auth_Call[T, R] {
	return &MockAuthentication_Auth_Call[T, R]{Call: _e.mock.On("Auth", input)}
}

func (_c *MockAuthentication_Auth_Call[T, R]) Run(run func(input T)) *MockAuthentication_Auth_Call[T, R] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(T))
	})
	return _c
}

func (_c *MockAuthentication_Auth_Call[T, R]) Return(_a0 R, _a1 error) *MockAuthentication_Auth_Call[T, R] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAuthentication_Auth_Call[T, R]) RunAndReturn(run func(T) (R, error)) *MockAuthentication_Auth_Call[T, R] {
	_c.Call.Return(run)
	return _c
}

// NewMockAuthentication creates a new instance of MockAuthentication. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAuthentication[T interface{}, R interface{}](t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAuthentication[T, R] {
	mock := &MockAuthentication[T, R]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
