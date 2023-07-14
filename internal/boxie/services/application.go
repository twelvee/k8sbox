// Package services contains buisness-logic methods of the models
package services

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/twelvee/boxie/pkg/boxie/structs"
	"github.com/twelvee/boxie/pkg/boxie/utils"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// NewApplicationService creates a new ApplicationService
func NewApplicationService() structs.ApplicationService {
	return structs.ApplicationService{
		ValidateApplications:          validateApplications,
		ExpandApplications:            ExpandApplications,
		Describe:                      describe,
		DescribePod:                   describePod,
		DescribePodTemplate:           describePodTemplate,
		DescribeReplicationController: describeReplicationController,
		DescribeReplicaSet:            describeReplicaSet,
		DescribeDeployment:            describeDeployment,
		DescribeStatefulSet:           describeStatefulSet,
		DescribeControllerRevision:    describeControllerRevision,
		DescribeDaemonSet:             describeDaemonSet,
		DescribeJob:                   describeJob,
		DescribeCronjob:               describeCronjob,
		DescribeHPA:                   describeHPA,
		DescribeService:               describeService,
		DescribeIngress:               describeIngress,
	}
}

func describe(k8sclient *kubernetes.Clientset, kind, namespace string, name string) (structs.ApplicationRuntimeData, error) {
	var runtimeData structs.ApplicationRuntimeData
	var describeFunc func(*kubernetes.Clientset, string, string) (structs.ApplicationRuntimeData, error)
	switch kind {
	case structs.KIND_POD:
		describeFunc = describePod
	case structs.KIND_POD_TEMPLATE:
		describeFunc = describePodTemplate
	case structs.KIND_REPLICATION_CONTROLLER:
		describeFunc = describeReplicationController
	case structs.KIND_REPLICA_SET:
		describeFunc = describeReplicaSet
	case structs.KIND_DEPLOYMENT:
		describeFunc = describeDeployment
	case structs.KIND_STATEFUL_SET:
		describeFunc = describeStatefulSet
	case structs.KIND_CONTROLLER_REVISION:
		describeFunc = describeControllerRevision
	case structs.KIND_DAEMON_SET:
		describeFunc = describeDaemonSet
	case structs.KIND_JOB:
		describeFunc = describeJob
	case structs.KIND_CRONJOB:
		describeFunc = describeCronjob
	case structs.KIND_HPA:
		describeFunc = describeHPA
	case structs.KIND_SERVICE:
		describeFunc = describeService
	case structs.KIND_INGRESS:
		describeFunc = describeIngress
	}

	if describeFunc != nil {
		rtData, err := describeFunc(k8sclient, namespace, name)
		if err != nil {
			return runtimeData, err
		}

		runtimeData = rtData
	}
	return runtimeData, nil
}

func describePod(k8sclient *kubernetes.Clientset, namespace string, name string) (structs.ApplicationRuntimeData, error) {
	var runtimeData structs.ApplicationRuntimeData
	p, err := k8sclient.CoreV1().Pods(namespace).Get(context.Background(), name, v1.GetOptions{})
	if err != nil {
		return runtimeData, err
	}
	m, err := utils.StructToMap(&p)
	if err != nil {
		return runtimeData, err
	}
	runtimeData.Data = append(runtimeData.Data, m)

	return runtimeData, nil
}

func describePodTemplate(k8sclient *kubernetes.Clientset, namespace string, name string) (structs.ApplicationRuntimeData, error) {
	var runtimeData structs.ApplicationRuntimeData
	pt, err := k8sclient.CoreV1().PodTemplates(namespace).Get(context.Background(), name, v1.GetOptions{})
	if err != nil {
		return runtimeData, err
	}
	m, err := utils.StructToMap(&pt)
	if err != nil {
		return runtimeData, err
	}
	runtimeData.Data = append(runtimeData.Data, m)

	return runtimeData, nil
}

func describeReplicationController(k8sclient *kubernetes.Clientset, namespace string, name string) (structs.ApplicationRuntimeData, error) {
	var runtimeData structs.ApplicationRuntimeData
	rc, err := k8sclient.CoreV1().ReplicationControllers(namespace).Get(context.Background(), name, v1.GetOptions{})
	if err != nil {
		return runtimeData, err
	}
	m, err := utils.StructToMap(&rc)
	if err != nil {
		return runtimeData, err
	}
	runtimeData.Data = append(runtimeData.Data, m)

	return runtimeData, nil
}

func describeReplicaSet(k8sclient *kubernetes.Clientset, namespace string, name string) (structs.ApplicationRuntimeData, error) {
	var runtimeData structs.ApplicationRuntimeData
	rs, err := k8sclient.AppsV1().ReplicaSets(namespace).Get(context.Background(), name, v1.GetOptions{})
	if err != nil {
		return runtimeData, err
	}
	m, err := utils.StructToMap(&rs)
	if err != nil {
		return runtimeData, err
	}
	runtimeData.Data = append(runtimeData.Data, m)

	return runtimeData, nil
}

func describeDeployment(k8sclient *kubernetes.Clientset, namespace string, name string) (structs.ApplicationRuntimeData, error) {
	var runtimeData structs.ApplicationRuntimeData
	d, err := k8sclient.AppsV1().Deployments(namespace).Get(context.Background(), name, v1.GetOptions{})
	if err != nil {
		return runtimeData, err
	}
	m, err := utils.StructToMap(&d)
	if err != nil {
		return runtimeData, err
	}
	runtimeData.Data = append(runtimeData.Data, m)

	return runtimeData, nil
}

