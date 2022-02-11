// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	ev "github.com/facebookincubator/symphony/pkg/ev"
	mock "github.com/stretchr/testify/mock"
)

// Emitter is an autogenerated mock type for the Emitter type
type Emitter struct {
	mock.Mock
}

// Emit provides a mock function with given fields: _a0, _a1
func (_m *Emitter) Emit(_a0 context.Context, _a1 *ev.Event) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *ev.Event) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Shutdown provides a mock function with given fields: _a0
func (_m *Emitter) Shutdown(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
