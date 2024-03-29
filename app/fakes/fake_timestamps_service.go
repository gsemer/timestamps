// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"
	"timestamps/domain"
)

type FakeTimestampsService struct {
	GetMatchingTimestampsStub        func(string, string, string, string, string) ([]string, error)
	getMatchingTimestampsMutex       sync.RWMutex
	getMatchingTimestampsArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 string
	}
	getMatchingTimestampsReturns struct {
		result1 []string
		result2 error
	}
	getMatchingTimestampsReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTimestampsService) GetMatchingTimestamps(arg1 string, arg2 string, arg3 string, arg4 string, arg5 string) ([]string, error) {
	fake.getMatchingTimestampsMutex.Lock()
	ret, specificReturn := fake.getMatchingTimestampsReturnsOnCall[len(fake.getMatchingTimestampsArgsForCall)]
	fake.getMatchingTimestampsArgsForCall = append(fake.getMatchingTimestampsArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 string
	}{arg1, arg2, arg3, arg4, arg5})
	stub := fake.GetMatchingTimestampsStub
	fakeReturns := fake.getMatchingTimestampsReturns
	fake.recordInvocation("GetMatchingTimestamps", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.getMatchingTimestampsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTimestampsService) GetMatchingTimestampsCallCount() int {
	fake.getMatchingTimestampsMutex.RLock()
	defer fake.getMatchingTimestampsMutex.RUnlock()
	return len(fake.getMatchingTimestampsArgsForCall)
}

func (fake *FakeTimestampsService) GetMatchingTimestampsCalls(stub func(string, string, string, string, string) ([]string, error)) {
	fake.getMatchingTimestampsMutex.Lock()
	defer fake.getMatchingTimestampsMutex.Unlock()
	fake.GetMatchingTimestampsStub = stub
}

func (fake *FakeTimestampsService) GetMatchingTimestampsArgsForCall(i int) (string, string, string, string, string) {
	fake.getMatchingTimestampsMutex.RLock()
	defer fake.getMatchingTimestampsMutex.RUnlock()
	argsForCall := fake.getMatchingTimestampsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeTimestampsService) GetMatchingTimestampsReturns(result1 []string, result2 error) {
	fake.getMatchingTimestampsMutex.Lock()
	defer fake.getMatchingTimestampsMutex.Unlock()
	fake.GetMatchingTimestampsStub = nil
	fake.getMatchingTimestampsReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeTimestampsService) GetMatchingTimestampsReturnsOnCall(i int, result1 []string, result2 error) {
	fake.getMatchingTimestampsMutex.Lock()
	defer fake.getMatchingTimestampsMutex.Unlock()
	fake.GetMatchingTimestampsStub = nil
	if fake.getMatchingTimestampsReturnsOnCall == nil {
		fake.getMatchingTimestampsReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.getMatchingTimestampsReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeTimestampsService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getMatchingTimestampsMutex.RLock()
	defer fake.getMatchingTimestampsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTimestampsService) recordInvocation(key string, args []interface{}) {
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

var _ domain.TimestampsService = new(FakeTimestampsService)
