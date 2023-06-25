// Package services contains buisness-logic methods of the models
package services

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/twelvee/k8sbox/pkg/k8sbox/structs"
	"github.com/twelvee/k8sbox/pkg/k8sbox/utils"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// NewApplicationService creates a new ApplicationService
func NewApplicationService() structs.ApplicationService {
	return structs.ApplicationService{
		ValidateApplications:          validateApplications,
		ExpandApplications:            ExpandApplications,
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

func describePod(k8sclient *kubernetes.Clientset, namespace string, name string) error {
	p, err := k8sclient.CoreV1().Pods(namespace).Get(context.Background(), name, v1.GetOptions{})
	if err != nil {
		return err
	}
	var result []string
	result = append(result, fmt.Sprintf("Pod %s:", p.Name))
	if len(p.Spec.Containers) > 0 {
		result = append(result, "Containers:")
		for _, container := range p.Spec.Containers {
			result = append(result, fmt.Sprintf("%s (image: %s)", container.Name, container.Image))
		}
	}
	fmt.Println(strings.Join(result, "\r\n"))
	return nil
}

func describePodTemplate(k8sclient *kubernetes.Clientset, namespace string, name string) error {
	pt, err := k8sclient.CoreV1().PodTemplates(namespace).Get(context.Background(), name, v1.GetOptions{})
	if err != nil {
		return err
	}
	var result []string
	result = append(result, fmt.Sprintf("Pod Template %s:", pt.Name))
	if len(pt.Template.Spec.Containers) > 0 {
		result = append(result, "Containers:")
		for _, container := range pt.Template.Spec.Containers {
			result = append(result, fmt.Sprintf("%s (image: %s)", container.Name, container.Image))
		}
	}
	fmt.Println(strings.Join(result, "\r\n"))
	return nil
}

func describeReplicationController(k8sclient *kubernetes.Clientset, namespace string, name string) error {
	rc, err := k8sclient.CoreV1().ReplicationControllers(namespace).Get(context.Background(), name, v1.GetOptions{})
	if err != nil {
		return err
	}
	var result []string
	result = append(result, fmt.Sprintf("Replication Controller %s:", rc.Name))
	result = append(result, fmt.Sprintf("Replicas: %s", utils.Int32ToString(rc.Status.Replicas)))
	fmt.Println(strings.Join(result, "\r\n"))
	return nil
}

func describeReplicaSet(k8sclient *kubernetes.Clientset, namespace string, name string) error {
	rs, err := k8sclient.AppsV1().ReplicaSets(namespace).Get(context.Background(), name, v1.GetOptions{})
	if err != nil {
		return err
	}
	var result []string
	result = append(result, fmt.Sprintf("Replica Set %s:", rs.Name))
	result = append(result, fmt.Sprintf("Available replicas: %s", utils.Int32ToString(rs.Status.AvailableReplicas)))
	result = append(result, fmt.Sprintf("Replicas: %s", utils.Int32ToString(rs.Status.Replicas)))
	fmt.Println(strings.Join(result, "\r\n"))
	return nil
}

func describeDeployment(k8sclient *kubernetes.Clientset, namespace string, name string) error {
	d, err := k8sclient.AppsV1().Deployments(namespace).Get(context.Background(), name, v1.GetOptions{})
	if err != nil {
		return err
	}
	var result []string
	result = append(result, fmt.Sprintf("Deployment %s:", d.Name))
	result = append(result, fmt.Sprintf("Replicas: %s", utils.Int32ToString(d.Status.Replicas)))
	if len(d.Spec.Template.Spec.Containers) > 0 {
		result = append(result, "Containers:")

		for _, container := range d.Spec.Template.Spec.Containers {
			result = append(result, fmt.Sprintf("%s (image: %s)", container.Name, container.Image))
		}
	}
	fmt.Println(strings.Join(result, "\r\n"))
	return nil
}

func describeStatefulSet(k8sclient *kubernetes.Clientset, namespace string, name string) error {
	ss, err := k8sclient.AppsV1().StatefulSets(namespace).Get(context.Background(), name, v1.GetOptions{})
	if err != nil {
		return err
	}
	var result []string
	result = append(result, fmt.Sprintf("Stateful Set %s:", ss.Name))
	result = append(result, fmt.Sprintf("Available replicas: %s", utils.Int32ToString(ss.Status.AvailableReplicas)))
	result = append(result, fmt.Sprintf("Replicas: %s", utils.Int32ToString(ss.Status.Replicas)))
	fmt.Println(strings.Join(result, "\r\n"))
	return nil
}

func describeControllerRevision(k8sclient *kubernetes.Clientset, namespace string, name string) error {
	cr, err := k8sclient.AppsV1().ControllerRevisions(namespace).Get(context.Background(), name, v1.GetOptions{})
	if err != nil {
		return err
	}
	var result []string
	result = append(result, fmt.Sprintf("Controller Revision %s:", cr.Name))
	result = append(result, fmt.Sprintf("Revision: %s", utils.Int32ToString(int32(cr.Revision))))
	fmt.Println(strings.Join(result, "\r\n"))
	return nil
}

func describeDaemonSet(k8sclient *kubernetes.Clientset, namespace string, name string) error {
	ds, err := k8sclient.AppsV1().DaemonSets(namespace).Get(context.Background(), name, v1.GetOptions{})
	if err != nil {
		return err
	}
	var result []string
	result = append(result, fmt.Sprintf("Daemon Set %s:", ds.Name))
	result = append(result, fmt.Sprintf("Number ready: %s", utils.Int32ToString(ds.Status.NumberReady)))
	if len(ds.Spec.Template.Spec.Containers) > 0 {
		result = append(result, "Containers:")
		for _, container := range ds.Spec.Template.Spec.Containers {
			result = append(result, fmt.Sprintf("%s (image: %s)", &container.Name, container.Image))
		}
	}
	fmt.Println(strings.Join(result, "\r\n"))
	return nil
}

func describeJob(k8sclient *kubernetes.Clientset, namespace string, name string) error {
	j, err := k8sclient.BatchV1().Jobs(namespace).Get(context.Background(), name, v1.GetOptions{})
	if err != nil {
		return err
	}
	var result []string
	result = append(result, fmt.Sprintf("Job %s:", j.Name))
	result = append(result, fmt.Sprintf("Active: %s", utils.Int32ToString(j.Status.Active)))
	result = append(result, fmt.Sprintf("Failed: %s", utils.Int32ToString(j.Status.Failed)))
	result = append(result, fmt.Sprintf("Succeeded: %s", utils.Int32ToString(j.Status.Succeeded)))
	result = append(result, fmt.Sprintf("Start time: %s", j.Status.StartTime.Time))
	result = append(result, fmt.Sprintf("Completion time: %s", j.Status.CompletionTime.Time))
	if len(j.Spec.Template.Spec.Containers) > 0 {
		result = append(result, "Containers:")
		for _, container := range j.Spec.Template.Spec.Containers {
			result = append(result, fmt.Sprintf("%s (image: %s)", container.Name, container.Image))
		}
	}
	fmt.Println(strings.Join(result, "\r\n"))
	return nil
}

func describeCronjob(k8sclient *kubernetes.Clientset, namespace string, name string) error {
	cj, err := k8sclient.BatchV1().CronJobs(namespace).Get(context.Background(), name, v1.GetOptions{})
	if err != nil {
		return err
	}
	var result []string
	result = append(result, fmt.Sprintf("Cron Job %s:", cj.Name))
	result = append(result, fmt.Sprintf("Schedule: %s", cj.Spec.Schedule))
	result = append(result, fmt.Sprintf("Time zone: %s", cj.Spec.TimeZone))

	if len(cj.Spec.JobTemplate.Spec.Template.Spec.Containers) > 0 {
		result = append(result, "Containers:")
		for _, container := range cj.Spec.JobTemplate.Spec.Template.Spec.Containers {
			result = append(result, fmt.Sprintf("%s (image: %s)", container.Name, container.Image))
		}
	}
	fmt.Println(strings.Join(result, "\r\n"))
	return nil
}

func describeHPA(k8sclient *kubernetes.Clientset, namespace string, name string) error {
	hpa, err := k8sclient.AutoscalingV1().HorizontalPodAutoscalers(namespace).Get(context.Background(), name, v1.GetOptions{})
	if err != nil {
		return err
	}
	var result []string
	result = append(result, fmt.Sprintf("Horizontal Pod Autoscaler %s:", hpa.Name))
	result = append(result, fmt.Sprintf("Min replicas: %s", utils.Int32ToString(*hpa.Spec.MinReplicas)))
	result = append(result, fmt.Sprintf("Max replicas: %s", utils.Int32ToString(hpa.Spec.MaxReplicas)))
	result = append(result, fmt.Sprintf("Target CPU Utilization Percentage: %s", utils.Int32ToString(*hpa.Spec.TargetCPUUtilizationPercentage)))
	fmt.Println(strings.Join(result, "\r\n"))
	return nil
}

func describeService(k8sclient *kubernetes.Clientset, namespace string, name string) error {
	s, err := k8sclient.CoreV1().Services(namespace).Get(context.Background(), name, v1.GetOptions{})
	if err != nil {
		return err
	}
	var result []string
	result = append(result, fmt.Sprintf("Service %s:", s.Name))
	result = append(result, fmt.Sprintf("Cluster IP: %s", s.Spec.ClusterIP))
	result = append(result, fmt.Sprintf("External Name: %s", s.Spec.ExternalName))
	result = append(result, fmt.Sprintf("External IPs: %s", strings.Join(s.Spec.ExternalIPs, ", ")))
	result = append(result, fmt.Sprintf("Load Balancer IP: %s", s.Spec.LoadBalancerIP))
	fmt.Println(strings.Join(result, "\r\n"))
	return nil
}

func describeIngress(k8sclient *kubernetes.Clientset, namespace string, name string) error {
	i, err := k8sclient.NetworkingV1().Ingresses(namespace).Get(context.Background(), name, v1.GetOptions{})
	if err != nil {
		return err
	}
	var result []string
	result = append(result, fmt.Sprintf("Ingres %s:", i.Name))
	result = append(result, "Rules: ")
	for _, rules := range i.Spec.Rules {
		result = append(result, fmt.Sprintf("Host: %s", rules.Host))
		for _, http := range rules.HTTP.Paths {
			result = append(result, fmt.Sprintf("%s => %s", http.Path, http.Backend.Service.Name))
		}
	}
	fmt.Println(strings.Join(result, "\r\n"))
	return nil
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
		_, err := os.Stat(application.Chart)
		if err != nil {
			messages = append(messages, fmt.Sprintf("--> Application %d: Chart file can't be opened (%s)", index, application.Chart))
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
