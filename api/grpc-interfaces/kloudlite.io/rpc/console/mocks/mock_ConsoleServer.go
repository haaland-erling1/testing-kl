// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	context "context"

	console "github.com/kloudlite/api/grpc-interfaces/kloudlite.io/rpc/console"

	mock "github.com/stretchr/testify/mock"
)

// console_server is an autogenerated mock type for the ConsoleServer type
type console_server struct {
	mock.Mock
}

type console_server_Expecter struct {
	mock *mock.Mock
}

func (_m *console_server) EXPECT() *console_server_Expecter {
	return &console_server_Expecter{mock: &_m.Mock}
}

// GetApp provides a mock function with given fields: _a0, _a1
func (_m *console_server) GetApp(_a0 context.Context, _a1 *console.AppIn) (*console.AppOut, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *console.AppOut
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *console.AppIn) (*console.AppOut, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *console.AppIn) *console.AppOut); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*console.AppOut)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *console.AppIn) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// console_server_GetApp_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetApp'
type console_server_GetApp_Call struct {
	*mock.Call
}

// GetApp is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *console.AppIn
func (_e *console_server_Expecter) GetApp(_a0 interface{}, _a1 interface{}) *console_server_GetApp_Call {
	return &console_server_GetApp_Call{Call: _e.mock.On("GetApp", _a0, _a1)}
}

func (_c *console_server_GetApp_Call) Run(run func(_a0 context.Context, _a1 *console.AppIn)) *console_server_GetApp_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*console.AppIn))
	})
	return _c
}

func (_c *console_server_GetApp_Call) Return(_a0 *console.AppOut, _a1 error) *console_server_GetApp_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *console_server_GetApp_Call) RunAndReturn(run func(context.Context, *console.AppIn) (*console.AppOut, error)) *console_server_GetApp_Call {
	_c.Call.Return(run)
	return _c
}

// GetManagedSvc provides a mock function with given fields: _a0, _a1
func (_m *console_server) GetManagedSvc(_a0 context.Context, _a1 *console.MSvcIn) (*console.MSvcOut, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *console.MSvcOut
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *console.MSvcIn) (*console.MSvcOut, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *console.MSvcIn) *console.MSvcOut); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*console.MSvcOut)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *console.MSvcIn) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// console_server_GetManagedSvc_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetManagedSvc'
type console_server_GetManagedSvc_Call struct {
	*mock.Call
}

// GetManagedSvc is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *console.MSvcIn
func (_e *console_server_Expecter) GetManagedSvc(_a0 interface{}, _a1 interface{}) *console_server_GetManagedSvc_Call {
	return &console_server_GetManagedSvc_Call{Call: _e.mock.On("GetManagedSvc", _a0, _a1)}
}

func (_c *console_server_GetManagedSvc_Call) Run(run func(_a0 context.Context, _a1 *console.MSvcIn)) *console_server_GetManagedSvc_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*console.MSvcIn))
	})
	return _c
}

func (_c *console_server_GetManagedSvc_Call) Return(_a0 *console.MSvcOut, _a1 error) *console_server_GetManagedSvc_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *console_server_GetManagedSvc_Call) RunAndReturn(run func(context.Context, *console.MSvcIn) (*console.MSvcOut, error)) *console_server_GetManagedSvc_Call {
	_c.Call.Return(run)
	return _c
}

// GetProjectName provides a mock function with given fields: _a0, _a1
func (_m *console_server) GetProjectName(_a0 context.Context, _a1 *console.ProjectIn) (*console.ProjectOut, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *console.ProjectOut
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *console.ProjectIn) (*console.ProjectOut, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *console.ProjectIn) *console.ProjectOut); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*console.ProjectOut)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *console.ProjectIn) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// console_server_GetProjectName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProjectName'
type console_server_GetProjectName_Call struct {
	*mock.Call
}

// GetProjectName is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *console.ProjectIn
func (_e *console_server_Expecter) GetProjectName(_a0 interface{}, _a1 interface{}) *console_server_GetProjectName_Call {
	return &console_server_GetProjectName_Call{Call: _e.mock.On("GetProjectName", _a0, _a1)}
}

func (_c *console_server_GetProjectName_Call) Run(run func(_a0 context.Context, _a1 *console.ProjectIn)) *console_server_GetProjectName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*console.ProjectIn))
	})
	return _c
}

func (_c *console_server_GetProjectName_Call) Return(_a0 *console.ProjectOut, _a1 error) *console_server_GetProjectName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *console_server_GetProjectName_Call) RunAndReturn(run func(context.Context, *console.ProjectIn) (*console.ProjectOut, error)) *console_server_GetProjectName_Call {
	_c.Call.Return(run)
	return _c
}

// SetupAccount provides a mock function with given fields: _a0, _a1
func (_m *console_server) SetupAccount(_a0 context.Context, _a1 *console.AccountSetupIn) (*console.AccountSetupVoid, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *console.AccountSetupVoid
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *console.AccountSetupIn) (*console.AccountSetupVoid, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *console.AccountSetupIn) *console.AccountSetupVoid); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*console.AccountSetupVoid)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *console.AccountSetupIn) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// console_server_SetupAccount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetupAccount'
type console_server_SetupAccount_Call struct {
	*mock.Call
}

// SetupAccount is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *console.AccountSetupIn
func (_e *console_server_Expecter) SetupAccount(_a0 interface{}, _a1 interface{}) *console_server_SetupAccount_Call {
	return &console_server_SetupAccount_Call{Call: _e.mock.On("SetupAccount", _a0, _a1)}
}

func (_c *console_server_SetupAccount_Call) Run(run func(_a0 context.Context, _a1 *console.AccountSetupIn)) *console_server_SetupAccount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*console.AccountSetupIn))
	})
	return _c
}

func (_c *console_server_SetupAccount_Call) Return(_a0 *console.AccountSetupVoid, _a1 error) *console_server_SetupAccount_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *console_server_SetupAccount_Call) RunAndReturn(run func(context.Context, *console.AccountSetupIn) (*console.AccountSetupVoid, error)) *console_server_SetupAccount_Call {
	_c.Call.Return(run)
	return _c
}

// mustEmbedUnimplementedConsoleServer provides a mock function with given fields:
func (_m *console_server) mustEmbedUnimplementedConsoleServer() {
	_m.Called()
}

// console_server_mustEmbedUnimplementedConsoleServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'mustEmbedUnimplementedConsoleServer'
type console_server_mustEmbedUnimplementedConsoleServer_Call struct {
	*mock.Call
}

// mustEmbedUnimplementedConsoleServer is a helper method to define mock.On call
func (_e *console_server_Expecter) mustEmbedUnimplementedConsoleServer() *console_server_mustEmbedUnimplementedConsoleServer_Call {
	return &console_server_mustEmbedUnimplementedConsoleServer_Call{Call: _e.mock.On("mustEmbedUnimplementedConsoleServer")}
}

func (_c *console_server_mustEmbedUnimplementedConsoleServer_Call) Run(run func()) *console_server_mustEmbedUnimplementedConsoleServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *console_server_mustEmbedUnimplementedConsoleServer_Call) Return() *console_server_mustEmbedUnimplementedConsoleServer_Call {
	_c.Call.Return()
	return _c
}

func (_c *console_server_mustEmbedUnimplementedConsoleServer_Call) RunAndReturn(run func()) *console_server_mustEmbedUnimplementedConsoleServer_Call {
	_c.Call.Return(run)
	return _c
}

// newConsole_server creates a new instance of console_server. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newConsole_server(t interface {
	mock.TestingT
	Cleanup(func())
}) *console_server {
	mock := &console_server{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
