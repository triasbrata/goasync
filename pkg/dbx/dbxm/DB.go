// Code generated by mockery v2.43.2. DO NOT EDIT.

package dbxm

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	dbx "github.com/triasbrata/golibs/pkg/dbx"

	sql "database/sql"

	sqlx "github.com/jmoiron/sqlx"
)

// DB is an autogenerated mock type for the DB type
type DB struct {
	mock.Mock
}

// BeginTxx provides a mock function with given fields: ctx, opts
func (_m *DB) BeginTxx(ctx context.Context, opts *sql.TxOptions) (dbx.Tx, error) {
	ret := _m.Called(ctx, opts)

	if len(ret) == 0 {
		panic("no return value specified for BeginTxx")
	}

	var r0 dbx.Tx
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *sql.TxOptions) (dbx.Tx, error)); ok {
		return rf(ctx, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *sql.TxOptions) dbx.Tx); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dbx.Tx)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *sql.TxOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BindNamed provides a mock function with given fields: query, arg
func (_m *DB) BindNamed(query string, arg interface{}) (string, []interface{}, error) {
	ret := _m.Called(query, arg)

	if len(ret) == 0 {
		panic("no return value specified for BindNamed")
	}

	var r0 string
	var r1 []interface{}
	var r2 error
	if rf, ok := ret.Get(0).(func(string, interface{}) (string, []interface{}, error)); ok {
		return rf(query, arg)
	}
	if rf, ok := ret.Get(0).(func(string, interface{}) string); ok {
		r0 = rf(query, arg)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, interface{}) []interface{}); ok {
		r1 = rf(query, arg)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]interface{})
		}
	}

	if rf, ok := ret.Get(2).(func(string, interface{}) error); ok {
		r2 = rf(query, arg)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Close provides a mock function with given fields:
func (_m *DB) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExecContext provides a mock function with given fields: ctx, query, args
func (_m *DB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ExecContext")
	}

	var r0 sql.Result
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) (sql.Result, error)); ok {
		return rf(ctx, query, args...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) sql.Result); ok {
		r0 = rf(ctx, query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sql.Result)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...interface{}) error); ok {
		r1 = rf(ctx, query, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetContext provides a mock function with given fields: ctx, dest, query, args
func (_m *DB) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, ctx, dest, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetContext")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, string, ...interface{}) error); ok {
		r0 = rf(ctx, dest, query, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NamedExecContext provides a mock function with given fields: ctx, query, arg
func (_m *DB) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	ret := _m.Called(ctx, query, arg)

	if len(ret) == 0 {
		panic("no return value specified for NamedExecContext")
	}

	var r0 sql.Result
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}) (sql.Result, error)); ok {
		return rf(ctx, query, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}) sql.Result); ok {
		r0 = rf(ctx, query, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sql.Result)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, interface{}) error); ok {
		r1 = rf(ctx, query, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NamedQuery provides a mock function with given fields: query, arg
func (_m *DB) NamedQuery(query string, arg interface{}) (*sqlx.Rows, error) {
	ret := _m.Called(query, arg)

	if len(ret) == 0 {
		panic("no return value specified for NamedQuery")
	}

	var r0 *sqlx.Rows
	var r1 error
	if rf, ok := ret.Get(0).(func(string, interface{}) (*sqlx.Rows, error)); ok {
		return rf(query, arg)
	}
	if rf, ok := ret.Get(0).(func(string, interface{}) *sqlx.Rows); ok {
		r0 = rf(query, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqlx.Rows)
		}
	}

	if rf, ok := ret.Get(1).(func(string, interface{}) error); ok {
		r1 = rf(query, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NamedQueryContext provides a mock function with given fields: ctx, query, arg
func (_m *DB) NamedQueryContext(ctx context.Context, query string, arg interface{}) (*sqlx.Rows, error) {
	ret := _m.Called(ctx, query, arg)

	if len(ret) == 0 {
		panic("no return value specified for NamedQueryContext")
	}

	var r0 *sqlx.Rows
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}) (*sqlx.Rows, error)); ok {
		return rf(ctx, query, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}) *sqlx.Rows); ok {
		r0 = rf(ctx, query, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqlx.Rows)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, interface{}) error); ok {
		r1 = rf(ctx, query, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectContext provides a mock function with given fields: ctx, dest, query, args
func (_m *DB) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, ctx, dest, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SelectContext")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, string, ...interface{}) error); ok {
		r0 = rf(ctx, dest, query, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewDB creates a new instance of DB. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDB(t interface {
	mock.TestingT
	Cleanup(func())
}) *DB {
	mock := &DB{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
