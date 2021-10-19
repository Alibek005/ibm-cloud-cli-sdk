// Code generated by counterfeiter. DO NOT EDIT.
package pluginfakes

import (
	"sync"

	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/plugin"
)

type FakePlugin struct {
	GetMetadataStub        func() plugin.PluginMetadata
	getMetadataMutex       sync.RWMutex
	getMetadataArgsForCall []struct {
	}
	getMetadataReturns struct {
		result1 plugin.PluginMetadata
	}
	getMetadataReturnsOnCall map[int]struct {
		result1 plugin.PluginMetadata
	}
	RunStub        func(plugin.PluginContext, []string)
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		arg1 plugin.PluginContext
		arg2 []string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePlugin) GetMetadata() plugin.PluginMetadata {
	fake.getMetadataMutex.Lock()
	ret, specificReturn := fake.getMetadataReturnsOnCall[len(fake.getMetadataArgsForCall)]
	fake.getMetadataArgsForCall = append(fake.getMetadataArgsForCall, struct {
	}{})
	fake.recordInvocation("GetMetadata", []interface{}{})
	fake.getMetadataMutex.Unlock()
	if fake.GetMetadataStub != nil {
		return fake.GetMetadataStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getMetadataReturns
	return fakeReturns.result1
}

func (fake *FakePlugin) GetMetadataCallCount() int {
	fake.getMetadataMutex.RLock()
	defer fake.getMetadataMutex.RUnlock()
	return len(fake.getMetadataArgsForCall)
}

func (fake *FakePlugin) GetMetadataCalls(stub func() plugin.PluginMetadata) {
	fake.getMetadataMutex.Lock()
	defer fake.getMetadataMutex.Unlock()
	fake.GetMetadataStub = stub
}

func (fake *FakePlugin) GetMetadataReturns(result1 plugin.PluginMetadata) {
	fake.getMetadataMutex.Lock()
	defer fake.getMetadataMutex.Unlock()
	fake.GetMetadataStub = nil
	fake.getMetadataReturns = struct {
		result1 plugin.PluginMetadata
	}{result1}
}

func (fake *FakePlugin) GetMetadataReturnsOnCall(i int, result1 plugin.PluginMetadata) {
	fake.getMetadataMutex.Lock()
	defer fake.getMetadataMutex.Unlock()
	fake.GetMetadataStub = nil
	if fake.getMetadataReturnsOnCall == nil {
		fake.getMetadataReturnsOnCall = make(map[int]struct {
			result1 plugin.PluginMetadata
		})
	}
	fake.getMetadataReturnsOnCall[i] = struct {
		result1 plugin.PluginMetadata
	}{result1}
}

func (fake *FakePlugin) Run(arg1 plugin.PluginContext, arg2 []string) {
	var arg2Copy []string
	if arg2 != nil {
		arg2Copy = make([]string, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.runMutex.Lock()
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
		arg1 plugin.PluginContext
		arg2 []string
	}{arg1, arg2Copy})
	fake.recordInvocation("Run", []interface{}{arg1, arg2Copy})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		fake.RunStub(arg1, arg2)
	}
}

func (fake *FakePlugin) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakePlugin) RunCalls(stub func(plugin.PluginContext, []string)) {
	fake.runMutex.Lock()
	defer fake.runMutex.Unlock()
	fake.RunStub = stub
}

func (fake *FakePlugin) RunArgsForCall(i int) (plugin.PluginContext, []string) {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	argsForCall := fake.runArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePlugin) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getMetadataMutex.RLock()
	defer fake.getMetadataMutex.RUnlock()
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePlugin) recordInvocation(key string, args []interface{}) {
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

var _ plugin.Plugin = new(FakePlugin)
