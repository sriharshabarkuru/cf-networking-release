// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"
	"time"
)

type MetricsSender struct {
	IncrementCounterStub        func(string)
	incrementCounterMutex       sync.RWMutex
	incrementCounterArgsForCall []struct {
		arg1 string
	}
	SendDurationStub        func(string, time.Duration)
	sendDurationMutex       sync.RWMutex
	sendDurationArgsForCall []struct {
		arg1 string
		arg2 time.Duration
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

func (fake *MetricsSender) SendDuration(arg1 string, arg2 time.Duration) {
	fake.sendDurationMutex.Lock()
	fake.sendDurationArgsForCall = append(fake.sendDurationArgsForCall, struct {
		arg1 string
		arg2 time.Duration
	}{arg1, arg2})
	fake.recordInvocation("SendDuration", []interface{}{arg1, arg2})
	fake.sendDurationMutex.Unlock()
	if fake.SendDurationStub != nil {
		fake.SendDurationStub(arg1, arg2)
	}
}

func (fake *MetricsSender) SendDurationCallCount() int {
	fake.sendDurationMutex.RLock()
	defer fake.sendDurationMutex.RUnlock()
	return len(fake.sendDurationArgsForCall)
}

func (fake *MetricsSender) SendDurationArgsForCall(i int) (string, time.Duration) {
	fake.sendDurationMutex.RLock()
	defer fake.sendDurationMutex.RUnlock()
	return fake.sendDurationArgsForCall[i].arg1, fake.sendDurationArgsForCall[i].arg2
}

func (fake *MetricsSender) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.incrementCounterMutex.RLock()
	defer fake.incrementCounterMutex.RUnlock()
	fake.sendDurationMutex.RLock()
	defer fake.sendDurationMutex.RUnlock()
	return fake.invocations
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
