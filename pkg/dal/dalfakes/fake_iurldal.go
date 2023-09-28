// Code generated by counterfeiter. DO NOT EDIT.
package dalfakes

import (
	"sync"
	"url-shortener/pkg/dal"
	"url-shortener/pkg/model"
)

type FakeIURLDAL struct {
	AddURLStub        func(*model.URL) error
	addURLMutex       sync.RWMutex
	addURLArgsForCall []struct {
		arg1 *model.URL
	}
	addURLReturns struct {
		result1 error
	}
	addURLReturnsOnCall map[int]struct {
		result1 error
	}
	FindByShortCodeStub        func(string) (*model.URL, error)
	findByShortCodeMutex       sync.RWMutex
	findByShortCodeArgsForCall []struct {
		arg1 string
	}
	findByShortCodeReturns struct {
		result1 *model.URL
		result2 error
	}
	findByShortCodeReturnsOnCall map[int]struct {
		result1 *model.URL
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeIURLDAL) AddURL(arg1 *model.URL) error {
	fake.addURLMutex.Lock()
	ret, specificReturn := fake.addURLReturnsOnCall[len(fake.addURLArgsForCall)]
	fake.addURLArgsForCall = append(fake.addURLArgsForCall, struct {
		arg1 *model.URL
	}{arg1})
	stub := fake.AddURLStub
	fakeReturns := fake.addURLReturns
	fake.recordInvocation("AddURL", []interface{}{arg1})
	fake.addURLMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIURLDAL) AddURLCallCount() int {
	fake.addURLMutex.RLock()
	defer fake.addURLMutex.RUnlock()
	return len(fake.addURLArgsForCall)
}

func (fake *FakeIURLDAL) AddURLCalls(stub func(*model.URL) error) {
	fake.addURLMutex.Lock()
	defer fake.addURLMutex.Unlock()
	fake.AddURLStub = stub
}

func (fake *FakeIURLDAL) AddURLArgsForCall(i int) *model.URL {
	fake.addURLMutex.RLock()
	defer fake.addURLMutex.RUnlock()
	argsForCall := fake.addURLArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeIURLDAL) AddURLReturns(result1 error) {
	fake.addURLMutex.Lock()
	defer fake.addURLMutex.Unlock()
	fake.AddURLStub = nil
	fake.addURLReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIURLDAL) AddURLReturnsOnCall(i int, result1 error) {
	fake.addURLMutex.Lock()
	defer fake.addURLMutex.Unlock()
	fake.AddURLStub = nil
	if fake.addURLReturnsOnCall == nil {
		fake.addURLReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addURLReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIURLDAL) FindByShortCode(arg1 string) (*model.URL, error) {
	fake.findByShortCodeMutex.Lock()
	ret, specificReturn := fake.findByShortCodeReturnsOnCall[len(fake.findByShortCodeArgsForCall)]
	fake.findByShortCodeArgsForCall = append(fake.findByShortCodeArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.FindByShortCodeStub
	fakeReturns := fake.findByShortCodeReturns
	fake.recordInvocation("FindByShortCode", []interface{}{arg1})
	fake.findByShortCodeMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeIURLDAL) FindByShortCodeCallCount() int {
	fake.findByShortCodeMutex.RLock()
	defer fake.findByShortCodeMutex.RUnlock()
	return len(fake.findByShortCodeArgsForCall)
}

func (fake *FakeIURLDAL) FindByShortCodeCalls(stub func(string) (*model.URL, error)) {
	fake.findByShortCodeMutex.Lock()
	defer fake.findByShortCodeMutex.Unlock()
	fake.FindByShortCodeStub = stub
}

func (fake *FakeIURLDAL) FindByShortCodeArgsForCall(i int) string {
	fake.findByShortCodeMutex.RLock()
	defer fake.findByShortCodeMutex.RUnlock()
	argsForCall := fake.findByShortCodeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeIURLDAL) FindByShortCodeReturns(result1 *model.URL, result2 error) {
	fake.findByShortCodeMutex.Lock()
	defer fake.findByShortCodeMutex.Unlock()
	fake.FindByShortCodeStub = nil
	fake.findByShortCodeReturns = struct {
		result1 *model.URL
		result2 error
	}{result1, result2}
}

func (fake *FakeIURLDAL) FindByShortCodeReturnsOnCall(i int, result1 *model.URL, result2 error) {
	fake.findByShortCodeMutex.Lock()
	defer fake.findByShortCodeMutex.Unlock()
	fake.FindByShortCodeStub = nil
	if fake.findByShortCodeReturnsOnCall == nil {
		fake.findByShortCodeReturnsOnCall = make(map[int]struct {
			result1 *model.URL
			result2 error
		})
	}
	fake.findByShortCodeReturnsOnCall[i] = struct {
		result1 *model.URL
		result2 error
	}{result1, result2}
}

func (fake *FakeIURLDAL) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addURLMutex.RLock()
	defer fake.addURLMutex.RUnlock()
	fake.findByShortCodeMutex.RLock()
	defer fake.findByShortCodeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeIURLDAL) recordInvocation(key string, args []interface{}) {
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

var _ dal.IURLDAL = new(FakeIURLDAL)
