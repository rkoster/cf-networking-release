// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/vishvananda/netlink"
)

type NetlinkAdapter struct {
	RouteListStub        func(netlink.Link, int) ([]netlink.Route, error)
	routeListMutex       sync.RWMutex
	routeListArgsForCall []struct {
		arg1 netlink.Link
		arg2 int
	}
	routeListReturns struct {
		result1 []netlink.Route
		result2 error
	}
	routeListReturnsOnCall map[int]struct {
		result1 []netlink.Route
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *NetlinkAdapter) RouteList(arg1 netlink.Link, arg2 int) ([]netlink.Route, error) {
	fake.routeListMutex.Lock()
	ret, specificReturn := fake.routeListReturnsOnCall[len(fake.routeListArgsForCall)]
	fake.routeListArgsForCall = append(fake.routeListArgsForCall, struct {
		arg1 netlink.Link
		arg2 int
	}{arg1, arg2})
	fake.recordInvocation("RouteList", []interface{}{arg1, arg2})
	fake.routeListMutex.Unlock()
	if fake.RouteListStub != nil {
		return fake.RouteListStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.routeListReturns.result1, fake.routeListReturns.result2
}

func (fake *NetlinkAdapter) RouteListCallCount() int {
	fake.routeListMutex.RLock()
	defer fake.routeListMutex.RUnlock()
	return len(fake.routeListArgsForCall)
}

func (fake *NetlinkAdapter) RouteListArgsForCall(i int) (netlink.Link, int) {
	fake.routeListMutex.RLock()
	defer fake.routeListMutex.RUnlock()
	return fake.routeListArgsForCall[i].arg1, fake.routeListArgsForCall[i].arg2
}

func (fake *NetlinkAdapter) RouteListReturns(result1 []netlink.Route, result2 error) {
	fake.RouteListStub = nil
	fake.routeListReturns = struct {
		result1 []netlink.Route
		result2 error
	}{result1, result2}
}

func (fake *NetlinkAdapter) RouteListReturnsOnCall(i int, result1 []netlink.Route, result2 error) {
	fake.RouteListStub = nil
	if fake.routeListReturnsOnCall == nil {
		fake.routeListReturnsOnCall = make(map[int]struct {
			result1 []netlink.Route
			result2 error
		})
	}
	fake.routeListReturnsOnCall[i] = struct {
		result1 []netlink.Route
		result2 error
	}{result1, result2}
}

func (fake *NetlinkAdapter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.routeListMutex.RLock()
	defer fake.routeListMutex.RUnlock()
	return fake.invocations
}

func (fake *NetlinkAdapter) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}