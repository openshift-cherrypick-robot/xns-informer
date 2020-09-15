// Code generated by xns-informer-gen. DO NOT EDIT.
package v1

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	informers "k8s.io/client-go/informers/core/v1"
	listers "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
)

type secretInformer struct {
	factory xnsinformers.InformerFactory
}

var _ informers.SecretInformer = &secretInformer{}

func (f *secretInformer) resource() schema.GroupVersionResource {
	return v1.SchemeGroupVersion.WithResource("secrets")
}

func (f *secretInformer) Informer() cache.SharedIndexInformer {
	return f.factory.NamespacedResource(f.resource()).Informer()
}

func (f *secretInformer) Lister() listers.SecretLister {
	idx := xnsinformers.NewCacheConverter(f.Informer().GetIndexer(), &v1.Secret{})
	return listers.NewSecretLister(idx)
}
