// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "testing-api/Domain"

	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: c, u
func (_m *UserRepository) Create(c context.Context, u domain.User) (*domain.User, error) {
	ret := _m.Called(c, u)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.User) (*domain.User, error)); ok {
		return rf(c, u)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.User) *domain.User); ok {
		r0 = rf(c, u)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.User) error); ok {
		r1 = rf(c, u)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: c, userID
func (_m *UserRepository) Delete(c context.Context, userID string) error {
	ret := _m.Called(c, userID)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(c, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: c
func (_m *UserRepository) Get(c context.Context) ([]domain.User, error) {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 []domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]domain.User, error)); ok {
		return rf(c)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []domain.User); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByEmail provides a mock function with given fields: c, email
func (_m *UserRepository) GetByEmail(c context.Context, email string) (*domain.User, error) {
	ret := _m.Called(c, email)

	if len(ret) == 0 {
		panic("no return value specified for GetByEmail")
	}

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*domain.User, error)); ok {
		return rf(c, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.User); ok {
		r0 = rf(c, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(c, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: c, userID
func (_m *UserRepository) GetByID(c context.Context, userID string) (*domain.User, error) {
	ret := _m.Called(c, userID)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*domain.User, error)); ok {
		return rf(c, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.User); ok {
		r0 = rf(c, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(c, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PromoteUser provides a mock function with given fields: c, userID
func (_m *UserRepository) PromoteUser(c context.Context, userID string) error {
	ret := _m.Called(c, userID)

	if len(ret) == 0 {
		panic("no return value specified for PromoteUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(c, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: c, userID, data
func (_m *UserRepository) Update(c context.Context, userID string, data domain.User) (*domain.User, error) {
	ret := _m.Called(c, userID, data)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, domain.User) (*domain.User, error)); ok {
		return rf(c, userID, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, domain.User) *domain.User); ok {
		r0 = rf(c, userID, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, domain.User) error); ok {
		r1 = rf(c, userID, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}