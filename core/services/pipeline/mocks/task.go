// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	logger "github.com/GoPlugin/pluginV2/core/logger"
	mock "github.com/stretchr/testify/mock"

	pipeline "github.com/GoPlugin/pluginV2/core/services/pipeline"

	time "time"
)

// Task is an autogenerated mock type for the Task type
type Task struct {
	mock.Mock
}

// Base provides a mock function with given fields:
func (_m *Task) Base() *pipeline.BaseTask {
	ret := _m.Called()

	var r0 *pipeline.BaseTask
	if rf, ok := ret.Get(0).(func() *pipeline.BaseTask); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pipeline.BaseTask)
		}
	}

	return r0
}

// DotID provides a mock function with given fields:
func (_m *Task) DotID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ID provides a mock function with given fields:
func (_m *Task) ID() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Inputs provides a mock function with given fields:
func (_m *Task) Inputs() []pipeline.TaskDependency {
	ret := _m.Called()

	var r0 []pipeline.TaskDependency
	if rf, ok := ret.Get(0).(func() []pipeline.TaskDependency); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]pipeline.TaskDependency)
		}
	}

	return r0
}

// OutputIndex provides a mock function with given fields:
func (_m *Task) OutputIndex() int32 {
	ret := _m.Called()

	var r0 int32
	if rf, ok := ret.Get(0).(func() int32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int32)
	}

	return r0
}

// Outputs provides a mock function with given fields:
func (_m *Task) Outputs() []pipeline.Task {
	ret := _m.Called()

	var r0 []pipeline.Task
	if rf, ok := ret.Get(0).(func() []pipeline.Task); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]pipeline.Task)
		}
	}

	return r0
}

// Run provides a mock function with given fields: ctx, lggr, vars, inputs
func (_m *Task) Run(ctx context.Context, lggr logger.Logger, vars pipeline.Vars, inputs []pipeline.Result) (pipeline.Result, pipeline.RunInfo) {
	ret := _m.Called(ctx, lggr, vars, inputs)

	var r0 pipeline.Result
	var r1 pipeline.RunInfo
	if rf, ok := ret.Get(0).(func(context.Context, logger.Logger, pipeline.Vars, []pipeline.Result) (pipeline.Result, pipeline.RunInfo)); ok {
		return rf(ctx, lggr, vars, inputs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, logger.Logger, pipeline.Vars, []pipeline.Result) pipeline.Result); ok {
		r0 = rf(ctx, lggr, vars, inputs)
	} else {
		r0 = ret.Get(0).(pipeline.Result)
	}

	if rf, ok := ret.Get(1).(func(context.Context, logger.Logger, pipeline.Vars, []pipeline.Result) pipeline.RunInfo); ok {
		r1 = rf(ctx, lggr, vars, inputs)
	} else {
		r1 = ret.Get(1).(pipeline.RunInfo)
	}

	return r0, r1
}

// TaskMaxBackoff provides a mock function with given fields:
func (_m *Task) TaskMaxBackoff() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// TaskMinBackoff provides a mock function with given fields:
func (_m *Task) TaskMinBackoff() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// TaskRetries provides a mock function with given fields:
func (_m *Task) TaskRetries() uint32 {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

// TaskTimeout provides a mock function with given fields:
func (_m *Task) TaskTimeout() (time.Duration, bool) {
	ret := _m.Called()

	var r0 time.Duration
	var r1 bool
	if rf, ok := ret.Get(0).(func() (time.Duration, bool)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	if rf, ok := ret.Get(1).(func() bool); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// Type provides a mock function with given fields:
func (_m *Task) Type() pipeline.TaskType {
	ret := _m.Called()

	var r0 pipeline.TaskType
	if rf, ok := ret.Get(0).(func() pipeline.TaskType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(pipeline.TaskType)
	}

	return r0
}

type mockConstructorTestingTNewTask interface {
	mock.TestingT
	Cleanup(func())
}

// NewTask creates a new instance of Task. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTask(t mockConstructorTestingTNewTask) *Task {
	mock := &Task{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
