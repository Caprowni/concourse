// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/db/lock"
)

type FakeCheckFactory struct {
	AcquireScanningLockStub        func(lager.Logger) (lock.Lock, bool, error)
	acquireScanningLockMutex       sync.RWMutex
	acquireScanningLockArgsForCall []struct {
		arg1 lager.Logger
	}
	acquireScanningLockReturns struct {
		result1 lock.Lock
		result2 bool
		result3 error
	}
	acquireScanningLockReturnsOnCall map[int]struct {
		result1 lock.Lock
		result2 bool
		result3 error
	}
	CheckStub        func(int) (db.Check, bool, error)
	checkMutex       sync.RWMutex
	checkArgsForCall []struct {
		arg1 int
	}
	checkReturns struct {
		result1 db.Check
		result2 bool
		result3 error
	}
	checkReturnsOnCall map[int]struct {
		result1 db.Check
		result2 bool
		result3 error
	}
	CreateCheckStub        func(int, int, int, int, bool, atc.Plan) (db.Check, bool, error)
	createCheckMutex       sync.RWMutex
	createCheckArgsForCall []struct {
		arg1 int
		arg2 int
		arg3 int
		arg4 int
		arg5 bool
		arg6 atc.Plan
	}
	createCheckReturns struct {
		result1 db.Check
		result2 bool
		result3 error
	}
	createCheckReturnsOnCall map[int]struct {
		result1 db.Check
		result2 bool
		result3 error
	}
	NotifyCheckerStub        func() error
	notifyCheckerMutex       sync.RWMutex
	notifyCheckerArgsForCall []struct {
	}
	notifyCheckerReturns struct {
		result1 error
	}
	notifyCheckerReturnsOnCall map[int]struct {
		result1 error
	}
	ResourceTypesStub        func() ([]db.ResourceType, error)
	resourceTypesMutex       sync.RWMutex
	resourceTypesArgsForCall []struct {
	}
	resourceTypesReturns struct {
		result1 []db.ResourceType
		result2 error
	}
	resourceTypesReturnsOnCall map[int]struct {
		result1 []db.ResourceType
		result2 error
	}
	ResourcesStub        func() ([]db.Resource, error)
	resourcesMutex       sync.RWMutex
	resourcesArgsForCall []struct {
	}
	resourcesReturns struct {
		result1 []db.Resource
		result2 error
	}
	resourcesReturnsOnCall map[int]struct {
		result1 []db.Resource
		result2 error
	}
	StartedChecksStub        func() ([]db.Check, error)
	startedChecksMutex       sync.RWMutex
	startedChecksArgsForCall []struct {
	}
	startedChecksReturns struct {
		result1 []db.Check
		result2 error
	}
	startedChecksReturnsOnCall map[int]struct {
		result1 []db.Check
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCheckFactory) AcquireScanningLock(arg1 lager.Logger) (lock.Lock, bool, error) {
	fake.acquireScanningLockMutex.Lock()
	ret, specificReturn := fake.acquireScanningLockReturnsOnCall[len(fake.acquireScanningLockArgsForCall)]
	fake.acquireScanningLockArgsForCall = append(fake.acquireScanningLockArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	fake.recordInvocation("AcquireScanningLock", []interface{}{arg1})
	fake.acquireScanningLockMutex.Unlock()
	if fake.AcquireScanningLockStub != nil {
		return fake.AcquireScanningLockStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.acquireScanningLockReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeCheckFactory) AcquireScanningLockCallCount() int {
	fake.acquireScanningLockMutex.RLock()
	defer fake.acquireScanningLockMutex.RUnlock()
	return len(fake.acquireScanningLockArgsForCall)
}

func (fake *FakeCheckFactory) AcquireScanningLockCalls(stub func(lager.Logger) (lock.Lock, bool, error)) {
	fake.acquireScanningLockMutex.Lock()
	defer fake.acquireScanningLockMutex.Unlock()
	fake.AcquireScanningLockStub = stub
}

func (fake *FakeCheckFactory) AcquireScanningLockArgsForCall(i int) lager.Logger {
	fake.acquireScanningLockMutex.RLock()
	defer fake.acquireScanningLockMutex.RUnlock()
	argsForCall := fake.acquireScanningLockArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCheckFactory) AcquireScanningLockReturns(result1 lock.Lock, result2 bool, result3 error) {
	fake.acquireScanningLockMutex.Lock()
	defer fake.acquireScanningLockMutex.Unlock()
	fake.AcquireScanningLockStub = nil
	fake.acquireScanningLockReturns = struct {
		result1 lock.Lock
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCheckFactory) AcquireScanningLockReturnsOnCall(i int, result1 lock.Lock, result2 bool, result3 error) {
	fake.acquireScanningLockMutex.Lock()
	defer fake.acquireScanningLockMutex.Unlock()
	fake.AcquireScanningLockStub = nil
	if fake.acquireScanningLockReturnsOnCall == nil {
		fake.acquireScanningLockReturnsOnCall = make(map[int]struct {
			result1 lock.Lock
			result2 bool
			result3 error
		})
	}
	fake.acquireScanningLockReturnsOnCall[i] = struct {
		result1 lock.Lock
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCheckFactory) Check(arg1 int) (db.Check, bool, error) {
	fake.checkMutex.Lock()
	ret, specificReturn := fake.checkReturnsOnCall[len(fake.checkArgsForCall)]
	fake.checkArgsForCall = append(fake.checkArgsForCall, struct {
		arg1 int
	}{arg1})
	fake.recordInvocation("Check", []interface{}{arg1})
	fake.checkMutex.Unlock()
	if fake.CheckStub != nil {
		return fake.CheckStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.checkReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeCheckFactory) CheckCallCount() int {
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	return len(fake.checkArgsForCall)
}

func (fake *FakeCheckFactory) CheckCalls(stub func(int) (db.Check, bool, error)) {
	fake.checkMutex.Lock()
	defer fake.checkMutex.Unlock()
	fake.CheckStub = stub
}

func (fake *FakeCheckFactory) CheckArgsForCall(i int) int {
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	argsForCall := fake.checkArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCheckFactory) CheckReturns(result1 db.Check, result2 bool, result3 error) {
	fake.checkMutex.Lock()
	defer fake.checkMutex.Unlock()
	fake.CheckStub = nil
	fake.checkReturns = struct {
		result1 db.Check
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCheckFactory) CheckReturnsOnCall(i int, result1 db.Check, result2 bool, result3 error) {
	fake.checkMutex.Lock()
	defer fake.checkMutex.Unlock()
	fake.CheckStub = nil
	if fake.checkReturnsOnCall == nil {
		fake.checkReturnsOnCall = make(map[int]struct {
			result1 db.Check
			result2 bool
			result3 error
		})
	}
	fake.checkReturnsOnCall[i] = struct {
		result1 db.Check
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCheckFactory) CreateCheck(arg1 int, arg2 int, arg3 int, arg4 int, arg5 bool, arg6 atc.Plan) (db.Check, bool, error) {
	fake.createCheckMutex.Lock()
	ret, specificReturn := fake.createCheckReturnsOnCall[len(fake.createCheckArgsForCall)]
	fake.createCheckArgsForCall = append(fake.createCheckArgsForCall, struct {
		arg1 int
		arg2 int
		arg3 int
		arg4 int
		arg5 bool
		arg6 atc.Plan
	}{arg1, arg2, arg3, arg4, arg5, arg6})
	fake.recordInvocation("CreateCheck", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6})
	fake.createCheckMutex.Unlock()
	if fake.CreateCheckStub != nil {
		return fake.CreateCheckStub(arg1, arg2, arg3, arg4, arg5, arg6)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.createCheckReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeCheckFactory) CreateCheckCallCount() int {
	fake.createCheckMutex.RLock()
	defer fake.createCheckMutex.RUnlock()
	return len(fake.createCheckArgsForCall)
}

func (fake *FakeCheckFactory) CreateCheckCalls(stub func(int, int, int, int, bool, atc.Plan) (db.Check, bool, error)) {
	fake.createCheckMutex.Lock()
	defer fake.createCheckMutex.Unlock()
	fake.CreateCheckStub = stub
}

func (fake *FakeCheckFactory) CreateCheckArgsForCall(i int) (int, int, int, int, bool, atc.Plan) {
	fake.createCheckMutex.RLock()
	defer fake.createCheckMutex.RUnlock()
	argsForCall := fake.createCheckArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6
}

func (fake *FakeCheckFactory) CreateCheckReturns(result1 db.Check, result2 bool, result3 error) {
	fake.createCheckMutex.Lock()
	defer fake.createCheckMutex.Unlock()
	fake.CreateCheckStub = nil
	fake.createCheckReturns = struct {
		result1 db.Check
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCheckFactory) CreateCheckReturnsOnCall(i int, result1 db.Check, result2 bool, result3 error) {
	fake.createCheckMutex.Lock()
	defer fake.createCheckMutex.Unlock()
	fake.CreateCheckStub = nil
	if fake.createCheckReturnsOnCall == nil {
		fake.createCheckReturnsOnCall = make(map[int]struct {
			result1 db.Check
			result2 bool
			result3 error
		})
	}
	fake.createCheckReturnsOnCall[i] = struct {
		result1 db.Check
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCheckFactory) NotifyChecker() error {
	fake.notifyCheckerMutex.Lock()
	ret, specificReturn := fake.notifyCheckerReturnsOnCall[len(fake.notifyCheckerArgsForCall)]
	fake.notifyCheckerArgsForCall = append(fake.notifyCheckerArgsForCall, struct {
	}{})
	fake.recordInvocation("NotifyChecker", []interface{}{})
	fake.notifyCheckerMutex.Unlock()
	if fake.NotifyCheckerStub != nil {
		return fake.NotifyCheckerStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.notifyCheckerReturns
	return fakeReturns.result1
}

func (fake *FakeCheckFactory) NotifyCheckerCallCount() int {
	fake.notifyCheckerMutex.RLock()
	defer fake.notifyCheckerMutex.RUnlock()
	return len(fake.notifyCheckerArgsForCall)
}

func (fake *FakeCheckFactory) NotifyCheckerCalls(stub func() error) {
	fake.notifyCheckerMutex.Lock()
	defer fake.notifyCheckerMutex.Unlock()
	fake.NotifyCheckerStub = stub
}

func (fake *FakeCheckFactory) NotifyCheckerReturns(result1 error) {
	fake.notifyCheckerMutex.Lock()
	defer fake.notifyCheckerMutex.Unlock()
	fake.NotifyCheckerStub = nil
	fake.notifyCheckerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCheckFactory) NotifyCheckerReturnsOnCall(i int, result1 error) {
	fake.notifyCheckerMutex.Lock()
	defer fake.notifyCheckerMutex.Unlock()
	fake.NotifyCheckerStub = nil
	if fake.notifyCheckerReturnsOnCall == nil {
		fake.notifyCheckerReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.notifyCheckerReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCheckFactory) ResourceTypes() ([]db.ResourceType, error) {
	fake.resourceTypesMutex.Lock()
	ret, specificReturn := fake.resourceTypesReturnsOnCall[len(fake.resourceTypesArgsForCall)]
	fake.resourceTypesArgsForCall = append(fake.resourceTypesArgsForCall, struct {
	}{})
	fake.recordInvocation("ResourceTypes", []interface{}{})
	fake.resourceTypesMutex.Unlock()
	if fake.ResourceTypesStub != nil {
		return fake.ResourceTypesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.resourceTypesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCheckFactory) ResourceTypesCallCount() int {
	fake.resourceTypesMutex.RLock()
	defer fake.resourceTypesMutex.RUnlock()
	return len(fake.resourceTypesArgsForCall)
}

func (fake *FakeCheckFactory) ResourceTypesCalls(stub func() ([]db.ResourceType, error)) {
	fake.resourceTypesMutex.Lock()
	defer fake.resourceTypesMutex.Unlock()
	fake.ResourceTypesStub = stub
}

func (fake *FakeCheckFactory) ResourceTypesReturns(result1 []db.ResourceType, result2 error) {
	fake.resourceTypesMutex.Lock()
	defer fake.resourceTypesMutex.Unlock()
	fake.ResourceTypesStub = nil
	fake.resourceTypesReturns = struct {
		result1 []db.ResourceType
		result2 error
	}{result1, result2}
}

func (fake *FakeCheckFactory) ResourceTypesReturnsOnCall(i int, result1 []db.ResourceType, result2 error) {
	fake.resourceTypesMutex.Lock()
	defer fake.resourceTypesMutex.Unlock()
	fake.ResourceTypesStub = nil
	if fake.resourceTypesReturnsOnCall == nil {
		fake.resourceTypesReturnsOnCall = make(map[int]struct {
			result1 []db.ResourceType
			result2 error
		})
	}
	fake.resourceTypesReturnsOnCall[i] = struct {
		result1 []db.ResourceType
		result2 error
	}{result1, result2}
}

func (fake *FakeCheckFactory) Resources() ([]db.Resource, error) {
	fake.resourcesMutex.Lock()
	ret, specificReturn := fake.resourcesReturnsOnCall[len(fake.resourcesArgsForCall)]
	fake.resourcesArgsForCall = append(fake.resourcesArgsForCall, struct {
	}{})
	fake.recordInvocation("Resources", []interface{}{})
	fake.resourcesMutex.Unlock()
	if fake.ResourcesStub != nil {
		return fake.ResourcesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.resourcesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCheckFactory) ResourcesCallCount() int {
	fake.resourcesMutex.RLock()
	defer fake.resourcesMutex.RUnlock()
	return len(fake.resourcesArgsForCall)
}

func (fake *FakeCheckFactory) ResourcesCalls(stub func() ([]db.Resource, error)) {
	fake.resourcesMutex.Lock()
	defer fake.resourcesMutex.Unlock()
	fake.ResourcesStub = stub
}

func (fake *FakeCheckFactory) ResourcesReturns(result1 []db.Resource, result2 error) {
	fake.resourcesMutex.Lock()
	defer fake.resourcesMutex.Unlock()
	fake.ResourcesStub = nil
	fake.resourcesReturns = struct {
		result1 []db.Resource
		result2 error
	}{result1, result2}
}

func (fake *FakeCheckFactory) ResourcesReturnsOnCall(i int, result1 []db.Resource, result2 error) {
	fake.resourcesMutex.Lock()
	defer fake.resourcesMutex.Unlock()
	fake.ResourcesStub = nil
	if fake.resourcesReturnsOnCall == nil {
		fake.resourcesReturnsOnCall = make(map[int]struct {
			result1 []db.Resource
			result2 error
		})
	}
	fake.resourcesReturnsOnCall[i] = struct {
		result1 []db.Resource
		result2 error
	}{result1, result2}
}

func (fake *FakeCheckFactory) StartedChecks() ([]db.Check, error) {
	fake.startedChecksMutex.Lock()
	ret, specificReturn := fake.startedChecksReturnsOnCall[len(fake.startedChecksArgsForCall)]
	fake.startedChecksArgsForCall = append(fake.startedChecksArgsForCall, struct {
	}{})
	fake.recordInvocation("StartedChecks", []interface{}{})
	fake.startedChecksMutex.Unlock()
	if fake.StartedChecksStub != nil {
		return fake.StartedChecksStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.startedChecksReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCheckFactory) StartedChecksCallCount() int {
	fake.startedChecksMutex.RLock()
	defer fake.startedChecksMutex.RUnlock()
	return len(fake.startedChecksArgsForCall)
}

func (fake *FakeCheckFactory) StartedChecksCalls(stub func() ([]db.Check, error)) {
	fake.startedChecksMutex.Lock()
	defer fake.startedChecksMutex.Unlock()
	fake.StartedChecksStub = stub
}

func (fake *FakeCheckFactory) StartedChecksReturns(result1 []db.Check, result2 error) {
	fake.startedChecksMutex.Lock()
	defer fake.startedChecksMutex.Unlock()
	fake.StartedChecksStub = nil
	fake.startedChecksReturns = struct {
		result1 []db.Check
		result2 error
	}{result1, result2}
}

func (fake *FakeCheckFactory) StartedChecksReturnsOnCall(i int, result1 []db.Check, result2 error) {
	fake.startedChecksMutex.Lock()
	defer fake.startedChecksMutex.Unlock()
	fake.StartedChecksStub = nil
	if fake.startedChecksReturnsOnCall == nil {
		fake.startedChecksReturnsOnCall = make(map[int]struct {
			result1 []db.Check
			result2 error
		})
	}
	fake.startedChecksReturnsOnCall[i] = struct {
		result1 []db.Check
		result2 error
	}{result1, result2}
}

func (fake *FakeCheckFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.acquireScanningLockMutex.RLock()
	defer fake.acquireScanningLockMutex.RUnlock()
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	fake.createCheckMutex.RLock()
	defer fake.createCheckMutex.RUnlock()
	fake.notifyCheckerMutex.RLock()
	defer fake.notifyCheckerMutex.RUnlock()
	fake.resourceTypesMutex.RLock()
	defer fake.resourceTypesMutex.RUnlock()
	fake.resourcesMutex.RLock()
	defer fake.resourcesMutex.RUnlock()
	fake.startedChecksMutex.RLock()
	defer fake.startedChecksMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCheckFactory) recordInvocation(key string, args []interface{}) {
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

var _ db.CheckFactory = new(FakeCheckFactory)
