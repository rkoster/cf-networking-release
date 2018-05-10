// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"
)

type MetricsSender struct {
	IncrementCounterStub        func(string)
	incrementCounterMutex       sync.RWMutex
	incrementCounterArgsForCall []struct {
		arg1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *MetricsSender) IncrementCounter(arg1 string) {
	fake.incrementCounterMutex.Lock()
	fake.incrementCounterArgsForCall = append(fake.incrementCounterArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("IncrementCounter", []interface{}{arg1})
	fake.incrementCounterMutex.Unlock()
	if fake.IncrementCounterStub != nil {
		fake.IncrementCounterStub(arg1)
	}
}

func (fake *MetricsSender) IncrementCounterCallCount() int {
	fake.incrementCounterMutex.RLock()
	defer fake.incrementCounterMutex.RUnlock()
	return len(fake.incrementCounterArgsForCall)
}

func (fake *MetricsSender) IncrementCounterArgsForCall(i int) string {
	fake.incrementCounterMutex.RLock()
	defer fake.incrementCounterMutex.RUnlock()
	return fake.incrementCounterArgsForCall[i].arg1
}

func (fake *MetricsSender) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.incrementCounterMutex.RLock()
	defer fake.incrementCounterMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *MetricsSender) recordInvocation(key string, args []interface{}) {
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