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

type podInformer struct {
	factory xnsinformers.InformerFactory
}

var _ informers.PodInformer = &podInformer{}

func (f *podInformer) resource() schema.GroupVersionResource {
	return v1.SchemeGroupVersion.WithResource("pods")
}

func (f *podInformer) Informer() cache.SharedIndexInformer {
	return f.factory.NamespacedResource(f.resource()).Informer()
}

func (f *podInformer) Lister() listers.PodLister {
	idx := xnsinformers.NewCacheConverter(f.Informer().GetIndexer(), &v1.Pod{})
	return listers.NewPodLister(idx)
}
