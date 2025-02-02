/*
Copyright Red Hat, Inc.

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

// Code generated by xns-informer-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	time "time"

	internalinterfaces "github.com/maistra/xns-informer/pkg/generated/istio/internalinterfaces"
	informers "github.com/maistra/xns-informer/pkg/informers"
	networkingv1beta1 "istio.io/client-go/pkg/apis/networking/v1beta1"
	versioned "istio.io/client-go/pkg/clientset/versioned"
	v1beta1 "istio.io/client-go/pkg/listers/networking/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ServiceEntryInformer provides access to a shared informer and lister for
// ServiceEntries.
type ServiceEntryInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.ServiceEntryLister
}

type serviceEntryInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespaces       informers.NamespaceSet
}

// NewServiceEntryInformer constructs a new informer for ServiceEntry type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewServiceEntryInformer(client versioned.Interface, namespaces informers.NamespaceSet,
	resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredServiceEntryInformer(client, namespaces, resyncPeriod, indexers, nil)
}

// NewFilteredServiceEntryInformer constructs a new informer for ServiceEntry type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredServiceEntryInformer(client versioned.Interface, namespaces informers.NamespaceSet,
	resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	newInformer := func(namespace string) cache.SharedIndexInformer {
		return cache.NewSharedIndexInformer(
			&cache.ListWatch{
				ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
					if tweakListOptions != nil {
						tweakListOptions(&options)
					}
					return client.NetworkingV1beta1().ServiceEntries(namespace).List(context.TODO(), options)
				},
				WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
					if tweakListOptions != nil {
						tweakListOptions(&options)
					}
					return client.NetworkingV1beta1().ServiceEntries(namespace).Watch(context.TODO(), options)
				},
			},
			&networkingv1beta1.ServiceEntry{},
			resyncPeriod,
			indexers,
		)
	}

	return informers.NewMultiNamespaceInformer(namespaces, resyncPeriod, newInformer)
}

func (f *serviceEntryInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredServiceEntryInformer(client, f.namespaces, resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *serviceEntryInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&networkingv1beta1.ServiceEntry{}, f.defaultInformer)
}

func (f *serviceEntryInformer) Lister() v1beta1.ServiceEntryLister {
	return v1beta1.NewServiceEntryLister(f.Informer().GetIndexer())
}
