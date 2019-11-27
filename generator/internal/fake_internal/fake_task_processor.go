// Code generated by counterfeiter. DO NOT EDIT.
package fake_internal

import (
	"sync"

	"code.cloudfoundry.org/executor"
	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/rep/generator/internal"
)

type FakeTaskProcessor struct {
	ProcessStub        func(lager.Logger, executor.Container)
	processMutex       sync.RWMutex
	processArgsForCall []struct {
		arg1 lager.Logger
		arg2 executor.Container
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTaskProcessor) Process(arg1 lager.Logger, arg2 executor.Container) {
	fake.processMutex.Lock()
	fake.processArgsForCall = append(fake.processArgsForCall, struct {
		arg1 lager.Logger
		arg2 executor.Container
	}{arg1, arg2})
	fake.recordInvocation("Process", []interface{}{arg1, arg2})
	processStubCopy := fake.ProcessStub
	fake.processMutex.Unlock()
	if processStubCopy != nil {
		processStubCopy(arg1, arg2)
	}
}

func (fake *FakeTaskProcessor) ProcessCallCount() int {
	fake.processMutex.RLock()
	defer fake.processMutex.RUnlock()
	return len(fake.processArgsForCall)
}

func (fake *FakeTaskProcessor) ProcessCalls(stub func(lager.Logger, executor.Container)) {
	fake.processMutex.Lock()
	defer fake.processMutex.Unlock()
	fake.ProcessStub = stub
}

func (fake *FakeTaskProcessor) ProcessArgsForCall(i int) (lager.Logger, executor.Container) {
	fake.processMutex.RLock()
	defer fake.processMutex.RUnlock()
	argsForCall := fake.processArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeTaskProcessor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.processMutex.RLock()
	defer fake.processMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTaskProcessor) recordInvocation(key string, args []interface{}) {
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

var _ internal.TaskProcessor = new(FakeTaskProcessor)
