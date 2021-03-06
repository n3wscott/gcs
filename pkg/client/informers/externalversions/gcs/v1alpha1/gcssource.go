/*
Copyright 2018 The Kubernetes Authors.

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

package v1alpha1

import (
	time "time"

	gcsv1alpha1 "github.com/vaikas-google/gcs/pkg/apis/gcs/v1alpha1"
	versioned "github.com/vaikas-google/gcs/pkg/client/clientset/versioned"
	internalinterfaces "github.com/vaikas-google/gcs/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/vaikas-google/gcs/pkg/client/listers/gcs/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// GCSSourceInformer provides access to a shared informer and lister for
// GCSSources.
type GCSSourceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.GCSSourceLister
}

type gCSSourceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewGCSSourceInformer constructs a new informer for GCSSource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewGCSSourceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredGCSSourceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredGCSSourceInformer constructs a new informer for GCSSource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredGCSSourceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SourcesV1alpha1().GCSSources(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SourcesV1alpha1().GCSSources(namespace).Watch(options)
			},
		},
		&gcsv1alpha1.GCSSource{},
		resyncPeriod,
		indexers,
	)
}

func (f *gCSSourceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredGCSSourceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *gCSSourceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&gcsv1alpha1.GCSSource{}, f.defaultInformer)
}

func (f *gCSSourceInformer) Lister() v1alpha1.GCSSourceLister {
	return v1alpha1.NewGCSSourceLister(f.Informer().GetIndexer())
}
