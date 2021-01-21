// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	caddyapi "github.com/conradludgate/terraform-provider-caddy/caddyapi"
	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// CreateHTTP provides a mock function with given fields: http
func (_m *Client) CreateHTTP(http caddyapi.HTTP) error {
	ret := _m.Called(http)

	var r0 error
	if rf, ok := ret.Get(0).(func(caddyapi.HTTP) error); ok {
		r0 = rf(http)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateServer provides a mock function with given fields: name, server
func (_m *Client) CreateServer(name string, server caddyapi.Server) (string, error) {
	ret := _m.Called(name, server)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, caddyapi.Server) string); ok {
		r0 = rf(name, server)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, caddyapi.Server) error); ok {
		r1 = rf(name, server)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteHTTP provides a mock function with given fields:
func (_m *Client) DeleteHTTP() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteServer provides a mock function with given fields: id
func (_m *Client) DeleteServer(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetHTTP provides a mock function with given fields:
func (_m *Client) GetHTTP() (*caddyapi.HTTP, error) {
	ret := _m.Called()

	var r0 *caddyapi.HTTP
	if rf, ok := ret.Get(0).(func() *caddyapi.HTTP); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(*caddyapi.HTTP)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetServer provides a mock function with given fields: id
func (_m *Client) GetServer(id string) (*caddyapi.Server, error) {
	ret := _m.Called(id)

	var r0 *caddyapi.Server
	if rf, ok := ret.Get(0).(func(string) *caddyapi.Server); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(*caddyapi.Server)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateHTTPGracePeriod provides a mock function with given fields: gracePeriod
func (_m *Client) UpdateHTTPGracePeriod(gracePeriod caddyapi.Duration) error {
	ret := _m.Called(gracePeriod)

	var r0 error
	if rf, ok := ret.Get(0).(func(caddyapi.Duration) error); ok {
		r0 = rf(gracePeriod)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateHTTPPort provides a mock function with given fields: httpPort
func (_m *Client) UpdateHTTPPort(httpPort int) error {
	ret := _m.Called(httpPort)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(httpPort)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateHTTPSPort provides a mock function with given fields: httpsPort
func (_m *Client) UpdateHTTPSPort(httpsPort int) error {
	ret := _m.Called(httpsPort)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(httpsPort)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateServerListen provides a mock function with given fields: id, listen
func (_m *Client) UpdateServerListen(id string, listen []string) error {
	ret := _m.Called(id, listen)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []string) error); ok {
		r0 = rf(id, listen)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateServerRoutes provides a mock function with given fields: id, routes
func (_m *Client) UpdateServerRoutes(id string, routes []caddyapi.Route) error {
	ret := _m.Called(id, routes)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []caddyapi.Route) error); ok {
		r0 = rf(id, routes)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
