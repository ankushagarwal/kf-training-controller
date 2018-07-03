// Code generated by informer-gen. DO NOT EDIT.

package v1alpha2

import (
	time "time"

	kubeflow_v1alpha2 "github.com/ankushagarwal/kf-training-controller/pkg/apis/kubeflow/v1alpha2"
	versioned "github.com/ankushagarwal/kf-training-controller/pkg/client/clientset/versioned"
	internalinterfaces "github.com/ankushagarwal/kf-training-controller/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha2 "github.com/ankushagarwal/kf-training-controller/pkg/client/listers/kubeflow/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PyTorchJobInformer provides access to a shared informer and lister for
// PyTorchJobs.
type PyTorchJobInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha2.PyTorchJobLister
}

type pyTorchJobInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPyTorchJobInformer constructs a new informer for PyTorchJob type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPyTorchJobInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPyTorchJobInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPyTorchJobInformer constructs a new informer for PyTorchJob type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPyTorchJobInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeflowV1alpha2().PyTorchJobs(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeflowV1alpha2().PyTorchJobs(namespace).Watch(options)
			},
		},
		&kubeflow_v1alpha2.PyTorchJob{},
		resyncPeriod,
		indexers,
	)
}

func (f *pyTorchJobInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPyTorchJobInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *pyTorchJobInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&kubeflow_v1alpha2.PyTorchJob{}, f.defaultInformer)
}

func (f *pyTorchJobInformer) Lister() v1alpha2.PyTorchJobLister {
	return v1alpha2.NewPyTorchJobLister(f.Informer().GetIndexer())
}
