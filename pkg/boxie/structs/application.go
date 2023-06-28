// Package structs contain every boxie public structs
package structs

import (
	"k8s.io/client-go/kubernetes"
)

// Application is your box application in a struct
type Application struct {
	Name  string `toml:"name"`
	Chart string `toml:"chart"`
}

// ApplicationService is a public ApplicationService
type ApplicationService struct {
	ValidateApplications          func([]Application) []string
	ExpandApplications            func([]Application) []Application
	DescribePod                   func(*kubernetes.Clientset, string, string) error
	DescribePodTemplate           func(*kubernetes.Clientset, string, string) error
	DescribeReplicationController func(*kubernetes.Clientset, string, string) error
	DescribeReplicaSet            func(*kubernetes.Clientset, string, string) error
	DescribeDeployment            func(*kubernetes.Clientset, string, string) error
	DescribeControllerRevision    func(*kubernetes.Clientset, string, string) error
	DescribeDaemonSet             func(*kubernetes.Clientset, string, string) error
	DescribeStatefulSet           func(*kubernetes.Clientset, string, string) error
	DescribeJob                   func(*kubernetes.Clientset, string, string) error
	DescribeCronjob               func(*kubernetes.Clientset, string, string) error
	DescribeHPA                   func(*kubernetes.Clientset, string, string) error
	DescribeService               func(*kubernetes.Clientset, string, string) error
	DescribeIngress               func(*kubernetes.Clientset, string, string) error
}

// GetApplicationAliases return a slice of application model name aliases
func GetApplicationAliases() []string {
	return []string{"application", "applications", "apps"}
}

const (
	KIND_POD                    string = "Pod"
	KIND_POD_TEMPLATE           string = "PodTemplate"
	KIND_REPLICATION_CONTROLLER string = "ReplicationController"
	KIND_REPLICA_SET            string = "ReplicaSet"
	KIND_DEPLOYMENT             string = "Deployment"
	KIND_CONTROLLER_REVISION    string = "ControllerRevision"
	KIND_DAEMON_SET             string = "DaemonSet"
	KIND_STATEFUL_SET           string = "StatefulSet"
	KIND_JOB                    string = "Job"
	KIND_CRONJOB                string = "CronJob"
	KIND_HPA                    string = "HorizontalPodAutoscaler"
	KIND_SERVICE                string = "Service"
	KIND_INGRESS                string = "Ingress"
)
