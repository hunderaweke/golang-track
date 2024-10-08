// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	domain "testing-api/Domain"

	mock "github.com/stretchr/testify/mock"
)

// UserUsecase is an autogenerated mock type for the UserUsecase type
type UserUsecase struct {
	mock.Mock
}

// Create provides a mock function with given fields: u
func (_m *UserUsecase) Create(u domain.User) (*domain.User, error) {
	ret := _m.Called(u)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(domain.User) (*domain.User, error)); ok {
		return rf(u)
	}
	if rf, ok := ret.Get(0).(func(domain.User) *domain.User); ok {
		r0 = rf(u)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(domain.User) error); ok {
		r1 = rf(u)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: userID
func (_m *UserUsecase) Delete(userID string) error {
	ret := _m.Called(userID)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields:
func (_m *UserUsecase) Get() ([]domain.User, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 []domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]domain.User, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []domain.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByEmail provides a mock function with given fields: email
func (_m *UserUsecase) GetByEmail(email string) (*domain.User, error) {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for GetByEmail")
	}

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.User, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.User); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: userID
func (_m *UserUsecase) GetByID(userID string) (*domain.User, error) {
	ret := _m.Called(userID)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.User, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.User); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: user
func (_m *UserUsecase) Login(user domain.User) (domain.User, error) {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(domain.User) (domain.User, error)); ok {
		return rf(user)
	}
	if rf, ok := ret.Get(0).(func(domain.User) domain.User); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	if rf, ok := ret.Get(1).(func(domain.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PromoteUser provides a mock function with given fields: userID
func (_m *UserUsecase) PromoteUser(userID string) error {
	ret := _m.Called(userID)

	if len(ret) == 0 {
		panic("no return value specified for PromoteUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: userID, data
func (_m *UserUsecase) Update(userID string, data domain.User) (*domain.User, error) {
	ret := _m.Called(userID, data)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string, domain.User) (*domain.User, error)); ok {
		return rf(userID, data)
	}
	if rf, ok := ret.Get(0).(func(string, domain.User) *domain.User); ok {
		r0 = rf(userID, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string, domain.User) error); ok {
		r1 = rf(userID, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserUsecase creates a new instance of UserUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserUsecase {
	mock := &UserUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
