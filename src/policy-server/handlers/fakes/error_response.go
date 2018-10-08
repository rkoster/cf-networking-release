// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"net/http"
	"sync"

	"code.cloudfoundry.org/lager"
)

type ErrorResponse struct {
	InternalServerErrorStub        func(lager.Logger, http.ResponseWriter, error, string)
	internalServerErrorMutex       sync.RWMutex
	internalServerErrorArgsForCall []struct {
		arg1 lager.Logger
		arg2 http.ResponseWriter
		arg3 error
		arg4 string
	}
	BadRequestStub        func(lager.Logger, http.ResponseWriter, error, string)
	badRequestMutex       sync.RWMutex
	badRequestArgsForCall []struct {
		arg1 lager.Logger
		arg2 http.ResponseWriter
		arg3 error
		arg4 string
	}
	NotFoundStub        func(lager.Logger, http.ResponseWriter, error, string)
	notFoundMutex       sync.RWMutex
	notFoundArgsForCall []struct {
		arg1 lager.Logger
		arg2 http.ResponseWriter
		arg3 error
		arg4 string
	}
	NotAcceptableStub        func(lager.Logger, http.ResponseWriter, error, string)
	notAcceptableMutex       sync.RWMutex
	notAcceptableArgsForCall []struct {
		arg1 lager.Logger
		arg2 http.ResponseWriter
		arg3 error
		arg4 string
	}
	ForbiddenStub        func(lager.Logger, http.ResponseWriter, error, string)
	forbiddenMutex       sync.RWMutex
	forbiddenArgsForCall []struct {
		arg1 lager.Logger
		arg2 http.ResponseWriter
		arg3 error
		arg4 string
	}
	UnauthorizedStub        func(lager.Logger, http.ResponseWriter, error, string)
	unauthorizedMutex       sync.RWMutex
	unauthorizedArgsForCall []struct {
		arg1 lager.Logger
		arg2 http.ResponseWriter
		arg3 error
		arg4 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ErrorResponse) InternalServerError(arg1 lager.Logger, arg2 http.ResponseWriter, arg3 error, arg4 string) {
	fake.internalServerErrorMutex.Lock()
	fake.internalServerErrorArgsForCall = append(fake.internalServerErrorArgsForCall, struct {
		arg1 lager.Logger
		arg2 http.ResponseWriter
		arg3 error
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("InternalServerError", []interface{}{arg1, arg2, arg3, arg4})
	fake.internalServerErrorMutex.Unlock()
	if fake.InternalServerErrorStub != nil {
		fake.InternalServerErrorStub(arg1, arg2, arg3, arg4)
	}
}

func (fake *ErrorResponse) InternalServerErrorCallCount() int {
	fake.internalServerErrorMutex.RLock()
	defer fake.internalServerErrorMutex.RUnlock()
	return len(fake.internalServerErrorArgsForCall)
}

func (fake *ErrorResponse) InternalServerErrorArgsForCall(i int) (lager.Logger, http.ResponseWriter, error, string) {
	fake.internalServerErrorMutex.RLock()
	defer fake.internalServerErrorMutex.RUnlock()
	return fake.internalServerErrorArgsForCall[i].arg1, fake.internalServerErrorArgsForCall[i].arg2, fake.internalServerErrorArgsForCall[i].arg3, fake.internalServerErrorArgsForCall[i].arg4
}

func (fake *ErrorResponse) BadRequest(arg1 lager.Logger, arg2 http.ResponseWriter, arg3 error, arg4 string) {
	fake.badRequestMutex.Lock()
	fake.badRequestArgsForCall = append(fake.badRequestArgsForCall, struct {
		arg1 lager.Logger
		arg2 http.ResponseWriter
		arg3 error
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("BadRequest", []interface{}{arg1, arg2, arg3, arg4})
	fake.badRequestMutex.Unlock()
	if fake.BadRequestStub != nil {
		fake.BadRequestStub(arg1, arg2, arg3, arg4)
	}
}

func (fake *ErrorResponse) BadRequestCallCount() int {
	fake.badRequestMutex.RLock()
	defer fake.badRequestMutex.RUnlock()
	return len(fake.badRequestArgsForCall)
}

func (fake *ErrorResponse) BadRequestArgsForCall(i int) (lager.Logger, http.ResponseWriter, error, string) {
	fake.badRequestMutex.RLock()
	defer fake.badRequestMutex.RUnlock()
	return fake.badRequestArgsForCall[i].arg1, fake.badRequestArgsForCall[i].arg2, fake.badRequestArgsForCall[i].arg3, fake.badRequestArgsForCall[i].arg4
}

func (fake *ErrorResponse) NotFound(arg1 lager.Logger, arg2 http.ResponseWriter, arg3 error, arg4 string) {
	fake.notFoundMutex.Lock()
	fake.notFoundArgsForCall = append(fake.notFoundArgsForCall, struct {
		arg1 lager.Logger
		arg2 http.ResponseWriter
		arg3 error
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("NotFound", []interface{}{arg1, arg2, arg3, arg4})
	fake.notFoundMutex.Unlock()
	if fake.NotFoundStub != nil {
		fake.NotFoundStub(arg1, arg2, arg3, arg4)
	}
}

func (fake *ErrorResponse) NotFoundCallCount() int {
	fake.notFoundMutex.RLock()
	defer fake.notFoundMutex.RUnlock()
	return len(fake.notFoundArgsForCall)
}

func (fake *ErrorResponse) NotFoundArgsForCall(i int) (lager.Logger, http.ResponseWriter, error, string) {
	fake.notFoundMutex.RLock()
	defer fake.notFoundMutex.RUnlock()
	return fake.notFoundArgsForCall[i].arg1, fake.notFoundArgsForCall[i].arg2, fake.notFoundArgsForCall[i].arg3, fake.notFoundArgsForCall[i].arg4
}

func (fake *ErrorResponse) NotAcceptable(arg1 lager.Logger, arg2 http.ResponseWriter, arg3 error, arg4 string) {
	fake.notAcceptableMutex.Lock()
	fake.notAcceptableArgsForCall = append(fake.notAcceptableArgsForCall, struct {
		arg1 lager.Logger
		arg2 http.ResponseWriter
		arg3 error
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("NotAcceptable", []interface{}{arg1, arg2, arg3, arg4})
	fake.notAcceptableMutex.Unlock()
	if fake.NotAcceptableStub != nil {
		fake.NotAcceptableStub(arg1, arg2, arg3, arg4)
	}
}

func (fake *ErrorResponse) NotAcceptableCallCount() int {
	fake.notAcceptableMutex.RLock()
	defer fake.notAcceptableMutex.RUnlock()
	return len(fake.notAcceptableArgsForCall)
}

func (fake *ErrorResponse) NotAcceptableArgsForCall(i int) (lager.Logger, http.ResponseWriter, error, string) {
	fake.notAcceptableMutex.RLock()
	defer fake.notAcceptableMutex.RUnlock()
	return fake.notAcceptableArgsForCall[i].arg1, fake.notAcceptableArgsForCall[i].arg2, fake.notAcceptableArgsForCall[i].arg3, fake.notAcceptableArgsForCall[i].arg4
}

func (fake *ErrorResponse) Forbidden(arg1 lager.Logger, arg2 http.ResponseWriter, arg3 error, arg4 string) {
	fake.forbiddenMutex.Lock()
	fake.forbiddenArgsForCall = append(fake.forbiddenArgsForCall, struct {
		arg1 lager.Logger
		arg2 http.ResponseWriter
		arg3 error
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("Forbidden", []interface{}{arg1, arg2, arg3, arg4})
	fake.forbiddenMutex.Unlock()
	if fake.ForbiddenStub != nil {
		fake.ForbiddenStub(arg1, arg2, arg3, arg4)
	}
}

func (fake *ErrorResponse) ForbiddenCallCount() int {
	fake.forbiddenMutex.RLock()
	defer fake.forbiddenMutex.RUnlock()
	return len(fake.forbiddenArgsForCall)
}

func (fake *ErrorResponse) ForbiddenArgsForCall(i int) (lager.Logger, http.ResponseWriter, error, string) {
	fake.forbiddenMutex.RLock()
	defer fake.forbiddenMutex.RUnlock()
	return fake.forbiddenArgsForCall[i].arg1, fake.forbiddenArgsForCall[i].arg2, fake.forbiddenArgsForCall[i].arg3, fake.forbiddenArgsForCall[i].arg4
}

func (fake *ErrorResponse) Unauthorized(arg1 lager.Logger, arg2 http.ResponseWriter, arg3 error, arg4 string) {
	fake.unauthorizedMutex.Lock()
	fake.unauthorizedArgsForCall = append(fake.unauthorizedArgsForCall, struct {
		arg1 lager.Logger
		arg2 http.ResponseWriter
		arg3 error
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("Unauthorized", []interface{}{arg1, arg2, arg3, arg4})
	fake.unauthorizedMutex.Unlock()
	if fake.UnauthorizedStub != nil {
		fake.UnauthorizedStub(arg1, arg2, arg3, arg4)
	}
}

func (fake *ErrorResponse) UnauthorizedCallCount() int {
	fake.unauthorizedMutex.RLock()
	defer fake.unauthorizedMutex.RUnlock()
	return len(fake.unauthorizedArgsForCall)
}

func (fake *ErrorResponse) UnauthorizedArgsForCall(i int) (lager.Logger, http.ResponseWriter, error, string) {
	fake.unauthorizedMutex.RLock()
	defer fake.unauthorizedMutex.RUnlock()
	return fake.unauthorizedArgsForCall[i].arg1, fake.unauthorizedArgsForCall[i].arg2, fake.unauthorizedArgsForCall[i].arg3, fake.unauthorizedArgsForCall[i].arg4
}

func (fake *ErrorResponse) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.internalServerErrorMutex.RLock()
	defer fake.internalServerErrorMutex.RUnlock()
	fake.badRequestMutex.RLock()
	defer fake.badRequestMutex.RUnlock()
	fake.notFoundMutex.RLock()
	defer fake.notFoundMutex.RUnlock()
	fake.notAcceptableMutex.RLock()
	defer fake.notAcceptableMutex.RUnlock()
	fake.forbiddenMutex.RLock()
	defer fake.forbiddenMutex.RUnlock()
	fake.unauthorizedMutex.RLock()
	defer fake.unauthorizedMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ErrorResponse) recordInvocation(key string, args []interface{}) {
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
