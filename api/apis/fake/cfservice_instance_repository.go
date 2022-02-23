// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"context"
	"sync"

	"code.cloudfoundry.org/cf-k8s-controllers/api/apis"
	"code.cloudfoundry.org/cf-k8s-controllers/api/authorization"
	"code.cloudfoundry.org/cf-k8s-controllers/api/repositories"
)

type CFServiceInstanceRepository struct {
	CreateServiceInstanceStub        func(context.Context, authorization.Info, repositories.CreateServiceInstanceMessage) (repositories.ServiceInstanceRecord, error)
	createServiceInstanceMutex       sync.RWMutex
	createServiceInstanceArgsForCall []struct {
		arg1 context.Context
		arg2 authorization.Info
		arg3 repositories.CreateServiceInstanceMessage
	}
	createServiceInstanceReturns struct {
		result1 repositories.ServiceInstanceRecord
		result2 error
	}
	createServiceInstanceReturnsOnCall map[int]struct {
		result1 repositories.ServiceInstanceRecord
		result2 error
	}
	DeleteServiceInstanceStub        func(context.Context, authorization.Info, repositories.DeleteServiceInstanceMessage) error
	deleteServiceInstanceMutex       sync.RWMutex
	deleteServiceInstanceArgsForCall []struct {
		arg1 context.Context
		arg2 authorization.Info
		arg3 repositories.DeleteServiceInstanceMessage
	}
	deleteServiceInstanceReturns struct {
		result1 error
	}
	deleteServiceInstanceReturnsOnCall map[int]struct {
		result1 error
	}
	GetServiceInstanceStub        func(context.Context, authorization.Info, string) (repositories.ServiceInstanceRecord, error)
	getServiceInstanceMutex       sync.RWMutex
	getServiceInstanceArgsForCall []struct {
		arg1 context.Context
		arg2 authorization.Info
		arg3 string
	}
	getServiceInstanceReturns struct {
		result1 repositories.ServiceInstanceRecord
		result2 error
	}
	getServiceInstanceReturnsOnCall map[int]struct {
		result1 repositories.ServiceInstanceRecord
		result2 error
	}
	ListServiceInstancesStub        func(context.Context, authorization.Info, repositories.ListServiceInstanceMessage) ([]repositories.ServiceInstanceRecord, error)
	listServiceInstancesMutex       sync.RWMutex
	listServiceInstancesArgsForCall []struct {
		arg1 context.Context
		arg2 authorization.Info
		arg3 repositories.ListServiceInstanceMessage
	}
	listServiceInstancesReturns struct {
		result1 []repositories.ServiceInstanceRecord
		result2 error
	}
	listServiceInstancesReturnsOnCall map[int]struct {
		result1 []repositories.ServiceInstanceRecord
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *CFServiceInstanceRepository) CreateServiceInstance(arg1 context.Context, arg2 authorization.Info, arg3 repositories.CreateServiceInstanceMessage) (repositories.ServiceInstanceRecord, error) {
	fake.createServiceInstanceMutex.Lock()
	ret, specificReturn := fake.createServiceInstanceReturnsOnCall[len(fake.createServiceInstanceArgsForCall)]
	fake.createServiceInstanceArgsForCall = append(fake.createServiceInstanceArgsForCall, struct {
		arg1 context.Context
		arg2 authorization.Info
		arg3 repositories.CreateServiceInstanceMessage
	}{arg1, arg2, arg3})
	stub := fake.CreateServiceInstanceStub
	fakeReturns := fake.createServiceInstanceReturns
	fake.recordInvocation("CreateServiceInstance", []interface{}{arg1, arg2, arg3})
	fake.createServiceInstanceMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CFServiceInstanceRepository) CreateServiceInstanceCallCount() int {
	fake.createServiceInstanceMutex.RLock()
	defer fake.createServiceInstanceMutex.RUnlock()
	return len(fake.createServiceInstanceArgsForCall)
}

func (fake *CFServiceInstanceRepository) CreateServiceInstanceCalls(stub func(context.Context, authorization.Info, repositories.CreateServiceInstanceMessage) (repositories.ServiceInstanceRecord, error)) {
	fake.createServiceInstanceMutex.Lock()
	defer fake.createServiceInstanceMutex.Unlock()
	fake.CreateServiceInstanceStub = stub
}

func (fake *CFServiceInstanceRepository) CreateServiceInstanceArgsForCall(i int) (context.Context, authorization.Info, repositories.CreateServiceInstanceMessage) {
	fake.createServiceInstanceMutex.RLock()
	defer fake.createServiceInstanceMutex.RUnlock()
	argsForCall := fake.createServiceInstanceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *CFServiceInstanceRepository) CreateServiceInstanceReturns(result1 repositories.ServiceInstanceRecord, result2 error) {
	fake.createServiceInstanceMutex.Lock()
	defer fake.createServiceInstanceMutex.Unlock()
	fake.CreateServiceInstanceStub = nil
	fake.createServiceInstanceReturns = struct {
		result1 repositories.ServiceInstanceRecord
		result2 error
	}{result1, result2}
}

func (fake *CFServiceInstanceRepository) CreateServiceInstanceReturnsOnCall(i int, result1 repositories.ServiceInstanceRecord, result2 error) {
	fake.createServiceInstanceMutex.Lock()
	defer fake.createServiceInstanceMutex.Unlock()
	fake.CreateServiceInstanceStub = nil
	if fake.createServiceInstanceReturnsOnCall == nil {
		fake.createServiceInstanceReturnsOnCall = make(map[int]struct {
			result1 repositories.ServiceInstanceRecord
			result2 error
		})
	}
	fake.createServiceInstanceReturnsOnCall[i] = struct {
		result1 repositories.ServiceInstanceRecord
		result2 error
	}{result1, result2}
}

func (fake *CFServiceInstanceRepository) DeleteServiceInstance(arg1 context.Context, arg2 authorization.Info, arg3 repositories.DeleteServiceInstanceMessage) error {
	fake.deleteServiceInstanceMutex.Lock()
	ret, specificReturn := fake.deleteServiceInstanceReturnsOnCall[len(fake.deleteServiceInstanceArgsForCall)]
	fake.deleteServiceInstanceArgsForCall = append(fake.deleteServiceInstanceArgsForCall, struct {
		arg1 context.Context
		arg2 authorization.Info
		arg3 repositories.DeleteServiceInstanceMessage
	}{arg1, arg2, arg3})
	stub := fake.DeleteServiceInstanceStub
	fakeReturns := fake.deleteServiceInstanceReturns
	fake.recordInvocation("DeleteServiceInstance", []interface{}{arg1, arg2, arg3})
	fake.deleteServiceInstanceMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *CFServiceInstanceRepository) DeleteServiceInstanceCallCount() int {
	fake.deleteServiceInstanceMutex.RLock()
	defer fake.deleteServiceInstanceMutex.RUnlock()
	return len(fake.deleteServiceInstanceArgsForCall)
}

func (fake *CFServiceInstanceRepository) DeleteServiceInstanceCalls(stub func(context.Context, authorization.Info, repositories.DeleteServiceInstanceMessage) error) {
	fake.deleteServiceInstanceMutex.Lock()
	defer fake.deleteServiceInstanceMutex.Unlock()
	fake.DeleteServiceInstanceStub = stub
}

func (fake *CFServiceInstanceRepository) DeleteServiceInstanceArgsForCall(i int) (context.Context, authorization.Info, repositories.DeleteServiceInstanceMessage) {
	fake.deleteServiceInstanceMutex.RLock()
	defer fake.deleteServiceInstanceMutex.RUnlock()
	argsForCall := fake.deleteServiceInstanceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *CFServiceInstanceRepository) DeleteServiceInstanceReturns(result1 error) {
	fake.deleteServiceInstanceMutex.Lock()
	defer fake.deleteServiceInstanceMutex.Unlock()
	fake.DeleteServiceInstanceStub = nil
	fake.deleteServiceInstanceReturns = struct {
		result1 error
	}{result1}
}

func (fake *CFServiceInstanceRepository) DeleteServiceInstanceReturnsOnCall(i int, result1 error) {
	fake.deleteServiceInstanceMutex.Lock()
	defer fake.deleteServiceInstanceMutex.Unlock()
	fake.DeleteServiceInstanceStub = nil
	if fake.deleteServiceInstanceReturnsOnCall == nil {
		fake.deleteServiceInstanceReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteServiceInstanceReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *CFServiceInstanceRepository) GetServiceInstance(arg1 context.Context, arg2 authorization.Info, arg3 string) (repositories.ServiceInstanceRecord, error) {
	fake.getServiceInstanceMutex.Lock()
	ret, specificReturn := fake.getServiceInstanceReturnsOnCall[len(fake.getServiceInstanceArgsForCall)]
	fake.getServiceInstanceArgsForCall = append(fake.getServiceInstanceArgsForCall, struct {
		arg1 context.Context
		arg2 authorization.Info
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.GetServiceInstanceStub
	fakeReturns := fake.getServiceInstanceReturns
	fake.recordInvocation("GetServiceInstance", []interface{}{arg1, arg2, arg3})
	fake.getServiceInstanceMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CFServiceInstanceRepository) GetServiceInstanceCallCount() int {
	fake.getServiceInstanceMutex.RLock()
	defer fake.getServiceInstanceMutex.RUnlock()
	return len(fake.getServiceInstanceArgsForCall)
}

func (fake *CFServiceInstanceRepository) GetServiceInstanceCalls(stub func(context.Context, authorization.Info, string) (repositories.ServiceInstanceRecord, error)) {
	fake.getServiceInstanceMutex.Lock()
	defer fake.getServiceInstanceMutex.Unlock()
	fake.GetServiceInstanceStub = stub
}

func (fake *CFServiceInstanceRepository) GetServiceInstanceArgsForCall(i int) (context.Context, authorization.Info, string) {
	fake.getServiceInstanceMutex.RLock()
	defer fake.getServiceInstanceMutex.RUnlock()
	argsForCall := fake.getServiceInstanceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *CFServiceInstanceRepository) GetServiceInstanceReturns(result1 repositories.ServiceInstanceRecord, result2 error) {
	fake.getServiceInstanceMutex.Lock()
	defer fake.getServiceInstanceMutex.Unlock()
	fake.GetServiceInstanceStub = nil
	fake.getServiceInstanceReturns = struct {
		result1 repositories.ServiceInstanceRecord
		result2 error
	}{result1, result2}
}

func (fake *CFServiceInstanceRepository) GetServiceInstanceReturnsOnCall(i int, result1 repositories.ServiceInstanceRecord, result2 error) {
	fake.getServiceInstanceMutex.Lock()
	defer fake.getServiceInstanceMutex.Unlock()
	fake.GetServiceInstanceStub = nil
	if fake.getServiceInstanceReturnsOnCall == nil {
		fake.getServiceInstanceReturnsOnCall = make(map[int]struct {
			result1 repositories.ServiceInstanceRecord
			result2 error
		})
	}
	fake.getServiceInstanceReturnsOnCall[i] = struct {
		result1 repositories.ServiceInstanceRecord
		result2 error
	}{result1, result2}
}

func (fake *CFServiceInstanceRepository) ListServiceInstances(arg1 context.Context, arg2 authorization.Info, arg3 repositories.ListServiceInstanceMessage) ([]repositories.ServiceInstanceRecord, error) {
	fake.listServiceInstancesMutex.Lock()
	ret, specificReturn := fake.listServiceInstancesReturnsOnCall[len(fake.listServiceInstancesArgsForCall)]
	fake.listServiceInstancesArgsForCall = append(fake.listServiceInstancesArgsForCall, struct {
		arg1 context.Context
		arg2 authorization.Info
		arg3 repositories.ListServiceInstanceMessage
	}{arg1, arg2, arg3})
	stub := fake.ListServiceInstancesStub
	fakeReturns := fake.listServiceInstancesReturns
	fake.recordInvocation("ListServiceInstances", []interface{}{arg1, arg2, arg3})
	fake.listServiceInstancesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CFServiceInstanceRepository) ListServiceInstancesCallCount() int {
	fake.listServiceInstancesMutex.RLock()
	defer fake.listServiceInstancesMutex.RUnlock()
	return len(fake.listServiceInstancesArgsForCall)
}

func (fake *CFServiceInstanceRepository) ListServiceInstancesCalls(stub func(context.Context, authorization.Info, repositories.ListServiceInstanceMessage) ([]repositories.ServiceInstanceRecord, error)) {
	fake.listServiceInstancesMutex.Lock()
	defer fake.listServiceInstancesMutex.Unlock()
	fake.ListServiceInstancesStub = stub
}

func (fake *CFServiceInstanceRepository) ListServiceInstancesArgsForCall(i int) (context.Context, authorization.Info, repositories.ListServiceInstanceMessage) {
	fake.listServiceInstancesMutex.RLock()
	defer fake.listServiceInstancesMutex.RUnlock()
	argsForCall := fake.listServiceInstancesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *CFServiceInstanceRepository) ListServiceInstancesReturns(result1 []repositories.ServiceInstanceRecord, result2 error) {
	fake.listServiceInstancesMutex.Lock()
	defer fake.listServiceInstancesMutex.Unlock()
	fake.ListServiceInstancesStub = nil
	fake.listServiceInstancesReturns = struct {
		result1 []repositories.ServiceInstanceRecord
		result2 error
	}{result1, result2}
}

func (fake *CFServiceInstanceRepository) ListServiceInstancesReturnsOnCall(i int, result1 []repositories.ServiceInstanceRecord, result2 error) {
	fake.listServiceInstancesMutex.Lock()
	defer fake.listServiceInstancesMutex.Unlock()
	fake.ListServiceInstancesStub = nil
	if fake.listServiceInstancesReturnsOnCall == nil {
		fake.listServiceInstancesReturnsOnCall = make(map[int]struct {
			result1 []repositories.ServiceInstanceRecord
			result2 error
		})
	}
	fake.listServiceInstancesReturnsOnCall[i] = struct {
		result1 []repositories.ServiceInstanceRecord
		result2 error
	}{result1, result2}
}

func (fake *CFServiceInstanceRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createServiceInstanceMutex.RLock()
	defer fake.createServiceInstanceMutex.RUnlock()
	fake.deleteServiceInstanceMutex.RLock()
	defer fake.deleteServiceInstanceMutex.RUnlock()
	fake.getServiceInstanceMutex.RLock()
	defer fake.getServiceInstanceMutex.RUnlock()
	fake.listServiceInstancesMutex.RLock()
	defer fake.listServiceInstancesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *CFServiceInstanceRepository) recordInvocation(key string, args []interface{}) {
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

var _ apis.CFServiceInstanceRepository = new(CFServiceInstanceRepository)
