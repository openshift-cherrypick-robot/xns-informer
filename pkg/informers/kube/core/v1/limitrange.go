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

type limitRangeInformer struct {
	factory xnsinformers.InformerFactory
}

var _ informers.LimitRangeInformer = &limitRangeInformer{}

func (f *limitRangeInformer) resource() schema.GroupVersionResource {
	return v1.SchemeGroupVersion.WithResource("limitranges")
}

func (f *limitRangeInformer) Informer() cache.SharedIndexInformer {
	return f.factory.NamespacedResource(f.resource()).Informer()
}

func (f *limitRangeInformer) Lister() listers.LimitRangeLister {
	idx := xnsinformers.NewCacheConverter(f.Informer().GetIndexer(), &v1.LimitRange{})
	return listers.NewLimitRangeLister(idx)
}
