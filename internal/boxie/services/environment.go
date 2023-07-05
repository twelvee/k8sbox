// Package services contains buisness-logic methods of the models
package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/twelvee/boxie/pkg/boxie/utils"
	"helm.sh/helm/v3/pkg/chartutil"
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
		CreateTempDir:              createTempDir,
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
func GetConfigFromKubeconfig(namespace string, kubeconfig string) (*rest.Config, error) {
	restClientGetter := kube.GetConfig(kubeconfig, "", namespace)
	rc, err := restClientGetter.ToRESTConfig()
	if err != nil {
		return nil, err
	}

	return rc, nil
}

func prepareToWorkWithNamespace(namespace string) error {
	restConfig, err := GetConfigFromKubeconfig(namespace, os.Getenv("KUBECONFIG"))
	if err != nil {
		return err
	}
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

func createTempDir(environment *structs.Environment) error {
	dir := "/tmp/boxie_" + utils.GetShortID(8)
	err := os.Mkdir(dir, 0750)
	if err != nil {
		return err
	}
	environment.TempDeployDirectory = dir
	for _, b := range environment.Boxes {
		err = os.Mkdir(dir+"/"+b.Name, 0750)
		if err != nil {
			return err
		}

		err = os.WriteFile(dir+"/"+b.Name+"/"+chartutil.ChartfileName, []byte(b.Chart), 0644)
		if err != nil {
			return err
		}
		if len(strings.TrimSpace(b.Values)) != 0 {
			err = os.WriteFile(dir+"/"+b.Name+"/"+chartutil.ValuesfileName, []byte(b.Values), 0644)
			if err != nil {
				return err
			}
		}

		err = os.Mkdir(dir+"/"+b.Name+"/templates", 0750)
		if err != nil {
			return err
		}
		for _, a := range b.Applications {
			err := os.WriteFile(dir+"/"+b.Name+"/templates/"+a.Name, []byte(a.Chart), 0644)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
