// Package services contains buisness-logic methods of the models
package services

import (
	"context"
	"github.com/twelvee/boxie/pkg/boxie/structs"
	"github.com/twelvee/boxie/pkg/boxie/utils"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"os"
)

// NewClusterService creates a new ClusterService
func NewClusterService() structs.ClusterService {
	return structs.ClusterService{
		TestConnection: testConnection,
	}
}

func testConnection(cluster structs.Cluster) (bool, error) {
	tempConfigPath := "/tmp/kubeconfig_" + utils.GetShortID(6)
	err := os.WriteFile(tempConfigPath, []byte(cluster.Kubeconfig), 0644)
	defer os.Remove(tempConfigPath)
	config, err := GetConfigFromKubeconfig("default", tempConfigPath)
	if err != nil {
		return false, err
	}
	cl, err := kubernetes.NewForConfig(config)
	if err != nil {
		return false, err
	}
	list, err := cl.CoreV1().Nodes().List(context.Background(), v1.ListOptions{})
	if err != nil {
		return false, err
	}
	if len(list.Items) > 0 {
		return true, nil
	}
	return false, nil
}
