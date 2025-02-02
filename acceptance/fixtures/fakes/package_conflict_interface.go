package fakes

import (
	"context"
	"sync"

	appsv1 "k8s.io/api/apps/v1"

	v1 "k8s.io/api/core/v1"
)

type PackageConflictInterface struct {
	ListDaemonsetsCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Ctx context.Context
		}
		Returns struct {
			DaemonSetList *appsv1.DaemonSetList
			Error         error
		}
		Stub func(context.Context) (*appsv1.DaemonSetList, error)
	}
	ListNodesCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Ctx context.Context
		}
		Returns struct {
			NodeList *v1.NodeList
			Error    error
		}
		Stub func(context.Context) (*v1.NodeList, error)
	}
}

func (f *PackageConflictInterface) ListDaemonsets(param1 context.Context) (*appsv1.DaemonSetList, error) {
	f.ListDaemonsetsCall.mutex.Lock()
	defer f.ListDaemonsetsCall.mutex.Unlock()
	f.ListDaemonsetsCall.CallCount++
	f.ListDaemonsetsCall.Receives.Ctx = param1
	if f.ListDaemonsetsCall.Stub != nil {
		return f.ListDaemonsetsCall.Stub(param1)
	}
	return f.ListDaemonsetsCall.Returns.DaemonSetList, f.ListDaemonsetsCall.Returns.Error
}
func (f *PackageConflictInterface) ListNodes(param1 context.Context) (*v1.NodeList, error) {
	f.ListNodesCall.mutex.Lock()
	defer f.ListNodesCall.mutex.Unlock()
	f.ListNodesCall.CallCount++
	f.ListNodesCall.Receives.Ctx = param1
	if f.ListNodesCall.Stub != nil {
		return f.ListNodesCall.Stub(param1)
	}
	return f.ListNodesCall.Returns.NodeList, f.ListNodesCall.Returns.Error
}
