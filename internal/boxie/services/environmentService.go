// Package services contains buisness-logic methods of the models
package services

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/twelvee/boxie/pkg/boxie/structs"
	"helm.sh/helm/v3/pkg/kube"
	v1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// NewEnvironmentService creates a new EnvironmentService
func NewEnvironmentService() structs.EnvironmentService {
	return structs.EnvironmentService{
		DeployEnvironment:          deployEnvironment,
		DeleteEnvironment:          deleteEnvironment,
		ValidateEnvironment:        validateEnvironment,
		ExpandVariables:            expandVariables,
		PrepareToWorkWithNamespace: prepareToWorkWithNamespace,
	}
}

func expandVariables(environment *structs.Environment) {
	environment.Name = os.ExpandEnv(environment.Name)
	environment.ID = os.ExpandEnv(environment.ID)
	environment.Namespace = os.ExpandEnv(environment.Namespace)
	environment.Variables = os.ExpandEnv(environment.Variables)
}

func deleteEnvironment(environment *structs.Environment) error {
	for _, box := range environment.Boxes {
		_, err := uninstallBox(*environment, box)
		if err != nil {
			return err
		}
	}
	err := deleteSavedEnvironment(*environment)
	if err != nil {
		return err
	}
	return nil
}

var k8sclient *kubernetes.Clientset
var restConfig *rest.Config

// GetConfigFromKubeconfig is loading your Kubeconfig into configuration struct
func GetConfigFromKubeconfig(namespace string) *rest.Config {
	restClientGetter := kube.GetConfig(os.Getenv("KUBECONFIG"), "", namespace)
	rc, err := restClientGetter.ToRESTConfig()
	if err != nil {
		panic(err)
	}

	return rc
}

func prepareToWorkWithNamespace(namespace string) error {
	restConfig = GetConfigFromKubeconfig(namespace)
	cl, err := kubernetes.NewForConfig(restConfig)
	k8sclient = cl
	if err != nil {
		return err
	}
	return createNamespaceIfNotExists(namespace)
}

func createNamespaceIfNotExists(namespace string) error {
	_, err := k8sclient.CoreV1().Namespaces().Create(context.Background(), &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: namespace,
		},
		TypeMeta: metav1.TypeMeta{
			Kind:       "Namespace",
			APIVersion: "v1",
		},
	}, metav1.CreateOptions{})
	if k8serrors.IsAlreadyExists(err) {
		return nil
	}
	return err
}

func deployEnvironment(environment *structs.Environment) error {
	saveEnvironment(*environment)
	for _, box := range environment.Boxes {
		_, err := installBox(&box, *environment)
		if err != nil {
			return err
		}
	}

	return nil
}

func validateEnvironment(environment *structs.Environment) error {
	var messages []string
	if len(strings.TrimSpace(environment.ID)) == 0 {
		messages = append(messages, "Environment id is missing")
	}

	if len(strings.TrimSpace(environment.Name)) == 0 {
		messages = append(messages, "Environment name is missing")
	}

	if len(strings.TrimSpace(environment.Variables)) > 0 {
		_, err := os.Stat(environment.Variables)
		if err != nil {
			messages = append(messages, fmt.Sprintf("Environment variables specified but the file is missing (%s)", environment.Variables))
		}
	}

	if len(environment.Boxes) == 0 {
		messages = append(messages, "Environment boxes are missing")
	}

	if len(messages) > 0 {
		return errors.New(strings.Join(messages, "\n\r"))
	}

	return nil
}
