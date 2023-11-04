/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	venusv1 "github.com/huweihuang/venus/pkg/apis/venus/v1"
	versioned "github.com/huweihuang/venus/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/huweihuang/venus/pkg/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/huweihuang/venus/pkg/generated/listers/venus/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// MemcacheInformer provides access to a shared informer and lister for
// Memcaches.
type MemcacheInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.MemcacheLister
}

type memcacheInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewMemcacheInformer constructs a new informer for Memcache type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMemcacheInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredMemcacheInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredMemcacheInformer constructs a new informer for Memcache type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredMemcacheInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.VenusV1().Memcaches(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.VenusV1().Memcaches(namespace).Watch(context.TODO(), options)
			},
		},
		&venusv1.Memcache{},
		resyncPeriod,
		indexers,
	)
}

func (f *memcacheInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredMemcacheInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *memcacheInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&venusv1.Memcache{}, f.defaultInformer)
}

func (f *memcacheInformer) Lister() v1.MemcacheLister {
	return v1.NewMemcacheLister(f.Informer().GetIndexer())
}
