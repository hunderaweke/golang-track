// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "testing-api/Domain"

	mock "github.com/stretchr/testify/mock"
)

// TaskRepository is an autogenerated mock type for the TaskRepository type
type TaskRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: c, t
func (_m *TaskRepository) Create(c context.Context, t domain.Task) (domain.Task, error) {
	ret := _m.Called(c, t)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 domain.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Task) (domain.Task, error)); ok {
		return rf(c, t)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.Task) domain.Task); ok {
		r0 = rf(c, t)
	} else {
		r0 = ret.Get(0).(domain.Task)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.Task) error); ok {
		r1 = rf(c, t)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: c, taskID
func (_m *TaskRepository) Delete(c context.Context, taskID string) error {
	ret := _m.Called(c, taskID)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(c, taskID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: c
func (_m *TaskRepository) Get(c context.Context) ([]domain.Task, error) {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 []domain.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]domain.Task, error)); ok {
		return rf(c)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []domain.Task); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Task)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: c, taskID
func (_m *TaskRepository) GetByID(c context.Context, taskID string) (domain.Task, error) {
	ret := _m.Called(c, taskID)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 domain.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (domain.Task, error)); ok {
		return rf(c, taskID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.Task); ok {
		r0 = rf(c, taskID)
	} else {
		r0 = ret.Get(0).(domain.Task)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(c, taskID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByUserID provides a mock function with given fields: c, userID
func (_m *TaskRepository) GetByUserID(c context.Context, userID string) ([]domain.Task, error) {
	ret := _m.Called(c, userID)

	if len(ret) == 0 {
		panic("no return value specified for GetByUserID")
	}

	var r0 []domain.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]domain.Task, error)); ok {
		return rf(c, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []domain.Task); ok {
		r0 = rf(c, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Task)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(c, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: c, taskID, data
func (_m *TaskRepository) Update(c context.Context, taskID string, data domain.Task) (*domain.Task, error) {
	ret := _m.Called(c, taskID, data)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *domain.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, domain.Task) (*domain.Task, error)); ok {
		return rf(c, taskID, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, domain.Task) *domain.Task); ok {
		r0 = rf(c, taskID, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Task)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, domain.Task) error); ok {
		r1 = rf(c, taskID, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTaskRepository creates a new instance of TaskRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTaskRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *TaskRepository {
	mock := &TaskRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
