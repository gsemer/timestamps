// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"
	"time"
	"timestamps/domain"
)

type FakeTimestampsRepository struct {
	DayStub        func(string, time.Time, time.Time, *time.Location, string) ([]string, error)
	dayMutex       sync.RWMutex
	dayArgsForCall []struct {
		arg1 string
		arg2 time.Time
		arg3 time.Time
		arg4 *time.Location
		arg5 string
	}
	dayReturns struct {
		result1 []string
		result2 error
	}
	dayReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	HourStub        func(string, time.Time, time.Time, *time.Location, string) ([]string, error)
	hourMutex       sync.RWMutex
	hourArgsForCall []struct {
		arg1 string
		arg2 time.Time
		arg3 time.Time
		arg4 *time.Location
		arg5 string
	}
	hourReturns struct {
		result1 []string
		result2 error
	}
	hourReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	MonthStub        func(string, time.Time, time.Time, *time.Location, string) ([]string, error)
	monthMutex       sync.RWMutex
	monthArgsForCall []struct {
		arg1 string
		arg2 time.Time
		arg3 time.Time
		arg4 *time.Location
		arg5 string
	}
	monthReturns struct {
		result1 []string
		result2 error
	}
	monthReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	YearStub        func(string, time.Time, time.Time, *time.Location, string) ([]string, error)
	yearMutex       sync.RWMutex
	yearArgsForCall []struct {
		arg1 string
		arg2 time.Time
		arg3 time.Time
		arg4 *time.Location
		arg5 string
	}
	yearReturns struct {
		result1 []string
		result2 error
	}
	yearReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTimestampsRepository) Day(arg1 string, arg2 time.Time, arg3 time.Time, arg4 *time.Location, arg5 string) ([]string, error) {
	fake.dayMutex.Lock()
	ret, specificReturn := fake.dayReturnsOnCall[len(fake.dayArgsForCall)]
	fake.dayArgsForCall = append(fake.dayArgsForCall, struct {
		arg1 string
		arg2 time.Time
		arg3 time.Time
		arg4 *time.Location
		arg5 string
	}{arg1, arg2, arg3, arg4, arg5})
	stub := fake.DayStub
	fakeReturns := fake.dayReturns
	fake.recordInvocation("Day", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.dayMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTimestampsRepository) DayCallCount() int {
	fake.dayMutex.RLock()
	defer fake.dayMutex.RUnlock()
	return len(fake.dayArgsForCall)
}

func (fake *FakeTimestampsRepository) DayCalls(stub func(string, time.Time, time.Time, *time.Location, string) ([]string, error)) {
	fake.dayMutex.Lock()
	defer fake.dayMutex.Unlock()
	fake.DayStub = stub
}

func (fake *FakeTimestampsRepository) DayArgsForCall(i int) (string, time.Time, time.Time, *time.Location, string) {
	fake.dayMutex.RLock()
	defer fake.dayMutex.RUnlock()
	argsForCall := fake.dayArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeTimestampsRepository) DayReturns(result1 []string, result2 error) {
	fake.dayMutex.Lock()
	defer fake.dayMutex.Unlock()
	fake.DayStub = nil
	fake.dayReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeTimestampsRepository) DayReturnsOnCall(i int, result1 []string, result2 error) {
	fake.dayMutex.Lock()
	defer fake.dayMutex.Unlock()
	fake.DayStub = nil
	if fake.dayReturnsOnCall == nil {
		fake.dayReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.dayReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeTimestampsRepository) Hour(arg1 string, arg2 time.Time, arg3 time.Time, arg4 *time.Location, arg5 string) ([]string, error) {
	fake.hourMutex.Lock()
	ret, specificReturn := fake.hourReturnsOnCall[len(fake.hourArgsForCall)]
	fake.hourArgsForCall = append(fake.hourArgsForCall, struct {
		arg1 string
		arg2 time.Time
		arg3 time.Time
		arg4 *time.Location
		arg5 string
	}{arg1, arg2, arg3, arg4, arg5})
	stub := fake.HourStub
	fakeReturns := fake.hourReturns
	fake.recordInvocation("Hour", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.hourMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTimestampsRepository) HourCallCount() int {
	fake.hourMutex.RLock()
	defer fake.hourMutex.RUnlock()
	return len(fake.hourArgsForCall)
}

func (fake *FakeTimestampsRepository) HourCalls(stub func(string, time.Time, time.Time, *time.Location, string) ([]string, error)) {
	fake.hourMutex.Lock()
	defer fake.hourMutex.Unlock()
	fake.HourStub = stub
}

func (fake *FakeTimestampsRepository) HourArgsForCall(i int) (string, time.Time, time.Time, *time.Location, string) {
	fake.hourMutex.RLock()
	defer fake.hourMutex.RUnlock()
	argsForCall := fake.hourArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeTimestampsRepository) HourReturns(result1 []string, result2 error) {
	fake.hourMutex.Lock()
	defer fake.hourMutex.Unlock()
	fake.HourStub = nil
	fake.hourReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeTimestampsRepository) HourReturnsOnCall(i int, result1 []string, result2 error) {
	fake.hourMutex.Lock()
	defer fake.hourMutex.Unlock()
	fake.HourStub = nil
	if fake.hourReturnsOnCall == nil {
		fake.hourReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.hourReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeTimestampsRepository) Month(arg1 string, arg2 time.Time, arg3 time.Time, arg4 *time.Location, arg5 string) ([]string, error) {
	fake.monthMutex.Lock()
	ret, specificReturn := fake.monthReturnsOnCall[len(fake.monthArgsForCall)]
	fake.monthArgsForCall = append(fake.monthArgsForCall, struct {
		arg1 string
		arg2 time.Time
		arg3 time.Time
		arg4 *time.Location
		arg5 string
	}{arg1, arg2, arg3, arg4, arg5})
	stub := fake.MonthStub
	fakeReturns := fake.monthReturns
	fake.recordInvocation("Month", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.monthMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTimestampsRepository) MonthCallCount() int {
	fake.monthMutex.RLock()
	defer fake.monthMutex.RUnlock()
	return len(fake.monthArgsForCall)
}

func (fake *FakeTimestampsRepository) MonthCalls(stub func(string, time.Time, time.Time, *time.Location, string) ([]string, error)) {
	fake.monthMutex.Lock()
	defer fake.monthMutex.Unlock()
	fake.MonthStub = stub
}

func (fake *FakeTimestampsRepository) MonthArgsForCall(i int) (string, time.Time, time.Time, *time.Location, string) {
	fake.monthMutex.RLock()
	defer fake.monthMutex.RUnlock()
	argsForCall := fake.monthArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeTimestampsRepository) MonthReturns(result1 []string, result2 error) {
	fake.monthMutex.Lock()
	defer fake.monthMutex.Unlock()
	fake.MonthStub = nil
	fake.monthReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeTimestampsRepository) MonthReturnsOnCall(i int, result1 []string, result2 error) {
	fake.monthMutex.Lock()
	defer fake.monthMutex.Unlock()
	fake.MonthStub = nil
	if fake.monthReturnsOnCall == nil {
		fake.monthReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.monthReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeTimestampsRepository) Year(arg1 string, arg2 time.Time, arg3 time.Time, arg4 *time.Location, arg5 string) ([]string, error) {
	fake.yearMutex.Lock()
	ret, specificReturn := fake.yearReturnsOnCall[len(fake.yearArgsForCall)]
	fake.yearArgsForCall = append(fake.yearArgsForCall, struct {
		arg1 string
		arg2 time.Time
		arg3 time.Time
		arg4 *time.Location
		arg5 string
	}{arg1, arg2, arg3, arg4, arg5})
	stub := fake.YearStub
	fakeReturns := fake.yearReturns
	fake.recordInvocation("Year", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.yearMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTimestampsRepository) YearCallCount() int {
	fake.yearMutex.RLock()
	defer fake.yearMutex.RUnlock()
	return len(fake.yearArgsForCall)
}

func (fake *FakeTimestampsRepository) YearCalls(stub func(string, time.Time, time.Time, *time.Location, string) ([]string, error)) {
	fake.yearMutex.Lock()
	defer fake.yearMutex.Unlock()
	fake.YearStub = stub
}

func (fake *FakeTimestampsRepository) YearArgsForCall(i int) (string, time.Time, time.Time, *time.Location, string) {
	fake.yearMutex.RLock()
	defer fake.yearMutex.RUnlock()
	argsForCall := fake.yearArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeTimestampsRepository) YearReturns(result1 []string, result2 error) {
	fake.yearMutex.Lock()
	defer fake.yearMutex.Unlock()
	fake.YearStub = nil
	fake.yearReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeTimestampsRepository) YearReturnsOnCall(i int, result1 []string, result2 error) {
	fake.yearMutex.Lock()
	defer fake.yearMutex.Unlock()
	fake.YearStub = nil
	if fake.yearReturnsOnCall == nil {
		fake.yearReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.yearReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeTimestampsRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.dayMutex.RLock()
	defer fake.dayMutex.RUnlock()
	fake.hourMutex.RLock()
	defer fake.hourMutex.RUnlock()
	fake.monthMutex.RLock()
	defer fake.monthMutex.RUnlock()
	fake.yearMutex.RLock()
	defer fake.yearMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTimestampsRepository) recordInvocation(key string, args []interface{}) {
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

var _ domain.TimestampsRepository = new(FakeTimestampsRepository)
