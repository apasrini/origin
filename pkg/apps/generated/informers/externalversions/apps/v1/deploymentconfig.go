// This file was automatically generated by informer-gen

package v1

import (
	apps_v1 "github.com/openshift/origin/pkg/apps/apis/apps/v1"
	clientset "github.com/openshift/origin/pkg/apps/generated/clientset"
	internalinterfaces "github.com/openshift/origin/pkg/apps/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/openshift/origin/pkg/apps/generated/listers/apps/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// DeploymentConfigInformer provides access to a shared informer and lister for
// DeploymentConfigs.
type DeploymentConfigInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.DeploymentConfigLister
}

type deploymentConfigInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

// NewDeploymentConfigInformer constructs a new informer for DeploymentConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDeploymentConfigInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				return client.AppsV1().DeploymentConfigs(namespace).List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				return client.AppsV1().DeploymentConfigs(namespace).Watch(options)
			},
		},
		&apps_v1.DeploymentConfig{},
		resyncPeriod,
		indexers,
	)
}

func defaultDeploymentConfigInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewDeploymentConfigInformer(client, meta_v1.NamespaceAll, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
}

func (f *deploymentConfigInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apps_v1.DeploymentConfig{}, defaultDeploymentConfigInformer)
}

func (f *deploymentConfigInformer) Lister() v1.DeploymentConfigLister {
	return v1.NewDeploymentConfigLister(f.Informer().GetIndexer())
}
