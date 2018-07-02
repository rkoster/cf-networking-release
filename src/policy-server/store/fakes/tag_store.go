// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"policy-server/store"
	"sync"
)

type TagStore struct {
	CreateTagStub        func(string, string) (store.Tag, error)
	createTagMutex       sync.RWMutex
	createTagArgsForCall []struct {
		arg1 string
		arg2 string
	}
	createTagReturns struct {
		result1 store.Tag
		result2 error
	}
	createTagReturnsOnCall map[int]struct {
		result1 store.Tag
		result2 error
	}
	TagsStub        func() ([]store.Tag, error)
	tagsMutex       sync.RWMutex
	tagsArgsForCall []struct{}
	tagsReturns     struct {
		result1 []store.Tag
		result2 error
	}
	tagsReturnsOnCall map[int]struct {
		result1 []store.Tag
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *TagStore) CreateTag(arg1 string, arg2 string) (store.Tag, error) {
	fake.createTagMutex.Lock()
	ret, specificReturn := fake.createTagReturnsOnCall[len(fake.createTagArgsForCall)]
	fake.createTagArgsForCall = append(fake.createTagArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("CreateTag", []interface{}{arg1, arg2})
	fake.createTagMutex.Unlock()
	if fake.CreateTagStub != nil {
		return fake.CreateTagStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createTagReturns.result1, fake.createTagReturns.result2
}

func (fake *TagStore) CreateTagCallCount() int {
	fake.createTagMutex.RLock()
	defer fake.createTagMutex.RUnlock()
	return len(fake.createTagArgsForCall)
}

func (fake *TagStore) CreateTagArgsForCall(i int) (string, string) {
	fake.createTagMutex.RLock()
	defer fake.createTagMutex.RUnlock()
	return fake.createTagArgsForCall[i].arg1, fake.createTagArgsForCall[i].arg2
}

func (fake *TagStore) CreateTagReturns(result1 store.Tag, result2 error) {
	fake.CreateTagStub = nil
	fake.createTagReturns = struct {
		result1 store.Tag
		result2 error
	}{result1, result2}
}

func (fake *TagStore) CreateTagReturnsOnCall(i int, result1 store.Tag, result2 error) {
	fake.CreateTagStub = nil
	if fake.createTagReturnsOnCall == nil {
		fake.createTagReturnsOnCall = make(map[int]struct {
			result1 store.Tag
			result2 error
		})
	}
	fake.createTagReturnsOnCall[i] = struct {
		result1 store.Tag
		result2 error
	}{result1, result2}
}

func (fake *TagStore) Tags() ([]store.Tag, error) {
	fake.tagsMutex.Lock()
	ret, specificReturn := fake.tagsReturnsOnCall[len(fake.tagsArgsForCall)]
	fake.tagsArgsForCall = append(fake.tagsArgsForCall, struct{}{})
	fake.recordInvocation("Tags", []interface{}{})
	fake.tagsMutex.Unlock()
	if fake.TagsStub != nil {
		return fake.TagsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.tagsReturns.result1, fake.tagsReturns.result2
}

func (fake *TagStore) TagsCallCount() int {
	fake.tagsMutex.RLock()
	defer fake.tagsMutex.RUnlock()
	return len(fake.tagsArgsForCall)
}

func (fake *TagStore) TagsReturns(result1 []store.Tag, result2 error) {
	fake.TagsStub = nil
	fake.tagsReturns = struct {
		result1 []store.Tag
		result2 error
	}{result1, result2}
}

func (fake *TagStore) TagsReturnsOnCall(i int, result1 []store.Tag, result2 error) {
	fake.TagsStub = nil
	if fake.tagsReturnsOnCall == nil {
		fake.tagsReturnsOnCall = make(map[int]struct {
			result1 []store.Tag
			result2 error
		})
	}
	fake.tagsReturnsOnCall[i] = struct {
		result1 []store.Tag
		result2 error
	}{result1, result2}
}

func (fake *TagStore) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createTagMutex.RLock()
	defer fake.createTagMutex.RUnlock()
	fake.tagsMutex.RLock()
	defer fake.tagsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *TagStore) recordInvocation(key string, args []interface{}) {
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

var _ store.TagStore = new(TagStore)