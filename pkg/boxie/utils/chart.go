// Package utils is a useful utils that boxie use. Methods are usually exported
package utils

import (
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/cli-runtime/pkg/resource"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/restmapper"
	"k8s.io/kubectl/pkg/scheme"
)

// NewRestClient provide a k8s rest client
func NewRestClient(restConfig rest.Config, gv schema.GroupVersion) (rest.Interface, error) {
	restConfig.ContentConfig = resource.UnstructuredPlusDefaultContentConfig()
	restConfig.GroupVersion = &gv
	if len(gv.Group) == 0 {
		restConfig.APIPath = "/api"
	} else {
		restConfig.APIPath = "/apis"
	}

	return rest.RESTClientFor(&restConfig)
}

// CreateRuntimeObject will create k8s runtime object from helm chart (template)
func CreateRuntimeObject(yaml string) (runtime.Object, error) {
	obj, _, err := scheme.Codecs.UniversalDeserializer().Decode([]byte(yaml), nil, nil)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

// CreateRestMapper will create a REST mapper that tracks information about the available resources in the cluster.
func CreateRestMapper(k8sclientSet *kubernetes.Clientset, obj runtime.Object) (*meta.RESTMapping, error) {
	groupResources, err := restmapper.GetAPIGroupResources(k8sclientSet.Discovery())
	if err != nil {
		return nil, err
	}
	rm := restmapper.NewDiscoveryRESTMapper(groupResources)

	gvk := obj.GetObjectKind().GroupVersionKind()
	gk := schema.GroupKind{Group: gvk.Group, Kind: gvk.Kind}
	mapping, err := rm.RESTMapping(gk, gvk.Version)
	if err != nil {
		return nil, err
	}

	return mapping, nil
}

// ConvertHelmRenderToYaml will convert helmcharts replaced render (tempalte) to k8s yaml format
func ConvertHelmRenderToYaml(m map[string]string) []string {
	var b []string
	for _, value := range m {
		b = append(b, value)
	}
	return b
}