func describeStatefulSet(k8sclient *kubernetes.Clientset, namespace string, name string) (structs.ApplicationRuntimeData, error) {
	var runtimeData structs.ApplicationRuntimeData
	ss, err := k8sclient.AppsV1().StatefulSets(namespace).Get(context.Background(), name, v1.GetOptions{})
	if err != nil {
		return runtimeData, err
	}
	m, err := utils.StructToMap(&ss)
	if err != nil {
		return runtimeData, err
	}
	runtimeData.Data = append(runtimeData.Data, m)

	return runtimeData, nil
}

func describeControllerRevision(k8sclient *kubernetes.Clientset, namespace string, name string) (structs.ApplicationRuntimeData, error) {
	var runtimeData structs.ApplicationRuntimeData
	cr, err := k8sclient.AppsV1().ControllerRevisions(namespace).Get(context.Background(), name, v1.GetOptions{})
	if err != nil {
		return runtimeData, err
	}
	m, err := utils.StructToMap(&cr)
	if err != nil {
		return runtimeData, err
	}
	runtimeData.Data = append(runtimeData.Data, m)

	return runtimeData, nil
}

func describeDaemonSet(k8sclient *kubernetes.Clientset, namespace string, name string) (structs.ApplicationRuntimeData, error) {
	var runtimeData structs.ApplicationRuntimeData
	ds, err := k8sclient.AppsV1().DaemonSets(namespace).Get(context.Background(), name, v1.GetOptions{})
	if err != nil {
		return runtimeData, err
	}
	m, err := utils.StructToMap(&ds)
	if err != nil {
		return runtimeData, err
	}
	runtimeData.Data = append(runtimeData.Data, m)

	return runtimeData, nil
}

func describeJob(k8sclient *kubernetes.Clientset, namespace string, name string) (structs.ApplicationRuntimeData, error) {
	var runtimeData structs.ApplicationRuntimeData
	j, err := k8sclient.BatchV1().Jobs(namespace).Get(context.Background(), name, v1.GetOptions{})
	if err != nil {
		return runtimeData, err
	}
	m, err := utils.StructToMap(&j)
	if err != nil {
		return runtimeData, err
	}
	runtimeData.Data = append(runtimeData.Data, m)

	return runtimeData, nil
}

func describeCronjob(k8sclient *kubernetes.Clientset, namespace string, name string) (structs.ApplicationRuntimeData, error) {
	var runtimeData structs.ApplicationRuntimeData
	cj, err := k8sclient.BatchV1().CronJobs(namespace).Get(context.Background(), name, v1.GetOptions{})
	if err != nil {
		return runtimeData, err
	}
	m, err := utils.StructToMap(&cj)
	if err != nil {
		return runtimeData, err
	}
	runtimeData.Data = append(runtimeData.Data, m)

	return runtimeData, nil
}

func describeHPA(k8sclient *kubernetes.Clientset, namespace string, name string) (structs.ApplicationRuntimeData, error) {
	var runtimeData structs.ApplicationRuntimeData
	hpa, err := k8sclient.AutoscalingV1().HorizontalPodAutoscalers(namespace).Get(context.Background(), name, v1.GetOptions{})
	if err != nil {
		return runtimeData, err
	}
	m, err := utils.StructToMap(&hpa)
	if err != nil {
		return runtimeData, err
	}
	runtimeData.Data = append(runtimeData.Data, m)

	return runtimeData, nil
}

func describeService(k8sclient *kubernetes.Clientset, namespace string, name string) (structs.ApplicationRuntimeData, error) {
	var runtimeData structs.ApplicationRuntimeData
	s, err := k8sclient.CoreV1().Services(namespace).Get(context.Background(), name, v1.GetOptions{})
	if err != nil {
		return runtimeData, err
	}
	m, err := utils.StructToMap(&s)
	if err != nil {
		return runtimeData, err
	}
	runtimeData.Data = append(runtimeData.Data, m)

	return runtimeData, nil
}

func describeIngress(k8sclient *kubernetes.Clientset, namespace string, name string) (structs.ApplicationRuntimeData, error) {
	var runtimeData structs.ApplicationRuntimeData
	i, err := k8sclient.NetworkingV1().Ingresses(namespace).Get(context.Background(), name, v1.GetOptions{})
	if err != nil {
		return runtimeData, err
	}
	m, err := utils.StructToMap(&i)
	if err != nil {
		return runtimeData, err
	}
	runtimeData.Data = append(runtimeData.Data, m)

	return runtimeData, nil
}

func validateApplications(applications []structs.Application) []string {
	var messages []string
	for index, application := range applications {
		if len(application.Name) == 0 {
			messages = append(messages, fmt.Sprintf("--> Application %d: Name is missing", index))
		}

		if len(strings.TrimSpace(application.Chart)) == 0 {
			messages = append(messages, fmt.Sprintf("--> Application %d: Chart is missing", index))
		}
	}
	return messages
}

// ExpandApplications expand environment variables in applications array
func ExpandApplications(applications []structs.Application) []structs.Application {
	var newApplications []structs.Application
	for _, a := range applications {
		a.Name = os.ExpandEnv(a.Name)
		a.Chart = os.ExpandEnv(a.Chart)
		newApplications = append(newApplications, a)
	}
	return newApplications
}
